package helpers

import (
	"math"
	"math/rand"
	"time"
)

func GenerateRandomNumber(len int) int {
	rand.Seed(time.Now().UnixNano())
	min := math.Pow(10, float64(len-1))
	max := math.Pow(10, float64(len)) - 1

	return rand.Intn(int(max-min)) + int(min)
}
