package main

import (
	"fmt"
	"time"
)

func main() {

	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(
		2012, 12, 12, 12, 12, 12, 66666, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then) //方法Sub返回一个Duration来表示两个时间点的间隔时间
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff)) //可以使用Add将时间后移一个时间间隔，或者使用一个-来将时间前移一个时间间隔
}
