package classfile

// Code	 method_info
/**
Code_attribute {
	u2 attribute_name_index
	u4 attribute_length
	u2 max_stack  // 操作数栈的最大深度
	u2 max_locals // 局部变量表大小
	u4 code_length
	u1 code[code_length] // 字节码
	u2 exception_table_length
	{
		u2 start_pc
		u2 end_pc
		u2 handler_pc
		u2 catch_type
	} exception_table[exception_table_length]
	u2 attributes_count
	attribute_info attributes[attributes_count]
}
*/

type AttributeCode struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlePc  uint16
	catchType uint16
}

func (self *AttributeCode) readInfo(cr *ClassReader) {
	self.maxStack = cr.readUint16()
	self.maxLocals = cr.readUint16()
	codeLength := cr.readUint32()
	self.code = cr.readBytes(codeLength)
	self.exceptionTable = readExceptionTable(cr)
	self.attributes = readAttributes(cr, self.cp)
}

func readExceptionTable(cr *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := cr.readUint16()
	exceptionTableEntry := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTableEntry {
		exceptionTableEntry[i] = &ExceptionTableEntry{
			startPc:   cr.readUint16(),
			endPc:     cr.readUint16(),
			handlePc:  cr.readUint16(),
			catchType: cr.readUint16(),
		}
	}
	return exceptionTableEntry
}
