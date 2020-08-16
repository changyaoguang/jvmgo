package classfile

// ConstantValue    		field_info
// ConstantValue_attribute { u2 attribute_name_index; u4 attribute_length; u2 constantvalue_index; }
// attribute_length的值必须是2
// constantvalue_index是常量池索引，但具体指向哪种常量因字段类型而异
// bool byte char short int  CONSTANT_Integer_Info
// float					 CONSTANT_Float_Info
// long					     CONSTANT_Long_Info
// double					 CONSTANT_Double_Info
// String 					 CONSTANT_String_Info
// 定长属性
// 用于表示常量表达式的值
type AttributeConstantValue struct {
	constantValueIndex uint16
}

func (self *AttributeConstantValue) readInfo(cr *ClassReader) {
	self.constantValueIndex = cr.readUint16()
}

func (self *AttributeConstantValue) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}



