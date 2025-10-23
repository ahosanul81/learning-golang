package main

import "fmt"

func array() {
	var destinations = [5]string{"Bangladesh", "India", "pakisatan ", "Nepal", "Thailand"}
	for i := 0; i < len(destinations); i++ {

		fmt.Println(destinations[i])
	}
}
