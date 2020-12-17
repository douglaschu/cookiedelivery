package main

// imports heap package to enable heap sorting methods
import hp "container/heap"

type path struct {
	value int
	houses []string
}

type minPath []path

func (h minPath) Len() int {
	return len(h)
}

// This function handles the phase of heap sorting that selects the lesser value
func (h minPath) Less(i, j int) bool {
	return h[i].value < h[j].value
}

// This function handles the swap phase of heap sorting
func (h minPath) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minPath) Push(x interface{}) {
	*h = append(*h, x.(path))
}

// This function removes the highest sorted value.
func (h *minPath) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type heap struct {
	values *minPath
}

func newHeap() *heap {
	return &heap{values: &minPath{}}
}

func (h *heap) push(p path) {
	hp.Push(h.values, p)
}

func (h *heap) pop() path {
	i := hp.Pop(h.values)
	return i.(path)
}
