package main

import (
	"slices"
	"strconv"
	"strings"
)

type PriorityQueue []int

func NewPQ() PriorityQueue {
	return make(PriorityQueue, 0, 1024)
}

func (p *PriorityQueue) Insert(val int) {
	*p = append(*p, val)
	p.shiftUp(p.Len() - 1)
}

func (p PriorityQueue) Remove() int {
	m := p[0]
	p.swap(0, p.Len()-1)
	p.removeLast()
	p.shiftDown(0)
	return m
}

func (p PriorityQueue) less(i, j int) bool {
	return p[i] > p[j]
}

func (p *PriorityQueue) removeLast() {
	*p = slices.Delete(*p, p.Len()-1, p.Len())
}

func (p PriorityQueue) swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorityQueue) parent(i int) int {
	return (i - 1) >> 1
}

func (p *PriorityQueue) leftChild(i int) int {
	return 2*i + 1
}

func (p *PriorityQueue) rightChild(i int) int {
	return 2*i + 2
}

func (p *PriorityQueue) isValidIndex(i int) bool {
	return i > 0 && i < p.Len()
}

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) shiftUp(cur int) {
	for cur > 0 && p.less(p.parent(cur), cur) {
		parent := p.parent(cur)
		p.swap(parent, cur)
		cur = parent
	}
}

func (p PriorityQueue) shiftDown(i int) {
	for p.leftChild(i) < p.Len()-1 {
		l := p.leftChild(i)
		if l < p.Len() && p.less(l, l+1) {
			l++
		}
		if !p.less(i, l) {
			break
		}

		p.swap(i, l)
		i = l
	}
}

func (p PriorityQueue) String() string {
	var sb strings.Builder

	for i, val := range p {
		sb.WriteString(string(rune(val)))
		sb.WriteByte('=')
		sb.WriteString(strconv.Itoa(val))
		if i != len(p)-1 {
			sb.WriteString(", ")
		}
	}

	return sb.String()
}
