package hnsw

import "math"

func Normalize(vec []float32) []float32 {
	var norm float32

	for _, x := range vec {
		norm += x * x
	}

	norm = float32(math.Sqrt(float64(norm)))

	if norm == 0 {
		result := make([]float32, len(vec))
		copy(result, vec)
		return result
	}

	result := make([]float32, len(vec))
	for i := range vec {
		result[i] = vec[i] / norm
	}
	return result
}
