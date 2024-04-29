package main

import (
	"fmt"

	utils "github.com/jffin/distinct-elements"
)

func main() {
	p := 14

	b := make(utils.Buckets, 1<<p)

	add := utils.Add(&b)

	times := 13

	m := map[string]struct{}{}

	for range 1 << times {

		s := utils.RandStringBytes(32)

		m[s] = struct{}{}

		add([]byte(s))
	}

	estimate := utils.Count(b)

	fmt.Println("times: ", 1<<times, "value: ", uint64(estimate))
	fmt.Println("times - unique: ", 1<<times-len(m))
	fmt.Println("diff: ", len(m)-int(estimate))
}
