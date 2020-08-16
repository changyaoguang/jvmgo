package classfile

// CONSTANT_Fieldref_info { u1 tag; u2 class_index; u2 name_and_type_index; }
type ConstantMemberRefInfo struct {
	cp               ConstantPool
	classIndex       uint16  // CONSTANT_Class_info
	nameAndTypeIndex uint16 // CONSTANT_NameAndType_info
}

func (self *ConstantMemberRefInfo) readInfo(cr *ClassReader) {
	self.classIndex = cr.readUint16()
	self.nameAndTypeIndex = cr.readUint16()
}

func (self *ConstantMemberRefInfo) ClassName(cr *ClassReader) string {
	return self.cp.getClassName(self.classIndex)
}

func (self *ConstantMemberRefInfo) NameAndDescriptor(cr *ClassReader) (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}
