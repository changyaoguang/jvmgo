package classfile

// SourceFile	ClassFile
// SourceFile_attribute { u2 attribute_name_index; u4 attribute_length; u2 sourcefile_index; }
// 可选定长属性
// 用于指出源文件名
type AttributeSourceFile struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (self *AttributeSourceFile) readInfo(cr *ClassReader) {
	self.sourceFileIndex = cr.readUint16()
}

func (self *AttributeSourceFile) FileName() string {
	return self.cp.getUtf8(self.sourceFileIndex)
}
