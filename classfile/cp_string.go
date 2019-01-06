package classfile

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (me *ConstantStringInfo) readInfo(reader *ClassReader) {
	me.stringIndex = reader.readUint16()
}

func (me *ConstantStringInfo) String() string {
	return me.cp.getUtf8(me.stringIndex)
}
