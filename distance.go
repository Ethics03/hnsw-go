package hnsw

import "math"

func CosineDistance(a, b []float32) float32 {
	Normalize(a)
	Normalize(b)
	var dot, normA, normB float32

	for i := range len(a) {
		dot += a[i] * b[i]
		normA += a[i] * a[i]
		normB += b[i] * b[i]
	}

	if normA == 0 || normB == 0 {
		return 1.0 // max distance if vector empty
	}

	similarity := dot / (float32(math.Sqrt(float64(normA))) * float32(math.Sqrt(float64(normB))))

	if similarity > 1 {
		similarity = 1
	} else if similarity < -1 {
		similarity = -1
	}
	return 1 - similarity
}
