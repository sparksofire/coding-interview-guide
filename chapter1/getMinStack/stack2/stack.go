package stack2

import (
	lls "github.com/emirpasic/gods/stacks/linkedliststack"
)

type getMinStack struct {
	stackData *lls.Stack
	stackMin  *lls.Stack
}

func New() *getMinStack {
	return &getMinStack{lls.New(), lls.New()}
}

func (stack *getMinStack) Push(newNum int) {
	stack.stackData.Push(newNum)
	min, ok := stack.GetMin();
	if stack.stackMin.Empty() || ok && min >= newNum {
		stack.stackMin.Push(newNum)
	} else {
		stack.stackMin.Push(min)
	}
}

func (stack *getMinStack) Pop() (value int, ok bool) {
	pop, ok := stack.stackData.Pop()
	if !ok {
		return
	}
	value = pop.(int)
	_, ok = stack.stackMin.Pop()
	return
}

func (stack *getMinStack) GetMin() (min int, ok bool) {
	peek, ok := stack.stackMin.Peek()
	if !ok {
		return
	}
	min = peek.(int)
	return
}
