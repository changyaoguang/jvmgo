package classfile

// string CONSTANT_String_info { u1 tag; u2 string_index; }
type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16 //CONSTANT_Utf8_info
}

func (self *ConstantStringInfo) readInfo(cr *ClassReader) {
	self.stringIndex = cr.readUint16()
}

func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}
