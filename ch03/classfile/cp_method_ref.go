package classfile

// 10 CONSTANT_Methodref_info { u1 tag; u2 class_index; u2 name_and_type_index; }
//  classIndex       uint16   必须是类
//	nameAndTypeIndex uint16
type ConstantMethodRefInfo struct {
	ConstantMemberRefInfo
}
