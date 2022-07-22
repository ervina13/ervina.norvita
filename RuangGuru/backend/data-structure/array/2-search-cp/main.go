package main

import "fmt"

func main() {
	var cars1 = []string{"Toyota", "Honda", "Nissan", "BMW", "Chevy", "Ford"}
	var cars2 = []string{"Ford", "BMW", "Audi", "Mercedes"}

	result, err := SearchMatch(cars1, cars2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Matched car brand: ", result)
	}
}

func SearchMatch(arr1 []string, arr2 []string) ([]string, error) {
	// sediakan variabel buat return []string
	var result []string

	// nested looping -> check val si arr1 == val si arr2 -> append variabel return
	for _, value := range arr1 {
		for _, value2 := range arr2 {
			if value == value2 {
				result = append(result, value)
			}
		}
	}
	// check length return masih kosong -> return error "no match"
	if len(result) == 0 {
		return nil, fmt.Errorf("no match")
	}
	// return result
	return result, nil
}
