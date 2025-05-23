package xjparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidString(t *testing.T) {
	_, err := Parse("non-json-string")
	_, ok := err.(JsonParsingError)
	assert.NotNil(t, err)
	assert.True(t, ok)
}

func TestNoTopLevelObject(t *testing.T) {
	_, err := Parse("{},{}")
	_, ok := err.(JsonParsingError)
	assert.NotNil(t, err)
	assert.True(t, ok)
}

func TestMultipleTopLevelObject(t *testing.T) {
	_, err := Parse("{}{}")
	_, ok := err.(JsonParsingError)
	assert.NotNil(t, err)
	assert.True(t, ok)
}

func TestEmptyObjectWithNoWhiteSpaces(t *testing.T) {
	json, err := Parse("{}")
	assert.Nil(t, err)
	jsonObject, ok := json.(*JsonObject)
	assert.True(t, ok)
	assert.Equal(t, 0, jsonObject.Len())
}

func TestEmptyObjectWithWhiteSpaces(t *testing.T) {
	json, err := Parse(" { } ")
	assert.Nil(t, err)
	jsonObject, ok := json.(*JsonObject)
	assert.True(t, ok)
	assert.Equal(t, 0, jsonObject.Len())
}

func TestObjectWithSingleSimpleKeyValue(t *testing.T) {
	json, err := Parse("{\"key1\": 1}")
	assert.Nil(t, err)
	jsonObject, ok := json.(*JsonObject)
	assert.True(t, ok)
	assert.Equal(t, 1, jsonObject.Len())

	jsonInt, ok := jsonObject.GetKey("key1").(*JsonInt)
	assert.True(t, ok)
	assert.Equal(t, int64(1), jsonInt.Value())
}

func TestObjectWithAllBasicTypes(t *testing.T) {
	json, err := Parse("{\"key1\": 1, \"key2\": 2.2, \"key3\": false, \"key4\": \"value\", \"key5\": null}")
	assert.Nil(t, err)
	jsonObject, ok := json.(*JsonObject)
	assert.True(t, ok)
	assert.Equal(t, 5, jsonObject.Len())

	jsonInt, ok := jsonObject.GetKey("key1").(*JsonInt)
	assert.True(t, ok)
	assert.Equal(t, int64(1), jsonInt.Value())

	jsonFloat, ok := jsonObject.GetKey("key2").(*JsonFloat)
	assert.True(t, ok)
	assert.Equal(t, float64(2.2), jsonFloat.Value())

	jsonBool, ok := jsonObject.GetKey("key3").(*JsonBool)
	assert.True(t, ok)
	assert.Equal(t, false, jsonBool.Value())

	jsonStr, ok := jsonObject.GetKey("key4").(*JsonStr)
	assert.True(t, ok)
	assert.Equal(t, "value", jsonStr.Value())

	jsonNull, ok := jsonObject.GetKey("key5").(*JsonNull)
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
	jsonObject1, ok := json.(*JsonObject)
	assert.True(t, ok)
	assert.Equal(t, 1, jsonObject1.Len())

	jsonObject2, ok := jsonObject1.GetKey("key1").(*JsonObject)
	assert.True(t, ok)
	assert.Equal(t, 1, jsonObject2.Len())

	jsonFloat, ok := jsonObject2.GetKey("key1").(*JsonFloat)
	assert.True(t, ok)
	assert.Equal(t, 2.2, jsonFloat.Value())
}

func TestEmptyArrayWithNoWhiteSpaces(t *testing.T) {
	json, err := Parse("[]")
	assert.Nil(t, err)
	jsonArray, ok := json.(*JsonArray)
	assert.True(t, ok)
	assert.Equal(t, 0, jsonArray.Len())
}

func TestEmptyArrayWithWhiteSpaces(t *testing.T) {
	json, err := Parse(" [ ] ")
	assert.Nil(t, err)
	jsonArray, ok := json.(*JsonArray)
	assert.True(t, ok)
	assert.Equal(t, 0, jsonArray.Len())
}

func TestArrayWithSingleObject(t *testing.T) {
	json, err := Parse("[{\"key1\": 1}]")
	assert.Nil(t, err)
	jsonArray, ok := json.(*JsonArray)
	assert.True(t, ok)
	assert.Equal(t, 1, jsonArray.Len())

	jsonObject, ok := jsonArray.GetIndex(0).(*JsonObject)
	assert.True(t, ok)
	jsonInt := jsonObject.GetKey("key1").(*JsonInt)
	assert.True(t, ok)
	assert.Equal(t, int64(1), jsonInt.Value())
}

func TestArrayWithMultipleObjectsWithin(t *testing.T) {
	json, err := Parse("[{\"key1\": 1}, {\"key2\": 2.2}, {\"key3\": false}, {\"key4\": \"value\"}, {\"key5\": null}]")
	assert.Nil(t, err)
	jsonArray, ok := json.(*JsonArray)
	assert.True(t, ok)
	assert.Equal(t, 5, jsonArray.Len())

	jsonObj, ok := jsonArray.GetIndex(0).(*JsonObject)
	assert.True(t, ok)
	jsonInt, ok := jsonObj.GetKey("key1").(*JsonInt)
	assert.True(t, ok)
	assert.Equal(t, int64(1), jsonInt.Value())

	jsonObj, ok = jsonArray.GetIndex(1).(*JsonObject)
	assert.True(t, ok)
	jsonFloat, ok := jsonObj.GetKey("key2").(*JsonFloat)
	assert.True(t, ok)
	assert.Equal(t, float64(2.2), jsonFloat.Value())

	jsonObj, ok = jsonArray.GetIndex(2).(*JsonObject)
	assert.True(t, ok)
	jsonBool, ok := jsonObj.GetKey("key3").(*JsonBool)
	assert.True(t, ok)
	assert.Equal(t, false, jsonBool.Value())

	jsonObj, ok = jsonArray.GetIndex(3).(*JsonObject)
	assert.True(t, ok)
	jsonStr, ok := jsonObj.GetKey("key4").(*JsonStr)
	assert.True(t, ok)
	assert.Equal(t, "value", jsonStr.Value())

	jsonObj, ok = jsonArray.GetIndex(4).(*JsonObject)
	assert.True(t, ok)
	jsonNull, ok := jsonObj.GetKey("key5").(*JsonNull)
	assert.NotNil(t, jsonNull)
	assert.True(t, ok)
}

func TestArrayWithMultipleTypes(t *testing.T) {
	json, err := Parse("[{\"key1\":\"value1\"}, [{\"key2\":\"value2\"}], 1, 1.1, true, \"plain-str\", null]")
	assert.Nil(t, err)
	jsonArray, ok := json.(*JsonArray)
	assert.True(t, ok)
	assert.Equal(t, 7, jsonArray.Len())

	jsonObj, ok := jsonArray.GetIndex(0).(*JsonObject)
	assert.True(t, ok)
	jsonStr, ok := jsonObj.GetKey("key1").(*JsonStr)
	assert.True(t, ok)
	assert.Equal(t, "value1", jsonStr.Value())

	subJsonArray, ok := jsonArray.GetIndex(1).(*JsonArray)
	assert.True(t, ok)
	jsonObj, ok = subJsonArray.GetIndex(0).(*JsonObject)
	jsonStr, ok = jsonObj.GetKey("key2").(*JsonStr)
	assert.True(t, ok)
	assert.Equal(t, "value2", jsonStr.Value())

	jsonInt, ok := jsonArray.GetIndex(2).(*JsonInt)
	assert.True(t, ok)
	assert.Equal(t, int64(1), jsonInt.Value())

	jsonFloat, ok := jsonArray.GetIndex(3).(*JsonFloat)
	assert.True(t, ok)
	assert.Equal(t, float64(1.1), jsonFloat.Value())

	jsonBool, ok := jsonArray.GetIndex(4).(*JsonBool)
	assert.True(t, ok)
	assert.Equal(t, true, jsonBool.Value())

	jsonStr, ok = jsonArray.GetIndex(5).(*JsonStr)
	assert.True(t, ok)
	assert.Equal(t, "plain-str", jsonStr.Value())

	_, ok = jsonArray.GetIndex(6).(*JsonNull)
	assert.True(t, ok)
}

func TestNestedObjectsAndArrays(t *testing.T) {
	json, err := Parse("{\"key1\": 5, \"key2\": {\"key3\": {\"key4\": 10, \"key5\": [{\"key6\": \"arrayValue1\"},{\"key7\": \"arrayValue2\"}]}}}")
	assert.Nil(t, err)
	jsonObject1, ok := json.(*JsonObject)
	assert.True(t, ok)
	assert.Equal(t, 2, jsonObject1.Len())

	jsonKey1, ok := jsonObject1.GetKey("key1").(*JsonInt)
	assert.True(t, ok)
	assert.Equal(t, int64(5), jsonKey1.Value())

	jsonObject2, ok := jsonObject1.GetKey("key2").(*JsonObject)
	assert.True(t, ok)
	assert.Equal(t, 1, jsonObject2.Len())

	jsonObject3, ok := jsonObject2.GetKey("key3").(*JsonObject)
	assert.True(t, ok)
	assert.Equal(t, 2, jsonObject3.Len())

	jsonInt, ok := jsonObject3.GetKey("key4").(*JsonInt)
	assert.True(t, ok)
	assert.Equal(t, int64(10), jsonInt.Value())

	jsonArray, ok := jsonObject3.GetKey("key5").(*JsonArray)
	assert.True(t, ok)
	assert.Equal(t, 2, jsonArray.Len())

	jsonArrayObj1, ok := jsonArray.GetIndex(0).(*JsonObject)
	assert.True(t, ok)
	assert.Equal(t, 1, jsonArrayObj1.Len())
	jsonArrayElem1Value, ok := jsonArrayObj1.GetKey("key6").(*JsonStr)
	assert.True(t, ok)
	assert.Equal(t, "arrayValue1", jsonArrayElem1Value.Value())

	jsonArrayObj2, ok := jsonArray.GetIndex(1).(*JsonObject)
	assert.True(t, ok)
	assert.Equal(t, 1, jsonArrayObj2.Len())
	jsonArrayElem2Value, ok := jsonArrayObj2.GetKey("key7").(*JsonStr)
	assert.True(t, ok)
	assert.Equal(t, "arrayValue2", jsonArrayElem2Value.Value())
}

func TestObjectEditingByReference(t *testing.T) {
	json, err := Parse("{\"key1\": {\"key2\": 1}}")
	assert.Nil(t, err)
	jsonObject1, ok := json.(*JsonObject)
	assert.True(t, ok)
	assert.Equal(t, 1, jsonObject1.Len())

	jsonObject2, ok := jsonObject1.GetKey("key1").(*JsonObject)
	jsonIntBefore, ok := jsonObject2.GetKey("key2").(*JsonInt)
	assert.Equal(t, int64(1), jsonIntBefore.Value())

	jsonObject2.SetKey("key2", NewJsonInt(2))
	jsonIntAfter, ok := jsonObject2.GetKey("key2").(*JsonInt)
	assert.Equal(t, int64(2), jsonIntAfter.Value())
}

func TestNegativeValues(t *testing.T) {
	json, err := Parse("{\"key1\": -100, \"key2\": -200.5, \"key3\": false, \"key4\": \"value\", \"key5\": null}")
	assert.Nil(t, err)
	jsonObject, ok := json.(*JsonObject)
	assert.True(t, ok)
	assert.Equal(t, 5, jsonObject.Len())

	jsonInt, ok := jsonObject.GetKey("key1").(*JsonInt)
	assert.True(t, ok)
	assert.Equal(t, int64(-100), jsonInt.Value())

	jsonFloat, ok := jsonObject.GetKey("key2").(*JsonFloat)
	assert.True(t, ok)
	assert.Equal(t, float64(-200.5), jsonFloat.Value())
}

func TestExponentialValues(t *testing.T) {
	json, _ := Parse("{\"e\": 0.123456789e-12}")
	jsonFloat := json.(*JsonObject).GetKey("e")
	assert.Equal(t, NewJsonFloat(0.000000000000123456789), jsonFloat)

	json, _ = Parse("{\"e\": 1.234567890E+34}")
	jsonFloat = json.(*JsonObject).GetKey("e")
	assert.Equal(t, NewJsonFloat(12345678900000000000000000000000000), jsonFloat)

	json, _ = Parse("{\"e\": 23456789012E66}")
	jsonFloat = json.(*JsonObject).GetKey("e")
	assert.Equal(t, NewJsonFloat(23456789012000000000000000000000000000000000000000000000000000000000000000000), jsonFloat)
}
