package main

import "fmt"

func main() {

	var breakAgain bool

	for x := 0; x < 10; x++ {

		for y := 0; y < 10; y++ {

			if y == 2 {

				breakAgain = true

				break
			}
		}

		if breakAgain {
			break
		}
	}

	fmt.Println("done")
}
