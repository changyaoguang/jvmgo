package classfile

import "fmt"

const (
	CONSTANT_Utf8               = 1
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_Class              = 7
	CONSTANT_String             = 8
	CONSTANT_FeildRef           = 9
	CONSTANT_MethodRef          = 10
	CONSTANT_InterfaceMethodRef = 11
	CONSTANT_NameAndType        = 12
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

type ConstantInfo interface {
	readInfo(cr *ClassReader)
}

func readConstantInfo(cr *ClassReader, cp ConstantPool) ConstantInfo {
	tag := cr.readUint8()
	fmt.Printf("tag1: %d\n",tag)
	c := newConstantInfo(tag, cp)
	c.readInfo(cr)
	return c
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_Class:
		return &ConstantClassInfo{cp: cp}
	case CONSTANT_String:
		return &ConstantStringInfo{cp: cp}
	case CONSTANT_FeildRef:
		return &ConstantFieldRefInfo{ConstantMemberRefInfo{cp: cp}}
	case CONSTANT_MethodRef:
		return &ConstantMethodRefInfo{ConstantMemberRefInfo{cp: cp}}
	case CONSTANT_InterfaceMethodRef:
		return &ConstantInterfaceMethodRefInfo{ConstantMemberRefInfo{cp: cp}}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandle{}
	case CONSTANT_MethodType:
		return &ConstantMethodType{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamic{}
	default:
		fmt.Printf("tag: %d\n",tag)
		panic("java.lang.ClassFormatError: constant pool tag!")
	}
}
