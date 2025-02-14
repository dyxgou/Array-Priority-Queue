package main

type PriorityQueue1 struct {
	N  int
	pq []Item
}

func New(n int) *PriorityQueue1 {
	return &PriorityQueue1{
		N:  0,
		pq: make([]Item, n),
	}
}

// func (p *PriorityQueue) Expired() iter.Seq[Item] {
// 	now := time.Now().Second()
// 	return func(yield func(Item) bool) {
//
// 	}
// }

func (p *PriorityQueue1) Len() int {
	return len(p.pq)
}

func (p *PriorityQueue1) IsEmpty() bool {
	return p.Len() == 0
}

func (p *PriorityQueue1) Insert(val Item) {
	// p.append(val)
	p.pq[p.N] = val
	p.shiftUp(p.N)
	p.N++
}

func (p *PriorityQueue1) Remove() Item {
	val := p.pq[0]
	p.swap(0, p.N-1)
	p.clearLast()
	p.shiftDown(0)
	return val
}

func (p *PriorityQueue1) clearLast() {
	clear(p.pq[p.N-1 : p.N])
	p.N--
}

func (p *PriorityQueue1) less(i, j int) bool {
	return p.pq[i].time < p.pq[j].time
}

func (p *PriorityQueue1) swap(i, j int) {
	p.pq[i], p.pq[j] = p.pq[j], p.pq[i]
}

func (p *PriorityQueue1) parent(i int) int {
	return (i - 1) >> 1
}

func (p *PriorityQueue1) leftChild(i int) int {
	return 2*i + 1
}

func (p *PriorityQueue1) rightChild(i int) int {
	return 2*i + 2
}

func (p PriorityQueue1) shiftUp(cur int) {
	for cur > 0 && p.less(p.parent(cur), cur) {
		parent := p.parent(cur)
		p.swap(parent, cur)
		cur = parent
	}
}

func (p PriorityQueue1) shiftDown(i int) {
	for p.leftChild(i) < p.N-1 {
		l := p.leftChild(i)
		if l < p.N && p.less(l, l+1) {
			l++
		}
		if !p.less(i, l) {
			break
		}

		p.swap(i, l)
		i = l
	}
}

// func (p *PriorityQueue1) String() string {
// 	var sb strings.Builder
//
// 	sb.WriteString("[ ")
// 	for i := 0; i < p.N; i++ {
// 		val := p.pq[i]
// 		sb.WriteString(string(rune(val)))
// 		sb.WriteByte('=')
// 		sb.WriteString(strconv.Itoa(val))
//
// 		if val != 0 {
// 			sb.WriteString(", ")
// 		}
// 	}
// 	sb.WriteString(" ]")
//
// 	return sb.String()
// }
