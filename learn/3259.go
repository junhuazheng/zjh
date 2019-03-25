//使用sort.Slice()函数，对代码重新优化：
package main

import (
	"fmt"
	"sort"
)

type HeroKind int

const (
	None = iota
	Tank
	Assassin
	Mage
)

type Hero struct {
	Name string
	Kind HeroKind
}

func main() {

	heros := []*Hero{
		{"猪", Tank},
		{"牛", Assassin},
		{"马", Mage},
		{"猴", Assassin},
		{"狗", Tank},
		{"虎", Mage},
	}

	sort.Slice(heros, func(i, j int) bool {
		if heros[i].Kind != heros[j].Kind {
			return heros[i].Kind < heros[j].Kind
		}

		return heros[i].Name < heros[j].Name
	})

	for _, v := range heros {
		fmt.Printf("%+v\n", v)
	}
}
