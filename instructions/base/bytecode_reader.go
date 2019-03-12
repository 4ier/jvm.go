package base

type BytecodeReader struct{
    code []byte
    pc int
}

func (me *BytecodeReader) Reset(code []byte, pc int){
    me.code = code
    me.pc = pc
}

func (me *BytecodeReader) ReadUint8() uint8 {
    i:= me.code[me.pc]
    me.pc++
    return i
}

func (me *BytecodeReader) ReadInt8() int8 {
    return int8(me.ReadUint8())
}

func (me *BytecodeReader) ReadUint16() uint16{
    byte1:= uint16(me.ReadUint8())
    byte2:= uint16(me.ReadUint8())
    return (byte1 << 8) | byte2
}

func (me *BytecodeReader) ReadInt16() int16{
    return int16(me.ReadUint16())
}

func (me *BytecodeReader) ReadInt32() int32{
    byte1:= int32(me.ReadUint8())
    byte2:= int32(me.ReadUint8())
    byte3:= int32(me.ReadUint8())
    byte4:= int32(me.ReadUint8())
    return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

