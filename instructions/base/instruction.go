package base

import "jvm.go/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {}

func (me *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// do nothing
}

type BranchInstruction struct {
	Offset int
}

func (me *BranchInstruction) FetchOperands(reader *BytecodeReader) {
    me.Offset = int(reader.ReadInt16())
}

type Index8Instruction struct {
	Index uint
}

func (me *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	me.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct{
    Index uint
}

func (me *Index16Instruction) FetchOperands(reader *BytecodeReader){
    me.Index = uint(reader.ReadInt16())
}

