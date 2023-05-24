package queue

type Task chan struct{}

type TaskQueue struct {
	capacity int
	array    []Task

	head int
	tail int
	len  int
}

func NewTaskQueue(capacity int) *TaskQueue {
	return &TaskQueue{
		capacity: capacity,
		array:    make([]Task, capacity),
	}
}

func (q *TaskQueue) Push(task Task) {
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

func (q *TaskQueue) Front() Task {
	if q.len == 0 {
		return nil
	}
	return q.array[q.head]
}

func (q *TaskQueue) Len() int {
	return q.len
}
