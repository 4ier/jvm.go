package rtda

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
    thread       *Thread
    nextPC       int
}

func NewFrame(maxLocals uint, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (me *Frame) LocalVars() LocalVars {
	return me.localVars
}

func (me *Frame) OperandStack() *OperandStack {
	return me.operandStack
}

func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
    return &Frame{
        thread:         thread,
        localVars:      newLocalVars(maxLocals),
        operandStack:   newOperandStack(maxStack),
    }
}