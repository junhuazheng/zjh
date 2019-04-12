/*映射（map），它将为一键映射到纸。键是用于在检索值的对象。
给定一个键和一个值就可以在map对象中设置值。这只存储值后，就可以使用其键检索它对应的值了*/
package main

import "fmt"

func main() {

	var ccm map[string]string

	ccm = make(map[string]string)

	ccm["France"] = "Paris"
	ccm["Italy"] = "Rome"
	ccm["Japan"] = "Tokyo"
	ccm["India"] = "Delhi"

	for country := range ccm {
		fmt.Println("Capital of", country, "is", ccm[country])
	}

	capital, ok := ccm["United States"]

	if ok {
		fmt.Println("Capital of United States is", capital)
	} else {
		fmt.Println("Capital of United States is not present")
	}

}
