package constants

import "jvm.go/instructions/base"
import "jvm.go/rtda"

type BIPUSH struct { val int8 }
type SIPUSH struct { val int16 }

func (me *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
    me.val = reader.ReadInt8()
}

func (me *BIPUSH) Execute(frame *rtda.Frame) {
    i := int32(me.val)
    frame.OperandStack().PushInt(i)
}

