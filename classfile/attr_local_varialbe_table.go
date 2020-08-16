package classfile

// LocalVariableTable			Code
/**
LocalVariableTable_attribute {
	u2 attribute_name_index
	u4 attribute_length
	u2 local_variable_table_length
	{
		u2 start_pc
		u2 length
		u2 name_index
		u2 descriptor_index
		u2 index
	} local_variable_table[local_variable_table_length]
}
*/

type AttributeLocalVariableTable struct {
	localVariableTableEntry []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc         uint16
	length          uint16
	nameIndex       uint16
	descriptorIndex uint16
	index           uint16
}

func (self *AttributeLocalVariableTable) readInfo(cr *ClassReader) {
	localVariableTableLength := cr.readUint16()
	self.localVariableTableEntry = make([]*LocalVariableTableEntry, localVariableTableLength)
	for i := range self.localVariableTableEntry {
		self.localVariableTableEntry[i] = &LocalVariableTableEntry{
			startPc:         cr.readUint16(),
			length:          cr.readUint16(),
			nameIndex:       cr.readUint16(),
			descriptorIndex: cr.readUint16(),
			index:           cr.readUint16(),
		}
	}
}
