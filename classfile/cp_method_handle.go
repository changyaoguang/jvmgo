package classfile

// 15 CONSTANT_MethodHandle_info { u1 tag; u1 reference_kind; u2 reference_index; }
//  reference_kind       uint16  1-9 方法句柄类型
// 1-4 CONSTANT_Fieldref_info
// 1 REF_getField
// 2 REF_getStatic
// 3 REF_putField
// 4 REF_putStatic

// 5-8 CONSTANT_Methodref_info
// 5 REF_invokeVirtual	 不能为<init> 或 <cinit>
// 6 REF_invokeStatic    不能为<init> 或 <cinit>  版本号>=52.0 可以为CONSTANT_InterfaceMethodref_info
// 7 REF_invokeSpecial   不能为<init> 或 <cinit>  版本号>=52.0 可以为CONSTANT_InterfaceMethodref_info
// 8 REF_newInvokeSpecial  必须是 <init>

// 9 CONSTANT_InterfaceMethodref_info
// 9 REF_invokeInterface  不能为<init> 或 <cinit>

//	reference_index		 uint16
type ConstantMethodHandle struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (self *ConstantMethodHandle) readInfo(cr *ClassReader) {
	self.referenceKind = cr.readUint8()
	self.referenceIndex = cr.readUint16()
}
