package classfile

// ConstantUtf8Info 字符串常量结构体
type ConstantUtf8Info struct {
	str string
}

func (me *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	me.str = decodeMUTF8(bytes)
}

func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
