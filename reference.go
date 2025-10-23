package main

func callByReference(p *int, val int) {
	*p = val
}

func reference() {
	var p *int
	x := 4

	p = &x
	*p = *p - 1

	callByReference(&x, 40)

	// fmt.Printf("x= %v\n", x)
	// fmt.Printf("memory addess of x = %v\n", p)
	// fmt.Printf("value of p = %v\n", *p)

}
