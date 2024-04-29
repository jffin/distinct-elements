package utils

import (
	"hash/fnv"
	"math"
)

const (
	p            = 14
	M            = 1 << p // number of buckets
	alphaM       = 0.7213 / (1 + 1.079/float64(M))
	maxBucketVal = 64 - p
)

type Buckets []uint8

func Add(b *Buckets) func([]byte) {
	return func(d []byte) {
		hash := fnv.New64()
		hash.Write(d)
		h := hash.Sum64()

		i := h >> maxBucketVal

		r := uint8(1)

		for (h & (1 << (maxBucketVal - 1))) == 0 {
			r++
			h <<= 1
		}

		if r > (*b)[i] {
			(*b)[i] = r
		}
	}
}

func Count(b Buckets) uint64 {
	var sum float64

	for _, r := range b {
		sum += 1.0 / math.Pow(2.0, float64(r))
	}

	estimate := alphaM * float64(M*M) / sum
	if estimate <= 2.5*float64(M) {
		zeros := 0
		for _, r := range b {
			if r == 0 {
				zeros++
			}
		}

		if zeros != 0 {
			estimate = float64(M) * math.Log(float64(M)/float64(zeros))
		}
	} else if estimate > (1<<32)/30.0 {
		estimate = -math.Pow(2, 32) * math.Log(1-estimate/math.Pow(2, 32))
	}

	return uint64(estimate)
}
