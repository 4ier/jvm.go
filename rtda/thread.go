package rtda

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (me *Thread) PushFrame(frame *Frame) {
	me.stack.push(frame)
}

func (me *Thread) PopFrame() *Frame {
	return me.stack.pop()
}

func (me *Thread) CurrentFrame() *Frame {
	return me.stack.top()
}

func (me *Thread) PC() int {
	return me.pc
}

func (me *Thread) SetPC(pc int) {
	me.pc = pc
}
