package classfile

type DeprecatedAttribute struct{ MarkerAttribute }
type SyntheticAttribute struct{ MarkerAttribute }

type MarkerAttribute struct{}

func (me *MarkerAttribute) readInfo(reader *ClassReader) {}
