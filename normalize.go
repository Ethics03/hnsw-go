package hnsw

import "math"

// normalizing for faster cosineDistance
func Normalize(vec []float32) {
	var norm float32

	for _, x := range vec {
		norm += x * x
	}

	norm = float32(math.Sqrt(float64(norm)))

	if norm == 0 {
		return
	}

	for i := range vec {
		vec[i] /= norm
	}
}
