package classfile

// utf8 CONSTANT_Utf8_info { u1 tag; u2 length; u1 bytes[length]; }
type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(cr *ClassReader)  {
	strCount := uint32(cr.readUint16())
	data := cr.readBytes(strCount)
	self.str = decodeMUTF8(data)
}

// 简化版
func decodeMUTF8(data []byte) string {
	return string(data)
}