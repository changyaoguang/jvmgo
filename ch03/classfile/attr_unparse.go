package classfile

type AttributeUnParse struct {
	name   string
	length uint32
	info   []byte
}

func (self *AttributeUnParse) readInfo(cr *ClassReader) {
	self.info = cr.readBytes(self.length)
}
