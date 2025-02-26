// example go script with maps
package main

import (
	"fmt"
)

func main() {
	//declaring a map with string keys and int values
	//the map is empty
	mapa := make(map[string]int)
	fmt.Println("Empty map declared with 'make':", mapa)

	//inserting values into the map
	mapa["key1"] = 1
	mapa["key2"] = 2
	mapa["key3"] = 3
	fmt.Println("Map with values:", mapa)

	_, exists := mapa["key2"]
	fmt.Println("Key2 exists:", exists)

	//deleting a key from the map
	delete(mapa, "key2")
	fmt.Println("Map with key2 removed:", mapa)

	//checking if a key exists in the map
	_, exists = mapa["key2"]
	fmt.Println("Key2 exists:", exists)
}
