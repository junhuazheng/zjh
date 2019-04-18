//下面的迭代还有什么问题？
package main

import (
	"fmt"
	"sync"
	"time"
)

type ThreadSafeSet struct {
	sync.RWMutex
	s []int
}

func (set *ThreadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.Lock()

		for elem := range set.s {
			ch <- elem
			fmt.Print("get", elem, ",")
		}

		close(ch)
		set.RUnlock()
	}()
	return ch
}

func main() {
	//reda()
	unRead()
}

func read() {
	set := ThreadSafeSet{}
	set.s = make([]int, 100)
	ch := set.Iter()
	closed := false

	for {
		select {
		case v, ok := <-ch:
			if ok {
				fmt.Print("read:", v, ",")
			} else {
				closed = true
			}
		}

		if closed {
			fmt.Print("closed")
			break
		}
	}
	fmt.Print("Done")
}

func unRead() {
	set := ThreadSafeSet{}
	set.s = make([]int, 100)
	ch := set.Iter()
	_ = ch
	time.Sleep(time.Second * 5)
	fmt.Print("Done")
}
