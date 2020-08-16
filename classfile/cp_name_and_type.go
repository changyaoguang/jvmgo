package classfile

// 12 CONSTANT_NameAndType_info { u1 tag; u2 name_index; u2 descriptor_index; }
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16 // CONSTANT_Utf8_info
	descriptorIndex uint16 // CONSTANT_Utf8_info
}

func (self *ConstantNameAndTypeInfo) readInfo(cr *ClassReader) {
	self.nameIndex = cr.readUint16()
	self.descriptorIndex = cr.readUint16()
}
