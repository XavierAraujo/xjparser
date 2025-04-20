package antlr4_parser

import (
	"testing"
	"xjparser/xjparser"

	"github.com/stretchr/testify/assert"
)

func TestInvalidString(t *testing.T) {
	_, err := Parse("non-json-string")
	_, ok := err.(JsonParsingError)
	assert.NotNil(t, err)
	assert.True(t, ok)
}

/*
	func TestNoTopLevelObject(t *testing.T) {
		_, err := Parse("{},{}")
		_, ok := err.(JsonParsingError)
		assert.NotNil(t, err)
		assert.True(t, ok)
	}
*/
func TestEmptyObjectWithNoWhiteSpaces(t *testing.T) {
	json, err := Parse("{}")
	assert.Nil(t, err)
	jsonObject, ok := (*json).(*xjparser.JsonObject)
	assert.True(t, ok)
	assert.Equal(t, 0, jsonObject.Len())
}

func TestEmptyObjectWithWhiteSpaces(t *testing.T) {
	json, err := Parse(" { } ")
	assert.Nil(t, err)
	jsonObject, ok := (*json).(*xjparser.JsonObject)
	assert.True(t, ok)
	assert.Equal(t, 0, jsonObject.Len())
}

func TestObjectWithSingleSimpleKeyValue(t *testing.T) {
	json, err := Parse("{\"key1\": 1}")
	assert.Nil(t, err)
	jsonObject, ok := (*json).(*xjparser.JsonObject)
	assert.True(t, ok)
	assert.Equal(t, 1, jsonObject.Len())

	jsonInt, ok := jsonObject.GetKey("key1").(*xjparser.JsonInt)
	assert.True(t, ok)
	assert.Equal(t, int64(1), jsonInt.Value())
}

func TestObjectWithAllBasicTypes(t *testing.T) {
	json, err := Parse("{\"key1\": 1, \"key2\": 2.2, \"key3\": false, \"key4\": \"value\", \"key5\": null}")
	assert.Nil(t, err)
	jsonObject, ok := (*json).(*xjparser.JsonObject)
	assert.True(t, ok)
	assert.Equal(t, 5, jsonObject.Len())

	jsonInt, ok := jsonObject.GetKey("key1").(*xjparser.JsonInt)
	assert.True(t, ok)
	assert.Equal(t, int64(1), jsonInt.Value())

	jsonFloat, ok := jsonObject.GetKey("key2").(*xjparser.JsonFloat)
	assert.True(t, ok)
	assert.Equal(t, float64(2.2), jsonFloat.Value())

	jsonBool, ok := jsonObject.GetKey("key3").(*xjparser.JsonBool)
	assert.True(t, ok)
	assert.Equal(t, false, jsonBool.Value())

	jsonStr, ok := jsonObject.GetKey("key4").(*xjparser.JsonStr)
	assert.True(t, ok)
	assert.Equal(t, "value", jsonStr.Value())

	jsonNull, ok := jsonObject.GetKey("key5").(*xjparser.JsonNull)
	assert.NotNil(t, jsonNull)
	assert.True(t, ok)
}

func TestRepeatedKeys(t *testing.T) {
	_, err := Parse("{\"key1\": 1, \"key1\": 2.2}")
	duplicatedKeyError, ok := err.(JsonDuplicatedKeyError)
	assert.NotNil(t, err)
	assert.True(t, ok)
	assert.Equal(t, "key1", duplicatedKeyError.key)
}

func TestNestedRepeatedKeys(t *testing.T) {
	json, err := Parse("{\"key1\": {\"key1\": 2.2}}")
	assert.Nil(t, err)
	jsonObject1, ok := (*json).(*xjparser.JsonObject)
	assert.True(t, ok)
	assert.Equal(t, 1, jsonObject1.Len())

	jsonObject2, ok := jsonObject1.GetKey("key1").(*xjparser.JsonObject)
	assert.True(t, ok)
	assert.Equal(t, 1, jsonObject2.Len())

	jsonFloat, ok := jsonObject2.GetKey("key1").(*xjparser.JsonFloat)
	assert.True(t, ok)
	assert.Equal(t, 2.2, jsonFloat.Value())
}

func TestEmptyArrayWithNoWhiteSpaces(t *testing.T) {
	json, err := Parse("[]")
	assert.Nil(t, err)
	jsonArray, ok := (*json).(*xjparser.JsonArray)
	assert.True(t, ok)
	assert.Equal(t, 0, jsonArray.Len())
}

func TestEmptyArrayWithWhiteSpaces(t *testing.T) {
	json, err := Parse(" [ ] ")
	assert.Nil(t, err)
	jsonArray, ok := (*json).(*xjparser.JsonArray)
	assert.True(t, ok)
	assert.Equal(t, 0, jsonArray.Len())
}

func TestArrayWithSingleObject(t *testing.T) {
	json, err := Parse("[{\"key1\": 1}]")
	assert.Nil(t, err)
	jsonArray, ok := (*json).(*xjparser.JsonArray)
	assert.True(t, ok)
	assert.Equal(t, 1, jsonArray.Len())

	jsonObject := jsonArray.GetIndex(0)
	jsonInt := jsonObject.GetKey("key1").(*xjparser.JsonInt)
	assert.True(t, ok)
	assert.Equal(t, int64(1), jsonInt.Value())
}

func TestArraytWithAllBasicTypes(t *testing.T) {
	json, err := Parse("[{\"key1\": 1}, {\"key2\": 2.2}, {\"key3\": false}, {\"key4\": \"value\"}, {\"key5\": null}]")
	assert.Nil(t, err)
	jsonArray, ok := (*json).(*xjparser.JsonArray)
	assert.True(t, ok)
	assert.Equal(t, 5, jsonArray.Len())

	jsonObj := jsonArray.GetIndex(0)
	jsonInt, ok := jsonObj.GetKey("key1").(*xjparser.JsonInt)
	assert.True(t, ok)
	assert.Equal(t, int64(1), jsonInt.Value())

	jsonObj = jsonArray.GetIndex(1)
	jsonFloat, ok := jsonObj.GetKey("key2").(*xjparser.JsonFloat)
	assert.True(t, ok)
	assert.Equal(t, float64(2.2), jsonFloat.Value())

	jsonObj = jsonArray.GetIndex(2)
	jsonBool, ok := jsonObj.GetKey("key3").(*xjparser.JsonBool)
	assert.True(t, ok)
	assert.Equal(t, false, jsonBool.Value())

	jsonObj = jsonArray.GetIndex(3)
	jsonStr, ok := jsonObj.GetKey("key4").(*xjparser.JsonStr)
	assert.True(t, ok)
	assert.Equal(t, "value", jsonStr.Value())

	jsonObj = jsonArray.GetIndex(4)
	jsonNull, ok := jsonObj.GetKey("key5").(*xjparser.JsonNull)
	assert.NotNil(t, jsonNull)
	assert.True(t, ok)
}

func TestNestedObjectsAndArrays(t *testing.T) {
	json, err := Parse("{\"key1\": 5, \"key2\": {\"key3\": {\"key4\": 10, \"key5\": [{\"key6\": \"arrayValue1\"},{\"key7\": \"arrayValue2\"}]}}}")
	assert.Nil(t, err)
	jsonObject1, ok := (*json).(*xjparser.JsonObject)
	assert.True(t, ok)
	assert.Equal(t, 2, jsonObject1.Len())

	jsonKey1, ok := jsonObject1.GetKey("key1").(*xjparser.JsonInt)
	assert.True(t, ok)
	assert.Equal(t, int64(5), jsonKey1.Value())

	jsonObject2, ok := jsonObject1.GetKey("key2").(*xjparser.JsonObject)
	assert.True(t, ok)
	assert.Equal(t, 1, jsonObject2.Len())

	jsonObject3, ok := jsonObject2.GetKey("key3").(*xjparser.JsonObject)
	assert.True(t, ok)
	assert.Equal(t, 2, jsonObject3.Len())

	jsonInt, ok := jsonObject3.GetKey("key4").(*xjparser.JsonInt)
	assert.True(t, ok)
	assert.Equal(t, int64(10), jsonInt.Value())

	jsonArray, ok := jsonObject3.GetKey("key5").(*xjparser.JsonArray)
	assert.True(t, ok)
	assert.Equal(t, 2, jsonArray.Len())

	jsonArrayElem1 := jsonArray.GetIndex(0)
	assert.Equal(t, 1, jsonArrayElem1.Len())
	jsonArrayElem1Value, ok := jsonArrayElem1.GetKey("key6").(*xjparser.JsonStr)
	assert.True(t, ok)
	assert.Equal(t, "arrayValue1", jsonArrayElem1Value.Value())

	jsonArrayElem2 := jsonArray.GetIndex(1)
	assert.Equal(t, 1, jsonArrayElem2.Len())
	jsonArrayElem2Value, ok := jsonArrayElem2.GetKey("key7").(*xjparser.JsonStr)
	assert.True(t, ok)
	assert.Equal(t, "arrayValue2", jsonArrayElem2Value.Value())
}

func TestObjectEditingByReference(t *testing.T) {
	json, err := Parse("{\"key1\": {\"key2\": 1}}")
	assert.Nil(t, err)
	jsonObject1, ok := (*json).(*xjparser.JsonObject)
	assert.True(t, ok)
	assert.Equal(t, 1, jsonObject1.Len())

	jsonObject2, ok := jsonObject1.GetKey("key1").(*xjparser.JsonObject)
	jsonIntBefore, ok := jsonObject2.GetKey("key2").(*xjparser.JsonInt)
	assert.Equal(t, int64(1), jsonIntBefore.Value())

	jsonObject2.SetKey("key2", xjparser.NewJsonInt(2))
	jsonIntAfter, ok := jsonObject2.GetKey("key2").(*xjparser.JsonInt)
	assert.Equal(t, int64(2), jsonIntAfter.Value())
}
