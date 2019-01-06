package classfile

type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (me *ExceptionsAttribute) readInfo(reader *ClassReader) {
	me.exceptionIndexTable = reader.readUint16s()
}

func (me *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return me.exceptionIndexTable
}
