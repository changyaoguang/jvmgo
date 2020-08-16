package classfile

/**
field_info {
	u2 access_flags
	u2 name_index
	u2 descriptor_index
	u2 attributes_count
	attribute_info attributes[attributes_count]
}
*/
type FieldInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readFields(cr *ClassReader, cp ConstantPool) []*FieldInfo {
	fieldCount := cr.readUint16()
	fieldInfos := make([]*FieldInfo, fieldCount)
	for i := range fieldInfos {
		fieldInfos[i] = readField(cr, cp)
	}
	return fieldInfos
}
func readField(cr *ClassReader, cp ConstantPool) *FieldInfo {
	return &FieldInfo{
		cp:              cp,
		accessFlags:     cr.readUint16(),
		nameIndex:       cr.readUint16(),
		descriptorIndex: cr.readUint16(),
		attributes:      readAttributes(cr, cp),
	}
}

func (self *FieldInfo) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *FieldInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

func (self *FieldInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}
