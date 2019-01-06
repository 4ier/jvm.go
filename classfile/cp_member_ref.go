package classfile

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (me *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	me.classIndex = reader.readUint16()
	me.nameAndTypeIndex = reader.readUint16()
}

func (me *ConstantMemberrefInfo) ClassName() string {
	return me.cp.getClassName(me.classIndex)
}

func (me *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return me.cp.getNameAndType(me.nameAndTypeIndex)
}

type ConstantFieldrefInfo struct{ ConstantMemberrefInfo }
type ConstantMethodrefInfo struct{ ConstantMemberrefInfo }
type ConstantInterfaceMethodrefInfo struct{ ConstantMemberrefInfo }
