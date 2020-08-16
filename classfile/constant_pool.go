package classfile


type ConstantPool []ConstantInfo

func readConstantPool(cr *ClassReader) ConstantPool {
	cpCount := int(cr.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(cr, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo,*ConstantDoubleInfo:
			i++;
		}
	}
	return cp
}

func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index];cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	nameAndTypeInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	_name := self.getUtf8(nameAndTypeInfo.nameIndex)
	_type := self.getUtf8(nameAndTypeInfo.descriptorIndex)
	return _name,_type
}

func (self ConstantPool) getClassName(index uint16) string {
	classNameInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	_name := self.getUtf8(classNameInfo.nameIndex)
	return _name
}
func (self ConstantPool) getUtf8(index uint16) string {
	utf8str := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8str.str
}
