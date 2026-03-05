package hnsw

type HNSW struct {
	EntryPoint int
	MaxLevel   int

	M              int
	EfConstruction int
	EfSearch       int

	Dist func(a, b []float32) float32

	Nodes []Node
}

// NewHNSW - contructor
func NewHNSW(m, efConstruction, efSearch int, dist func(a, b []float32) float32) *HNSW {
	return &HNSW{
		EntryPoint: -1,
		MaxLevel:   -1,

		M:              m,
		EfConstruction: efConstruction,
		EfSearch:       efSearch,

		Dist: dist,

		Nodes: make([]Node, 0),
	}
}
