package hnsw

import (
	"container/heap"
	"fmt"
)

type Candidate struct {
	ID       int
	Distance float32
}

// MinHeap definition and operations (closest candidates)
// top element (closest element)
type MinHeap []Candidate

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool {
	return h[i].Distance < h[j].Distance
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Candidate))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)

	item := old[n-1]
	*h = old[:n-1]

	return item
}

// MaxHeap definition
type MaxHeap []Candidate

func (h MaxHeap) Len() int { return len(h) }

func (h MaxHeap) Less(i, j int) bool {
	return h[i].Distance > h[j].Distance
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Candidate))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)

	item := old[n-1]
	*h = old[:n-1]

	return item
}

func (h MaxHeap) Peek() Candidate {
	return h[0]
}

func (h *HNSW) Search(startID int, query []float32, level, ef int) ([]Candidate, error) {
	if ef <= 0 {
		return nil, fmt.Errorf("ef must be positive")
	}

	if startID < 0 || startID >= len(h.Nodes) {
		return nil, fmt.Errorf("start node %d not found", startID)
	}

	visited := make(map[int]struct{}, ef*4)

	candidates := &MinHeap{}
	results := &MaxHeap{}

	heap.Init(candidates)
	heap.Init(results)

	startDist := h.Dist(h.Nodes[startID].Vector, query)

	start := Candidate{
		ID:       startID,
		Distance: startDist,
	}

	heap.Push(candidates, start)
	heap.Push(results, start)

	visited[startID] = struct{}{}

	for candidates.Len() > 0 {
		current := heap.Pop(candidates).(Candidate)

		if results.Len() == 0 {
			continue
		}

		worst := results.Peek()

		if current.Distance > worst.Distance {
			break
		}

		node := h.Nodes[current.ID]

		if level >= len(node.Neighbors) {
			continue
		}

		for _, neighborID := range node.Neighbors[level] {
			if neighborID < 0 || neighborID >= len(h.Nodes) {
				continue
			}

			if _, ok := visited[neighborID]; ok {
				continue
			}

			visited[neighborID] = struct{}{}

			neighborNode := h.Nodes[neighborID]
			dist := h.Dist(neighborNode.Vector, query)

			if results.Len() < ef || dist < results.Peek().Distance {

				c := Candidate{
					ID:       neighborID,
					Distance: dist,
				}

				heap.Push(candidates, c)
				heap.Push(results, c)

				if results.Len() > ef {
					heap.Pop(results)
				}
			}
		}
	}

	res := make([]Candidate, results.Len())

	for i := results.Len() - 1; i >= 0; i-- {
		res[i] = heap.Pop(results).(Candidate)
	}

	return res, nil
}
