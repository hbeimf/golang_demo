package stack

import (
	llstack "github.com/emirpasic/gods/stacks/linkedliststack"
	"fmt"
)

type Min struct {
	top *llstack.Stack
	min_top *llstack.Stack
}

func (self *Min) Push(ele int32) {
	fmt.Println(ele)

	// set val
	self.top.Push(ele)

	// set min
	min_val, err := self.min_top.Pop()
	if err {
		self.min_top.Push(min_val)
		if ele <= min_val.(int32)  {
			self.min_top.Push(ele)
		} else {
			self.min_top.Push(min_val)	
		}
	} else {
		self.min_top.Push(ele)
	}
	
}

func (self *Min) Pop() (int32, int32) {
	val, _ := self.top.Pop()
	min_val, _ := self.min_top.Pop()

	return val.(int32), min_val.(int32)
}



func Test() {
	// stack := llstack.New()  // empty
	// stack.Push(1)       // 1
	// stack.Push(2)       // 1, 2
	// fmt.Println("val", stack.Values())
	// stack.Values()      // 2, 1 (LIFO order)
	// _, _ = stack.Peek() // 2,true
	// _, _ = stack.Pop()  // 2, true
	// _, _ = stack.Pop()  // 1, true
	// _, _ = stack.Pop()  // nil, false (nothing to pop)
	// stack.Push(1)       // 1
	// stack.Clear()       // empty
	// stack.Empty()       // true
	// stack.Size()        // 0

	min := &Min{
		top: llstack.New(),
		min_top:llstack.New(),
	}

	min.Push(1)
	min.Push(2)
	min.Push(3)

	// fmt.Println("stack:", min)

	val, min_val := min.Pop()

	fmt.Println("pop:", val, min_val)

}