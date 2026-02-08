package hnsw

import "math"

func CosineDistance(a, b []float64) float64 {
	Normalize(a)
	Normalize(b)
	var dot, normA, normB float64

	for i := 0; i < len(a); i++ {
		dot += a[i] * b[i]
		normA += a[i] * a[i]
		normB += b[i] * b[i]
	}

	if normA == 0 || normB == 0 {
		return 1.0 // max distance if vector empty
	}

	similarity := dot / (math.Sqrt(normA) * math.Sqrt(normB))

	if similarity > 1 {
		similarity = 1
	} else if similarity < -1 {
		similarity = -1
	}
	return 1 - similarity
}
