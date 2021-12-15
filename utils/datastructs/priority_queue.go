package datastructs

import (
	"container/heap"
	"fmt"
)

type PriorityQueue []*PQItem
type PQItem struct {
	Value    interface{}
	Priority int
	Index    int
}

func (pq PriorityQueue) Init() {
	heap.Init(&pq)
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*PQItem)
	item.Index = n
	*pq = append(*pq, item)
}

// Min priority mode
func (pq *PriorityQueue) MinPop() interface{} {
	old := *pq
	n := len(old)
	item := old[0]
	old[0] = nil
	item.Index = -1
	*pq = old[1:n]
	return item
}

// Max priority mode
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) Update(item *PQItem, value interface{}, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}

func (pq PriorityQueue) Print() {
	for pq.Len() > 0 {
		item := pq.Pop().(*PQItem)
		val := item.Value
		fmt.Printf("%d:%v ", item.Priority, val)
	}
}
