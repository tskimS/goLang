package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	v *list.List
}

func (r *Queue) Push(val interface{}) {
	r.v.PushBack(val)
}

func (r *Queue) Pop() interface{} {
	front := r.v.Front()
	if front != nil {
		return r.v.Remove(front)
	}

	return nil
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func testQueue() {
	queue := NewQueue()
	for i := 0; i < 5; i++ {
		queue.Push(i)
	}

	v := queue.Pop()
	for nil != v {
		fmt.Printf("%v -> ", v)
		v = queue.Pop()
	}
}
