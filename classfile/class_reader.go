package classfile

import (
	"encoding/binary"
)

// ClassReader Class文件阅读器
type ClassReader struct {
	data []byte
}

// read u1 data
func (me *ClassReader) readUint8() uint8 {
	val := me.data[0]
	me.data = me.data[1:]
	return val
}

// read u2 data
func (me *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(me.data)
	me.data = me.data[2:]
	return val
}

// read u4 data
func (me *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(me.data)
	me.data = me.data[4:]
	return val
}

// read u8 data
func (me *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(me.data)
	me.data = me.data[8:]
	return val
}

func (me *ClassReader) readUint16s() []uint16 {
	n := me.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = me.readUint16()
	}
	return s
}

func (me *ClassReader) readBytes(length uint32) []byte {
	bytes := me.data[:length]
	me.data = me.data[length:]
	return bytes
}
