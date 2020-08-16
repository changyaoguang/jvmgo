package classfile

// CONSTANT_Class_info { u1 tag; u2 name_index; }
type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16 // CONSTANT_Utf8_info
}

func (self *ConstantClassInfo) readInfo(cr *ClassReader) {
	self.nameIndex = cr.readUint16()
}

func (self ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
