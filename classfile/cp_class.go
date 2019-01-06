package classfile

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (me *ConstantClassInfo) readInfo(reader *ClassReader) {
	me.nameIndex = reader.readUint16()
}

func (me *ConstantClassInfo) Name() string {
	return me.cp.getUtf8(me.nameIndex)
}
