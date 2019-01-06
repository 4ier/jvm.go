package classfile

type ConstantValueAttribute struct {
	constantValueAttribute uint16
}

func (me *ConstantValueAttribute) readInfo(reader *ClassReader) {
	me.constantValueAttribute = reader.readUint16()
}

func (me *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return me.constantValueAttribute
}
