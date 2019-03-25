//对结构体数据进行排序
//1、完整实现sort.Interface进行结构体排序

package main

import (
	"fmt"
	"sort"
)

type HeroKind int

const (
	None HeroKind = iota
	Tank
	Assassin
	Mage
)

type Hero struct {
	Name string
	Kind HeroKind
}

type Heros []*Hero

func (s Heros) len() int {
	return len(s)
}

func (s Heros) Less(i, j int) bool {
	
	if s[i],Kind != s[j].Kind {
		return s[i].Kind < s[j].Kind
	}
	
	return s[i].Name < s[j].Name
}

func (s Heros) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	
	heros := Heros{
		&Hero{"猪", Tank},
		&Hero{"牛", Assassin},
		&Hero{"马", Mage},
		&Hero{"狗", Assassin},
		&Hero{"猴", Tank},
		&Hero{"虎", Mage},
	}
	
	sort.Sort(heros)
	
	for _, v := range heros {
		fmt.Printf("%+v\n", v)
	}
}