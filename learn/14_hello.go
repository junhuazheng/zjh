package main

import "fmt"

func main() {
	var a int
	a = 10
	fmt.Printf("a = %d\n", a)

	b := 20
	fmt.Println("b = ", b)

	c := 30
	fmt.Printf("c type is %T\n", c)

	var d int
	var e int
	var f int
	d, e, f = 10, 20, 30
	fmt.Printf("d = %d, e = %d, f = %d\n", d, f, e)

	g := 3.14
	fmt.Printf("g type is %T\n", g)

	const (
		h = iota
		j = iota
		k = iota
	)
	fmt.Printf("h = %d, j = %d, k = %d\n", h, j, k)

	var ch byte
	ch = 98
	fmt.Printf("%c, %d\n", ch, ch)

	ch = 'b'
	fmt.Printf("%c, %d\n", ch, ch)
	fmt.Printf("%c\n", 'b'-32)
	fmt.Printf("%c\n", 'B'+32)

	var m bool
	fmt.Println("m = ", m)
}
