package rtda

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
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
