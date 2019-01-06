package classfile

type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (me *CodeAttribute) readInfo(reader *ClassReader) {
	me.maxStack = reader.readUint16()
	me.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	me.code = reader.readBytes(codeLength)
	me.exceptionTable = readExceptionTable(reader)
	me.attributes = readAttributes(reader, me.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}
