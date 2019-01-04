package classfile

// ConstantPool 常量池结构体
type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)

	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}

	return cp
}

func (me ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := me[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invaild constant pool index!")
}

func (me ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := me.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := me.getUtf8(ntInfo.nameIndex)
	_type := me.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func (me ConstantPool) getClassName(index uint16) string {
	classInfo := me.getConstantInfo(index).(*ConstantClassInfo)
	return me.getUtf8(classInfo.nameIndex)
}

func (me ConstantPool) getUtf8(index uint16) string {
	utf8Info := me.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
