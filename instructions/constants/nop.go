package constants

import "jvm.go/instructions/base"
import "jvm.go/rtda"

type NOP struct{ base.NoOperandsInstruction }

func (me *NOP) Execute(frame *rtda.frame) {
    // do nothing
}