package classfile

type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (me *UnparsedAttribute) readInfo(reader *ClassReader) {
	me.info = reader.readBytes(me.length)
}
