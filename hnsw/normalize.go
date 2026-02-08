package hnsw

import "math"

// normalizing for faster cosineDistance
func Normalize(vec []float64) {
	var norm float64

	for _, x := range vec {
		norm += x * x
	}

	norm = math.Sqrt(norm)

	if norm == 0 {
		return
	}

	for i := range vec {
		vec[i] /= norm
	}
}
