package main

import (
	"log"
	"math"
	"math/rand"
	"sync"

	utils "github.com/jffin/distinct-elements"
)

func main() {
	var wg sync.WaitGroup

	c := make(chan string)

	buckets := make(utils.Buckets, utils.M)

	b := &buckets

	add := utils.Add(b)

	var m = map[string]struct{}{}
	for range 1 << 13 {
		m[utils.RandStringBytes(rand.Intn(32))] = struct{}{}
	}

	var s []string
	for k := range m {
		s = append(s, k)
	}

	var unique = map[string]struct{}{}
	var elements uint32

	for range 2 << 11 {
		wg.Add(1)

		go func() {
			defer wg.Done()

			rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })

			s := s[:rand.Intn(len(s))]

			for i := 0; i < len(s); i++ {
				if rand.Intn(120) >= 119 {
					n := rand.Intn(10)
					for n > i {
						n -= 1
					}

					i -= n
				}

				v := s[rand.Intn(len(s))]
				c <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for value := range c {
		unique[value] = struct{}{}

		elements++

		add([]byte(value))
	}

	r := utils.Count(*b)

	log.Printf("result: %d | elements: %d | unique: %d\n", r, elements, len(unique))
	log.Printf("diff: %d\n", (len(unique)-int(r))-int(math.Pow(2, math.Log2(float64(len(unique)))-3)))
}
