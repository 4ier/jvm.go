package rtda

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (me *Stack) push(frame *Frame) {
	if me.size >= me.maxSize {
		panic("java.lang.StackOverflowError!")
	}

	if me._top != nil {
		frame.lower = me._top
	}

	me._top = frame
	me.size++
}

func (me *Stack) pop() *Frame {
	if me._top == nil {
		panic("jvm stack is empty!")
	}

	top := me._top
	me._top = top.lower
	top.lower = nil
	me.size--
	return top
}

func (me *Stack) top() *Frame {
	if me._top != nil {
		return me._top
	}

	panic("jvm stack is empty!")
}
