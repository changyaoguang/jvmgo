package classfile

// 11 CONSTANT_InterfaceMethodref_info { u1 tag; u2 class_index; u2 name_and_type_index; }
//  classIndex       uint16   必须是接口
//	nameAndTypeIndex uint16
type ConstantInterfaceMethodRefInfo struct {
	ConstantMemberRefInfo
}
