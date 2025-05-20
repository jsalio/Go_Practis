package main

import (
	"fmt"
)

func create_data() map[string]int {
	var dataMap map[string]int = make(map[string]int)
	dataMap["alex"] = 12
	dataMap["ass"] = 2
	dataMap["xina"] = 21
	dataMap["olef"] = 8
	dataMap["xela"] = 10
	return dataMap
}

func delete_from_map(key string, set map[string]int) (map[string]int, bool) {

	if _, ok := set[key]; !ok {
		return set, false
	}
	delete(set, key)
	//realmente no es necesario retornar el map ya que el dato se borra por referencia
	return set, true
}

func display_content(set map[string]int) {
	for key, value := range set {
		fmt.Printf("Key {%s} , value {%d} \n", key, value)
	}
}

func main() {
	fmt.Println("Eliminando data de un map")
	dataSet := create_data()
	display_content(dataSet)
	delete_from_map("olef", dataSet)
	fmt.Println("before delete olef")
	display_content(dataSet)
}
