package main

import(
	"bytes"
	"fmt"
)

func printTypevalue(slist ...interface{}) string {
	
	var b bytes.Buffer
	
	for _, s := range slist {
		
		str := fmt.Sprintf("%v", s)
		
		var typeString string
		
		swaitch s.(type) {
			case bool:
			typeString = "bool"
			case string:
			typeString = "string"
			case int:
			typeString = "int"
		}
		
		b.WriteString("value:")
		b.WriteString(str)
		b.WriteString("type:")
		b.WriteString(typeString)
		b.WriteString("\n")
	}
	return b.String()
}

func main() {
	fmt.Println(printTypevalue(100, "str", true))
}