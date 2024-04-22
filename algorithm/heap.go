package main

import "fmt"

type heap struct {
	data []int
}

func new() *heap {
	return &heap{data: make([]int, 0)}
}

func (h *heap) left(i int) int {
	return 2*i + 1
}

func (h *heap) right(i int) int {
	return 2*i + 2
}

func (h *heap) parent(i int) int {
	return (i - 1) / 2
}

func (h *heap) peek() int {
	return h.data[0]
}

func (h *heap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *heap) size() int {
	return len(h.data)
}

func (h *heap) siftUpHeapify(i int) {
	for {
		p := h.parent(i)
		if p < 0 || h.data[i] <= h.data[p] {
			break
		}

		h.swap(i, p)

		i = p
	}
}

func (h *heap) siftDownHeapify(i int) {
	size := len(h.data)

	for {
		l, r, max := h.left(i), h.right(i), i

		if l < size && h.data[l] > h.data[max] {
			max = l
		}

		if r < size && h.data[r] > h.data[max] {
			max = l
		}

		if max == i {
			break
		}

		h.swap(i, max)

		i = max
	}
}

func (h *heap) push(val int) {
	h.data = append(h.data, val)

	h.siftUpHeapify(len(h.data) - 1)
}

func (h *heap) isEmpty() bool {
	return len(h.data) == 0
}

func (h *heap) pop() int {
	if h.isEmpty() {
		return -1
	}

	size := len(h.data)

	h.swap(0, size-1)

	val := h.data[size-1]

	h.data = h.data[:size-1]

	h.siftDownHeapify(0)

	return val
}

func main() {
	h := &heap{data: []int{3, 2, 1}}
	h.push(5)
	h.push(9)
	fmt.Println(h.data)

	h.pop()

	fmt.Println(h.data)
}
