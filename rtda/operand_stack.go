package rtda

import "math"

type OperandStack struct {
	size  uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

func (me *OperandStack) PushInt(val int32) {
	me.slots[me.size].num = val
	me.size++
}

func (me *OperandStack) PopInt() int32 {
	me.size--
	return me.slots[me.size].num
}

func (me *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	me.PushInt(int32(bits))
}

func (me *OperandStack) PopFloat() float32 {
	return math.Float32frombits(uint32(me.PopInt()))
}

func (me *OperandStack) PushLong(val int64) {
	me.slots[me.size].num = int32(val)
	me.slots[me.size+1].num = int32(val >> 32)
	me.size += 2
}

func (me *OperandStack) PopLong() int64 {
	me.size -= 2
	low := uint32(me.slots[me.size].num)
	high := uint32(me.slots[me.size+1].num)
	return int64(high)<<32 | int64(low)
}

func (me *OperandStack) PushDouble(val float64) {
	me.PushLong(int64(math.Float64bits(val)))
}

func (me *OperandStack) PopDouble() float64 {
	return math.Float64frombits(uint64(me.PopLong()))
}

func (me *OperandStack) PushRef(ref *Object) {
	me.slots[me.size].ref = ref
	me.size++
}

func (me *OperandStack) PopRef() *Object {
	me.size--
	ref := me.slots[me.size].ref
	// 帮助go语言的垃圾回收当前引用
	me.slots[me.size].ref = nil
	return ref
}
