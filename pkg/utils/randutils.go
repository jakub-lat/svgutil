package utils

import (
	"golang.org/x/exp/constraints"
	"math/rand"
)

func RandRange[T constraints.Integer | constraints.Float](min, max T) T {
	minf, maxf := float64(min), float64(max)
	return T(rand.Float64()*(maxf-minf) + minf)
}

func TruncatedNormal(mean, stdDev, low, high float64) int {
	if low >= high {
		panic("high must be greater than low")
	}

	for {
		x := rand.NormFloat64()*stdDev + mean
		if low <= x && x < high {
			return int(x)
		}
	}
}
