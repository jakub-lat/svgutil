package utils

import "math/rand"

func RandRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func RandRangeFloat(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
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
