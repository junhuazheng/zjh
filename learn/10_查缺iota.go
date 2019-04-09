package main

import "fmt"

func main() {

	type Weapon int

	const (
		Arrow = iota
		Shur
		Snil
		Rif
		Blow
	)

	fmt.Println(Arrow, Shur, Snil, Rif, Blow)

	var weapon Weapon = Blow
	fmt.Println(weapon)

	//在代码中编写一些标志位时，我们往往手动编写常量值，如果常量值特别多时
	//很容易重复或者写错。因此，使用iota自动生成较为方便

	const (
		Fa = 1 << iota
		Fb
		Fc
		Fd
	)

	fmt.Printf("%d %d %d\n", Fa, Fb, Fc)
	fmt.Printf("%b %b %b\n", Fa, Fb, Fc)
}
