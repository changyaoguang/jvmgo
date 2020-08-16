package classfile

// 18 CONSTANT_InterfaceMethodref_info { u1 tag; u2 bootstrap_method_attr_index; u2 name_and_type_index; }
//bootstrap_method_attr_index
// name_and_type_index
type ConstantInvokeDynamic struct {
	bootstrapMethodAttrIndex uint16 //
	nameAndTypeIndex         uint16 //CONSTANT_NameAndType_info
}

func (self *ConstantInvokeDynamic) readInfo(cr *ClassReader) {
	self.bootstrapMethodAttrIndex = cr.readUint16()
	self.nameAndTypeIndex = cr.readUint16()
}
