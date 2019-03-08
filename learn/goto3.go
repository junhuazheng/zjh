package main

import "fmt"

func main() {

	err := firstCheckError()
	if err != nil {
		goto onExit
	}

	err = secondCheckError()
	if err != nil {
		goto onExit
	}

	fmt.Println("done")

	return

onExit:
	fmt.Println(err)
	exitProcess()
}
