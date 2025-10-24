package main

import "fmt"

func higherOrderFn(x int, y int, op func(m int, n int)) {
	z := x + y
	op(x, y)
	fmt.Println("in higher order fn =", z)
}

func add(a int, b int) {
	c := a + b
	fmt.Println("add= ", c)
}

func higherOrderFnMain() {
	higherOrderFn(2, 4, add)
}
