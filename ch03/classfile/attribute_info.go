package classfile

// attribute_info { u2 attribute_name_index; u4 attribute_length; u1 info[attribute_length]; }
type AttributeInfo interface {
	readInfo(cr *ClassReader)
}

func readAttributes(cr *ClassReader, cp ConstantPool) []AttributeInfo {
	attributeCount := cr.readUint16()
	attributes := make([]AttributeInfo, attributeCount)
	for i := range attributes {
		attributes[i] = readAttribute(cr, cp)
	}
	return attributes
}

func readAttribute(cr *ClassReader, cp ConstantPool) AttributeInfo {
	attributeNameIndex := cr.readUint16() // u2 attribute_name_index
	attributeName := cp.getUtf8(attributeNameIndex)
	attributeLength := cr.readUint32() // u4 attribute_length
	attributeInfo := newAttributeInfo(attributeName, attributeLength, cp)
	attributeInfo.readInfo(cr)
	return attributeInfo
}

// 第一组属性是实现 Java虚拟机所必需的，共有5种；
// ConstantValue    		field_info --ConstantValue_attribute { u2 attribute_name_index; u4 attribute_length; u2 constantvalue_index; }
// Code						method_info
// Exceptions				method_info --Exceptions_attribute { u2 attribute_name_index; u4 attribute_length; u2 number_of_exceptions; u2 exception_index_table[number_of_exceptions]; }
// StackMapTable			Code
// BootstrapMethods			ClassFile

// 第二组属性是Java类库所必需的， 共有12种；
// InnerClasses								ClassFile
// EnclosingMethod							ClassFile
// Synthetic								ClassFile, field_info, method_info	--Synthetic_attribute { u2 attribute_name_index; u4 attribute_length; }
// Signature								ClassFile, field_info, method_info
// RuntimeVisibleAnnotations				ClassFile, field_info, method_info
// RuntimeInvisibleAnnotations				ClassFile, field_info, method_info
// RuntimeVisibleTypeAnnotations			ClassFile, field_info, method_info
// RuntimeInvisibleTypeAnnotations			ClassFile, field_info, method_info
// RuntimeVisibleParameterAnnotations		method_info
// RuntimeInvisibleParameterAnnotations		method_info
// AnnotationDefault						method_info
// MethodParameters							method_info

// 第三组属性主要提供给工具使用，共有6种。
// SourceFile					ClassFile	--SourceFile_attribute { u2 attribute_name_index; u4 attribute_length; u2 sourcefile_index; }
// SourceDebugExtension			ClassFile
// LineNumberTable				Code
// LocalVariableTable			Code
// LocalVariableTypeTable		Code
// Deprecated					ClassFile, field_info, method_info  -- Deprecated_attribute { u2 attribute_name_index; u4 attribute_length; }

// 第三组属性 是可选的，也就是说可以不出现在class文件中。
// 如果class文件中存在第三组属性，Java虚拟机实现或者Java类库也是可以利用它们的，比如使用LineNumberTable属性在异常堆栈中显示行号。
func newAttributeInfo(attributeName string, attributeLength uint32, cp ConstantPool) AttributeInfo {
	switch attributeName {
	case "ConstantValue":
		return &AttributeConstantValue{}
	case "Code":
		return &AttributeCode{cp: cp}
	case "Exceptions":
		return &AttributeExceptions{}
	case "Deprecated":
		return &AttributeDeprecated{}
	case "Synthetic":
		return &AttributeSynthetic{}
	case "LineNumberTable":
		return &AttributeLineNumberTable{}
	case "LocalVariableTable":
		return &AttributeLocalVariableTable{}
	case "SourceFile":
		return &AttributeSourceFile{cp: cp}
	default:
		return &AttributeUnParse{attributeName, attributeLength, nil}
	}
}
