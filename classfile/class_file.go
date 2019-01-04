package classfile

import "fmt"

// ClassFile 类文件
type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlag   uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttrbuteInfo
}

func Parse(classData []byte) (cf *ClassFile, err erros) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (me *ClassFile) read(reader *ClassReader) {
	me.readAndCheckMagic(reader)
	me.readAndCheckVersion(reader)
	me.constantPool = readConstantPool(reader)
	me.accessFlag = reader.readUint16()
	me.thisClass = reader.readUint16()
	me.superClass = reader.readUint16()
	me.interfaces = reader.readUint16s()
	me.fields = readMembers(reader, me.constantPool)
	me.methods = readMembers(reader, me.constantPool)
	me.attributes = readAttributes(reader, me.constantPool)
}

func (me *ClassFile) MinorVersion() uint16       { return me.minorVersion }
func (me *ClassFile) MajorVersion() uint16       { return me.majorVersion }
func (me *ClassFile) ConstantPool() ConstantPool { return me.constantPool }
func (me *ClassFile) AccessFlag() uint16         { return me.accessFlag }
func (me *ClassFile) ThisClass() uint16          { return me.thisClass }
func (me *ClassFile) SuperClass() uint16         { return me.superClass }
func (me *ClassFile) Interfaces() []uint16       { return me.interfaces }
func (me *ClassFile) Fields() []*MemberInfo      { return me.fields }
func (me *ClassFile) Methods() []*MemberInfo     { return me.methods }
func (me *ClassFile) Attributes() []AttrbuteInfo { return me.attributes }

func (me *ClassFile) ClassName() string {
	return me.constantPool.getClassName(me.thisClass)
}

func (me *ClassFile) SuperClassName() string {
	if me.superClass > 0 {
		return me.constantPool.getClassName(me.SuperClass)
	}
	return ""
}

func (me *ClassFile) InterfaceName() []string {
	interfaceNames := make([]string, len(me.interfaces))
	for i, cpIndex := range me.interfaces {
		interfaceNames[i] = me.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

func (me *ClassFile) readAndCheckMagic(reader *ClassReader) {
	me.magic = reader.readUint32()
	if me.magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

func (me *ClassFile) readAndCheckVersion(reader *ClassReader) {
	me.minorVersion = reader.readUint16()
	me.majorVersion = reader.readUint16()
	switch me.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if me.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}
