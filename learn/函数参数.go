package main

import (
	"bytes"
	"fmt"
)

func joinStrings(slist ...string) string {

	var v bytes.Buffer

	for _, s := range slist {
		b.WriteString(s)
	}

	return b.String()
}

func main() {

	fmt.Println(joinStrings("pig", "and", "rat"))
	fmt.Println(joinStrings("hammer", "mom", "and", "hawk"))
}
