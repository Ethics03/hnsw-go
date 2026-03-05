package hnsw

func CosineDistance(a, b []float32) float32 {
	aNorm := Normalize(a)
	bNorm := Normalize(b)

	var dot float32
	for i := range len(aNorm) {
		dot += aNorm[i] * bNorm[i]
	}

	similarity := dot

	if similarity > 1 {
		similarity = 1
	} else if similarity < -1 {
		similarity = -1
	}
	return 1 - similarity
}
