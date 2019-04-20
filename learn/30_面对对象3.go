/*
多态
*/

package main

import "fmt"

type Human interface {
	speak(language string)
}

type Chinese struct{}

type American struct{}

func (ch Chinese) speak(language string) {
	fmt.Printf("speak %s\n", language)
}

func (am American) speak(language string) {
	fmt.Printf("speak %s\n", language)
}

func main() {
	var ch Human
	var am Human

	ch = Chinese{}
	am = American{}

	ch.speak("Chinese")
	am.speak("English")

	abcd := Chinese{} //这个结构体实现了接口的speak方法。
	abcd.speak("Riyu")
}

/*
面向接口编程，结构体中只要有和接口中定义的同名方法，就称该结构体是实现了哪个接口。
*/
