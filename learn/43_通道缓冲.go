package main

import "fmt"

func main() {

	messaage := make(chan string, 2)

	messaage <- "buffered"
	messaage <- "channel"

	fmt.Println(<-messaage)
	fmt.Println(<-messaage)
}
