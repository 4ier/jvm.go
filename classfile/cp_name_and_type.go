package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (me *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	me.nameIndex = reader.readUint16()
	me.descriptorIndex = reader.readUint16()
}
