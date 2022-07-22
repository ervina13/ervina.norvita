package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	people := []Person{{name: "Bob", age: 21}, {name: "Sam", age: 28}, {name: "Ann", age: 21}, {name: "Joe", age: 22}, {name: "Ben", age: 28}}
	fmt.Println(AgeDistribution(people))
	fmt.Println(FilterByAge(people, 21))
}

func AgeDistribution(people []Person) map[int]int {
	var result map[int]int
	result = map[int]int{}

	for i := 0; i < len(people); i++ {
		value, ok := result[people[i].age]
		if ok {
			value++
			result[people[i].age] = value
		} else {
			result[people[i].age] = 1
		}
	}

	return result // TODO: replace this
}

func FilterByAge(people []Person, age int) []Person {
	var person []Person
	for i := 0; i < len(people); i++ {
		if people[i].age == age {
			newPerson := Person{
				age:  people[i].age,
				name: people[i].name,
			}
			person = append(person, newPerson)
		}
	}
	return person // TODO: replace this
}
