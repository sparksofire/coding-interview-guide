package twoStacksQueue

import (
	lls "github.com/emirpasic/gods/stacks/linkedliststack"
)

type queue struct {
	stackPush *lls.Stack
	stackPop *lls.Stack
}

func New() *queue {
	return &queue{lls.New(), lls.New()}
}

func (q *queue) Add(value int) {
	q.stackPush.Push(value)
}

func (q *queue) Poll() (int, bool) {
	if q.stackPush.Empty() && q.stackPop.Empty() {
		return 0, false
	} else if q.stackPop.Empty() {
		for {
			if value, ok := q.stackPush.Pop(); ok {
				q.stackPop.Push(value)
			} else {
				break
			}
		}
	}
	value, ok := q.stackPop.Pop()
	if !ok {
		return 0, false
	}
	return value.(int), ok
}

func (q *queue) Peek() (int, bool) {
	if q.stackPush.Empty() && q.stackPop.Empty() {
		return 0, false
	} else if q.stackPop.Empty() {
		for {
			if value, ok := q.stackPush.Pop(); ok {
				q.stackPop.Push(value)
			} else {
				break
			}
		}
	}
	value, ok := q.stackPop.Peek()
	if !ok {
		return 0, false
	}
	return value.(int), ok
}
