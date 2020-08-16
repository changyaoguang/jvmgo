package classfile

// Synthetic	 ClassFile, field_info, method_info	--Synthetic_attribute { u2 attribute_name_index; u4 attribute_length; }
// Deprecated	 ClassFile, field_info, method_info  -- Deprecated_attribute { u2 attribute_name_index; u4 attribute_length; }
type AttributeDeprecated struct {
	AttributeMarker
}
