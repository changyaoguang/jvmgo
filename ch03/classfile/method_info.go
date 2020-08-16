package classfile

/**
method_info {
	u2 access_flags
	u2 name_index
	u2 descriptor_index
	u2 attributes_count
	attribute_info attributes[attributes_count]
}
*/
type MethodInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(cr *ClassReader, cp ConstantPool) []*MethodInfo {
	memberCount := cr.readUint16()
	members := make([]*MethodInfo,memberCount)
	for i,_ := range members {
		members[i] = readMember(cr,cp)
	}
	return members
}

func readMember(cr *ClassReader, cp1 ConstantPool) *MethodInfo {
	return &MethodInfo{
		cp : cp1,
		accessFlags: cr.readUint16(),
		nameIndex: cr.readUint16(),
		descriptorIndex: cr.readUint16(),
		attributes: readAttributes(cr,cp1),
	}
}

func (self *MethodInfo) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *MethodInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)

}
func (self *MethodInfo) Descriptor() string  {
	return self.cp.getUtf8(self.descriptorIndex)

}
