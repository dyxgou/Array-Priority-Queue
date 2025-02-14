package main

import (
	"container/heap"
	"slices"
	"testing"
)

var n = 1024

func TestInsertAndRemoveOne(t *testing.T) {
	tt := struct {
		key  string
		time int64
	}{
		key:  "first",
		time: 1,
	}
	pq := New(1)

	pq.Insert(NewItem(tt.key, tt.time))

	m := pq.Remove()

	if m.key != tt.key {
		t.Errorf("val key expected=%q. got=%q", tt.key, m.key)
	}

	if m.time != tt.time {
		t.Errorf("time expected=%d. got=%d", tt.time, m.time)
	}
}

func TestInsertAndRemove(t *testing.T) {
	tests := []struct {
		key  string
		time int64
	}{
		{
			key:  "first",
			time: 1,
		},
		{
			key:  "second",
			time: 2,
		},
		{
			key:  "third",
			time: 3,
		},
		{
			key:  "four",
			time: 4,
		},
		{
			key:  "five",
			time: 5,
		},
		{
			key:  "six",
			time: 6,
		},
	}

	pq := New(len(tests))

	for _, item := range tests {
		pq.Insert(NewItem(item.key, item.time))
	}

	for _, tt := range slices.Backward(tests) {
		m := pq.Remove()

		if m.key != tt.key {
			t.Errorf("val key expected=%q. got=%q", tt.key, m.key)
		}

		if m.time != tt.time {
			t.Errorf("time expected=%d. got=%d", tt.time, m.time)
		}
	}
}

func BenchmarkPQNative(b *testing.B) {
	pq := New(b.N)
	b.ResetTimer()

	for i := range b.N {
		pq.Insert(NewItem("key", int64(i)))
	}

	b.Cleanup(func() {
		pq = nil
	})
}

func BenchmarkPQInter(b *testing.B) {
	h := make(PriorityQueue, 0)
	heap.Init(&h)
	b.ResetTimer()

	for i := range b.N {
		ni := NewItem("key", int64(i))
		heap.Push(&h, &ni)
	}

	b.Cleanup(func() {
		h = nil
	})
}
