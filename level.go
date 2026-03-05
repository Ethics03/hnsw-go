package hnsw

import (
	"math"
	"math/rand/v2"
)

func RandomLevel(m int) int {
	mL := 1.0 / math.Log(float64(m))
	level := int(-math.Log(rand.Float64()) * mL)
	return level
}
