package classfile

import "math"

// ConstantIntegerInfo 整型常量结构体
type ConstantIntegerInfo struct {
	val int32
}

func (me *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	me.val = int32(bytes)
}

// ConstantFloatInfo 浮点型常量结构体
type ConstantFloatInfo struct {
	val float32
}

func (me *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	me.val = math.Float32frombits(bytes)
}

type ConstantLongInfo struct {
	val int64
}

func (me *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	me.val = int64(bytes)
}

type ConstantDoubleInfo struct {
	val float64
}

func (me *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	me.val = math.Float64frombits(bytes)
}
