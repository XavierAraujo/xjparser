package xjparser

import (
	"fmt"
)

type JSON_TYPE int

const (
	JSON_OBJECT JSON_TYPE = iota
	JSON_ARRAY
	JSON_INT
	JSON_FLOAT
	JSON_BOOL
	JSON_STR
	JSON_NULL
)

var jsonNullObject = &JsonNull{} // Small optimization: we can always use the same JsonNull object

type JsonValue interface {
	GetType() JSON_TYPE
}

/* Json Object */

type JsonObject struct {
	values map[string]JsonValue
}

func NewJsonObject() *JsonObject {
	return &JsonObject{values: make(map[string]JsonValue)}
}

func (jsonObject *JsonObject) GetType() JSON_TYPE {
	return JSON_OBJECT
}

func (jsonObject *JsonObject) GetKey(key string) JsonValue {
	return jsonObject.values[key]
}

func (jsonObject *JsonObject) SetKey(key string, value JsonValue) {
	jsonObject.values[key] = value
}

func (jsonObject *JsonObject) HasKey(key string) bool {
	_, keyExists := jsonObject.values[key]
	return keyExists
}

func (jsonObject *JsonObject) Len() int {
	return len(jsonObject.values)
}

func (jsonObject *JsonObject) String() string {
	return fmt.Sprint(jsonObject.values)
}

/* Json Array */

type JsonArray struct {
	values []JsonValue
}

func NewJsonArray() *JsonArray {
	return &JsonArray{values: make([]JsonValue, 0)}
}

func (jsonArray *JsonArray) GetType() JSON_TYPE {
	return JSON_ARRAY
}

func (jsonArray *JsonArray) GetIndex(index int) JsonValue {
	return jsonArray.values[index]
}

func (jsonArray *JsonArray) Add(jsonValue JsonValue) {
	jsonArray.values = append(jsonArray.values, jsonValue)
}

func (jsonArray *JsonArray) String() string {
	return fmt.Sprint(jsonArray.values)
}

func (jsonArray *JsonArray) Len() int {
	return len(jsonArray.values)
}

/* Json Int */

type JsonInt struct {
	value int64
}

func NewJsonInt(value int64) *JsonInt {
	return &JsonInt{value}
}

func (jsonInt *JsonInt) GetType() JSON_TYPE {
	return JSON_INT
}

func (jsonInt *JsonInt) Value() int64 {
	return jsonInt.value
}

func (jsonInt *JsonInt) String() string {
	return fmt.Sprint(jsonInt.value)
}

/* Json Float */

type JsonFloat struct {
	value float64
}

func NewJsonFloat(value float64) *JsonFloat {
	return &JsonFloat{value}
}

func (jsonFloat *JsonFloat) GetType() JSON_TYPE {
	return JSON_FLOAT
}

func (jsonFloat *JsonFloat) Value() float64 {
	return jsonFloat.value
}

func (jsonFloat *JsonFloat) String() string {
	return fmt.Sprint(jsonFloat.value)
}

/* Json Bool */

type JsonBool struct {
	value bool
}

func NewJsonBool(value bool) *JsonBool {
	return &JsonBool{value}
}

func (jsonBool *JsonBool) GetType() JSON_TYPE {
	return JSON_BOOL
}

func (jsonBool *JsonBool) Value() bool {
	return jsonBool.value
}

func (jsonBool *JsonBool) String() string {
	return fmt.Sprint(jsonBool.value)
}

/* Json Str */

type JsonStr struct {
	value string
}

func NewJsonStr(value string) *JsonStr {
	return &JsonStr{value}
}

func (jsonStr *JsonStr) GetType() JSON_TYPE {
	return JSON_STR
}

func (jsonStr *JsonStr) Value() string {
	return jsonStr.value
}

func (jsonStr *JsonStr) String() string {
	return jsonStr.value
}

/* Json Null */

type JsonNull struct {
}

func NewJsonNull() *JsonNull {
	return jsonNullObject
}

func (jsonNull *JsonNull) GetType() JSON_TYPE {
	return JSON_NULL
}

func (jsonNull *JsonNull) String() string {
	return "null"
}

/* ----- */

func IsObject(value JsonValue) bool {
	return value.GetType() == JSON_OBJECT
}

func IsArray(value JsonValue) bool {
	return value.GetType() == JSON_ARRAY
}

func IsInt(value JsonValue) bool {
	return value.GetType() == JSON_INT
}

func IsFloat(value JsonValue) bool {
	return value.GetType() == JSON_FLOAT
}

func IsBool(value JsonValue) bool {
	return value.GetType() == JSON_BOOL
}

func IsStr(value JsonValue) bool {
	return value.GetType() == JSON_STR
}

func IsNull(value JsonValue) bool {
	return value.GetType() == JSON_NULL
}
