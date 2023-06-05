package queue

type TaskQueue struct {
	capacity int
	array    []int64

	head int
	tail int
	len  int
}

func NewTaskQueue(capacity int) *TaskQueue {
	return &TaskQueue{
		capacity: capacity,
		array:    make([]int64, capacity),
	}
}

func (q *TaskQueue) Push(task int64) {
	if q.len == q.capacity {
		return
	}
	q.array[q.tail] = task
	q.tail = (q.tail + 1) % q.capacity
	q.len++
}

func (q *TaskQueue) Pop() {
	if q.len == 0 {
		return
	}
	q.head = (q.head + 1) % q.capacity
	q.len--
}

func (q *TaskQueue) Front() int64 {
	if q.len == 0 {
		return -1
	}
	return q.array[q.head]
}

func (q *TaskQueue) Len() int {
	return q.len
}
