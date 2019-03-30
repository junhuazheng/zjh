package main

import "fmt"

const (
	Awsl = 99

	Ansl = Awsl * 10

	Atsl = Ansl * 10
)

func awslTsl(a int) (w, n, t int) {

	t = a / Atsl
	n = a / Ansl
	w = a / Awsl

	return
}

func main() {

	fmt.Println(awslTsl(999))

	t, n, w := awslTsl(99999)

	fmt.Println(t, n, w)
}
