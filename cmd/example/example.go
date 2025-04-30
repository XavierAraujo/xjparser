package main

import (
	"fmt"
	"xjparser/xjparser"
)

func main() {
	input := "{\"key1\": {\"key1.1\": \"bli\"}, \"key2\": 10, \"key3\": [10E10]}"
	fmt.Println("Parsing string: ", input)

	topLevelEntity, err := xjparser.Parse(input)
	if err != nil {
		fmt.Println("And error occurred parsing the JSON string")
		return
	}
	topLevelObject := topLevelEntity.(*xjparser.JsonObject)

	fmt.Println("Is top level entity an object: ", xjparser.IsObject(topLevelObject))
	fmt.Println("Is top level entity an array: ", xjparser.IsArray(topLevelObject))

	nonExistentKeyObject := topLevelObject.GetKey("non-existent-key")
	fmt.Println("Non existent key object: ", nonExistentKeyObject)

	key1Object := topLevelObject.GetKey("key1")
	fmt.Println("Key1 Object: ", key1Object)

	key2Object := topLevelObject.GetKey("key2")
	fmt.Println("Key2 Object: ", key2Object)
	fmt.Println("Is Key2 Object an int: ", xjparser.IsInt(topLevelObject))
	fmt.Println("Is Key2 Object a float: ", xjparser.IsFloat(topLevelObject))

	key3Array := topLevelObject.GetKey("key3")
	fmt.Println("Key3 Array: ", key3Array)
}
