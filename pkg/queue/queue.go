package queue

import (
	"container/list"
	"errors"
)

type Queue struct {
	list.List
}

func (q *Queue) Enqueue(elem interface{}) error {
	if elem == nil {
		err := errors.New("Enqueued element is nil")
		return err
	}
	e := list.Element{Value: elem}
	q.List.PushBack(e)
	return nil
}

func (q *Queue) Dequeue() (interface{}, error) {
	elem := q.Front()
	if elem.Value == nil {
		err := errors.New("Nothing to deque")
		return elem.Value, err
	}
	e := q.List.Remove(elem)
	return e, nil
}
