package main

import "fmt"

type FIFO struct {
	queue []interface{}
}

func New() *FIFO {
	return &FIFO{
		queue: make([]interface{}, 0),
	}
}

func (f *FIFO) Push(node interface{}) {
	f.queue = append(f.queue, node)
}

func (f *FIFO) Front() interface{} {
	if len(f.queue) == 0 {
		return nil
	}

	node := f.queue[0]
	f.queue[0] = nil
	f.queue = f.queue[1:]
	return node
}

func main() {
	vals := []int{1, 2, 3, 4}
	fifo := New()

	for _, val := range vals {
		fifo.Push(val)
	}

	for {
		front := fifo.Front()
		if front == nil {
			break
		}
		fmt.Println(front)
	}
}
