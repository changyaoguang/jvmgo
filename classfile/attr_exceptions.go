package classfile

// Exceptions method_info
/**
Exceptions_attribute {
	u2 attribute_name_index
	u4 attribute_length
	u2 number_of_exceptions
	u2 exception_index_table[number_of_exceptions]
}
*/
// 变长属性
// 记录方法抛出的异常表
type AttributeExceptions struct {
	exceptionIndexTable []uint16
}

func (self *AttributeExceptions) readInfo(cr *ClassReader) {
	self.exceptionIndexTable = cr.readUint16s()
}

func (self *AttributeExceptions) ExceptionIndexTable() []uint16 {
	return self.exceptionIndexTable
}
