package customcontainer

type PriorityElement struct {
	Prioty any
	Value  any
}

type PriorityQueue struct {
	Queue      []PriorityElement
	Comparator func(PriorityElement, PriorityElement) bool
}

func NewPriorityQueue(size int, comparator func(PriorityElement, PriorityElement) bool) *PriorityQueue {
	return &PriorityQueue{Queue: make([]PriorityElement, 0, size), Comparator: comparator}
}

func (pq *PriorityQueue) Len() int { return len(pq.Queue) }

func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.Comparator(pq.Queue[i], pq.Queue[j])
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.Queue[i], pq.Queue[j] = pq.Queue[j], pq.Queue[i]
}

func (pq *PriorityQueue) Push(ele any) {
	pq.Queue = append(pq.Queue, ele.(PriorityElement))
}

func (pq *PriorityQueue) Pop() any {
	deleted_ele := pq.Queue[len(pq.Queue)-1]
	pq.Queue = pq.Queue[:len(pq.Queue)-1]
	return deleted_ele

}
