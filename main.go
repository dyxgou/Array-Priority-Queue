package main

import (
	"container/heap"
	"fmt"
	"iter"
	"time"
)

type timestamp struct {
	key  string
	time time.Time
}

func newTimestamp(k string, t int64) timestamp {
	return timestamp{
		key:  k,
		time: time.Now().Add(time.Duration(t) * time.Second),
	}
}

type times []timestamp

func (t times) Len() int {
	return len(t)
}

func (t times) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t times) Less(i, j int) bool {
	return t[i].time.Before(t[j].time)
}

func (t times) Timestamps() iter.Seq[timestamp] {
	now := time.Now()
	fmt.Printf("now=%s\n", now)
	return func(yield func(timestamp) bool) {
		for _, ts := range t {
			fmt.Printf("cur=%s\n", ts.time)
			if now.Before(ts.time) {
				return
			}

			if !yield(ts) {
				return
			}
		}
	}
}

func (t *times) Push(x any) {
	ts, ok := x.(timestamp)
	if !ok {
		panic("argument given in push should be a timestamp")
	}

	*t = append(*t, ts)
}

func (t *times) Pop() any {
	old := *t
	n := len(old)
	x := old[n-1]
	*t = old[0 : n-1]

	return x
}

func main() {
	t := &times{}
	heap.Init(t)
	heap.Push(t, newTimestamp("key 1", 2))
	heap.Push(t, newTimestamp("key 2", 10))
	heap.Push(t, newTimestamp("key 3", 3))
	heap.Push(t, newTimestamp("key 4", 4))
	time.Sleep(3 * time.Second)

	for ts := range t.Timestamps() {
		fmt.Printf("time=%+v\n", ts)
	}
}
