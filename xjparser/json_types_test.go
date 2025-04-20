package xjparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsObject(t *testing.T) {
	jsonObject := NewJsonObject()
	jsonArray := NewJsonArray()
	jsonInt := NewJsonInt(1)
	jsonFloat := NewJsonFloat(1)
	jsonBool := NewJsonBool(false)
	jsonStr := NewJsonStr("bla")
	jsonNull := NewJsonNull()

	assert.True(t, IsObject(jsonObject))
	assert.False(t, IsObject(jsonArray))
	assert.False(t, IsObject(jsonInt))
	assert.False(t, IsObject(jsonFloat))
	assert.False(t, IsObject(jsonBool))
	assert.False(t, IsObject(jsonStr))
	assert.False(t, IsObject(jsonNull))
}

func TestIsArray(t *testing.T) {
	jsonObject := NewJsonObject()
	jsonArray := NewJsonArray()
	jsonInt := NewJsonInt(1)
	jsonFloat := NewJsonFloat(1)
	jsonBool := NewJsonBool(false)
	jsonStr := NewJsonStr("bla")
	jsonNull := NewJsonNull()

	assert.False(t, IsArray(jsonObject))
	assert.True(t, IsArray(jsonArray))
	assert.False(t, IsArray(jsonInt))
	assert.False(t, IsArray(jsonFloat))
	assert.False(t, IsArray(jsonBool))
	assert.False(t, IsArray(jsonStr))
	assert.False(t, IsArray(jsonNull))
}

func TestIsInt(t *testing.T) {
	jsonObject := NewJsonObject()
	jsonArray := NewJsonArray()
	jsonInt := NewJsonInt(1)
	jsonFloat := NewJsonFloat(1)
	jsonBool := NewJsonBool(false)
	jsonStr := NewJsonStr("bla")
	jsonNull := NewJsonNull()

	assert.False(t, IsInt(jsonObject))
	assert.False(t, IsInt(jsonArray))
	assert.True(t, IsInt(jsonInt))
	assert.False(t, IsInt(jsonFloat))
	assert.False(t, IsInt(jsonBool))
	assert.False(t, IsInt(jsonStr))
	assert.False(t, IsInt(jsonNull))
}

func TestIsFloat(t *testing.T) {
	jsonObject := NewJsonObject()
	jsonArray := NewJsonArray()
	jsonInt := NewJsonInt(1)
	jsonFloat := NewJsonFloat(1)
	jsonBool := NewJsonBool(false)
	jsonStr := NewJsonStr("bla")
	jsonNull := NewJsonNull()

	assert.False(t, IsFloat(jsonObject))
	assert.False(t, IsFloat(jsonArray))
	assert.False(t, IsFloat(jsonInt))
	assert.True(t, IsFloat(jsonFloat))
	assert.False(t, IsFloat(jsonBool))
	assert.False(t, IsFloat(jsonStr))
	assert.False(t, IsFloat(jsonNull))
}

func TestIsBool(t *testing.T) {
	jsonObject := NewJsonObject()
	jsonArray := NewJsonArray()
	jsonInt := NewJsonInt(1)
	jsonFloat := NewJsonFloat(1)
	jsonBool := NewJsonBool(false)
	jsonStr := NewJsonStr("bla")
	jsonNull := NewJsonNull()

	assert.False(t, IsBool(jsonObject))
	assert.False(t, IsBool(jsonArray))
	assert.False(t, IsBool(jsonInt))
	assert.False(t, IsBool(jsonFloat))
	assert.True(t, IsBool(jsonBool))
	assert.False(t, IsBool(jsonStr))
	assert.False(t, IsBool(jsonNull))
}

func TestIsStr(t *testing.T) {
	jsonObject := NewJsonObject()
	jsonArray := NewJsonArray()
	jsonInt := NewJsonInt(1)
	jsonFloat := NewJsonFloat(1)
	jsonBool := NewJsonBool(false)
	jsonStr := NewJsonStr("bla")
	jsonNull := NewJsonNull()

	assert.False(t, IsStr(jsonObject))
	assert.False(t, IsStr(jsonArray))
	assert.False(t, IsStr(jsonInt))
	assert.False(t, IsStr(jsonFloat))
	assert.False(t, IsStr(jsonBool))
	assert.True(t, IsStr(jsonStr))
	assert.False(t, IsStr(jsonNull))
}

func TestIsNull(t *testing.T) {
	jsonObject := NewJsonObject()
	jsonArray := NewJsonArray()
	jsonInt := NewJsonInt(1)
	jsonFloat := NewJsonFloat(1)
	jsonBool := NewJsonBool(false)
	jsonStr := NewJsonStr("bla")
	jsonNull := NewJsonNull()

	assert.False(t, IsNull(jsonObject))
	assert.False(t, IsNull(jsonArray))
	assert.False(t, IsNull(jsonInt))
	assert.False(t, IsNull(jsonFloat))
	assert.False(t, IsNull(jsonBool))
	assert.False(t, IsNull(jsonStr))
	assert.True(t, IsNull(jsonNull))
}
