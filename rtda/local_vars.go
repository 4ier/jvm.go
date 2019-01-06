package rtda

import "math"

type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (me LocalVars) SetInt(index uint, val int32) {
	me[index].num = val
}

func (me LocalVars) GetInt(index uint) int32 {
	return me[index].num
}

func (me LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	me[index].num = int32(bits)
}

func (me LocalVars) GetFloat(index uint) float32 {
	bits := uint32(me[index].num)
	return math.Float32frombits(bits)
}

func (me LocalVars) SetLong(index uint, val int64) {
	me[index].num = int32(val)
	me[index+1].num = int32(val >> 32)
}

func (me LocalVars) GetLong(index uint) int64 {
	low := uint32(me[index].num)
	high := uint32(me[index+1].num)
	return int64(high)<<32 | int64(low)
}

func (me LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	me.SetLong(index, int64(bits))
}

func (me LocalVars) GetDouble(index uint) float64 {
	bits := uint64(me.GetLong(index))
	return math.Float64frombits(bits)
}

func (me LocalVars) SetRef(index uint, ref *Object) {
	me[index].ref = ref
}

func (me LocalVars) GetRef(index uint) *Object {
	return me[index].ref
}
