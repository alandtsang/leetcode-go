package minstack

import "github.com/alandtsang/leetcode-go/stack"

type Stack = stack.Stack

type MinStack struct {
	min  *Stack
	data *Stack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	minStack := stack.NewStack()
	dataStack := stack.NewStack()
	return MinStack{
		min:  minStack,
		data: dataStack,
	}
}

func (this *MinStack) Push(x int) {
	this.data.Push(x)
	if this.min.Len() == 0 {
		this.min.Push(x)
	} else {
		val := this.GetMin()
		if x < val {
			this.min.Push(x)
		} else {
			this.min.Push(val)
		}
	}
}

func (this *MinStack) Pop() {
	this.data.Pop()
	this.min.Pop()
}

func (this *MinStack) Top() int {
	top := this.data.Peek()
	return top.(int)
}

func (this *MinStack) GetMin() int {
	m := this.min.Peek()
	return m.(int)
}
