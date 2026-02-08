package hnsw

type HNSW struct {
	EntryPoint *Node
	MaxLevel   int
	M          int
	EfConst    int
	Dist       func(a, b []float32) float32
	Nodes      []*Node
}
