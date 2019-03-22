package main

import "fmt"

type Actor struct {
}

func (a *Actor) OnEvent(param interface{}) {
	fmt.Println("actor event:", param)
}

func GlobalEvent(param interface{}) {
	fmt.Println("global event:", param)
}

func main() {

	a := new(Actor)

	REgisterEvent("OnSkill", a.OnEvent)

	REgisterEvent("Onskill", GlobalEvent)

	CallEvent("Onskill", 100)
}
