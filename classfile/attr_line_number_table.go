package classfile

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (me *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	me.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range me.lineNumberTable {
		me.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
