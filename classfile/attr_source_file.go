package classfile

type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (me *SourceFileAttribute) readInfo(reader *ClassReader) {
	me.sourceFileIndex = reader.readUint16()
}

func (me *SourceFileAttribute) FileName() string {
	return me.cp.getUtf8(me.sourceFileIndex)
}
