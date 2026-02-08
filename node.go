package hnsw

type Node struct {
	ID        int
	Vector    []float32
	Level     int
	Neighbors map[int][]*Node // level -> neigbor
}

