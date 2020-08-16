package classfile


// 16 CONSTANT_MethodHandle_info { u1 tag; u2 descriptor_index; }
type ConstantMethodType struct {
	descriptorIndex uint16 //CONSTANT_Utf8_info
}

func (self *ConstantMethodType) readInfo(cr *ClassReader) {
	self.descriptorIndex = cr.readUint16()
}
