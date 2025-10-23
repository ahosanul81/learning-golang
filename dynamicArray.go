package main

import "fmt"

func dynamicArray() {
	var studentsName []string
	var studentName string
	var numOfStudents int

	fmt.Printf("Number of students: ")
	fmt.Scan(&numOfStudents)
	for i := 0; i < numOfStudents; i++ {
		fmt.Print("Enter your name")
		fmt.Scan(&studentName)
		studentsName = append(studentsName, studentName)
	}
	fmt.Println(studentsName)
}
