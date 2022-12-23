package main

import "fmt"

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

func main() {
	fmt.Println(employees["Benjamin"])

	over21 := 0

	for _, value := range employees {
		if value >= 21 {
			over21++
		}
	}

	employees["Federico"] = 25

	delete(employees, "Pedro")

	fmt.Println(employees)
}
