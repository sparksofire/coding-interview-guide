package stack1

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
	if stack.stackMin.Empty() {
		stack.stackMin.Push(newNum)
	} else if min, ok := stack.GetMin(); ok && min >= newNum  {
		stack.stackMin.Push(newNum)
	}
}

func (stack *getMinStack) Pop() (value int, ok bool) {
	pop, ok := stack.stackData.Pop()
	if !ok {
		return
	}
	value = pop.(int)
	if getMin, ok := stack.GetMin();ok && value == getMin {
		stack.stackMin.Pop()
	}
	return
}

func (stack *getMinStack) GetMin() (min int, ok bool) {
	peek, ok := stack.stackMin.Peek()
	min = peek.(int)
	return
}
