package classfile

// 9 CONSTANT_Fieldref_info { u1 tag; u2 class_index; u2 name_and_type_index; }
//  classIndex       uint16   类或接口
//	nameAndTypeIndex uint16
type ConstantFieldRefInfo struct {
	ConstantMemberRefInfo
}