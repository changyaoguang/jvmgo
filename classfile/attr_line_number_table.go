package classfile

// LineNumberTable				Code
/**
LineNumberTable_attribute {
	u2 attribute_name_index
	u4 attribute_length
	u2 line_number_table_length
	{
		u2 start_pc
		u2 line_number
	} line_number_table[line_number_table_length]
}
*/

type AttributeLineNumberTable struct {
	lineNumberTableEntry []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (self *AttributeLineNumberTable) readInfo(cr *ClassReader) {
	lineNumberTableLength := cr.readUint16()
	self.lineNumberTableEntry = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range self.lineNumberTableEntry {
		self.lineNumberTableEntry[i] = &LineNumberTableEntry{
			startPc:    cr.readUint16(),
			lineNumber: cr.readUint16(),
		}
	}
}
