package classfile

import (
	"fmt"
)

/**
ClassFile {
	u4 magic
	u2 minor_version
	u2 major_version
	u2 constant_pool_count
	cp_info constant_pool[constant_pool_count-1]
	u2 access_flags
	u2 this_class
	u2 super_class
	u2 interfaces_count
	u2 interfaces[interfaces_count]
	u2 fields_count
	field_info fields[fields_count]
	u2 methods_count
	method_info methods[methods_count]
	u2 attributes_count
	attribute_info attributes[attributes_count]
}
 */
//解析class文件

type ClassFile struct {
	//magic               uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool //constant_pool_count  cp_info constant_pool[constant_pool_count-1]
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16         //interfaces_count
	fields       []*FieldInfo    //fields_count   field_info
	methods      []*MethodInfo    //methods_count  method_info
	attributes   []AttributeInfo //attributes_count    attribute_info
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (self *ClassFile) read(cr *ClassReader) {
	self.readAndCheckMagic(cr)
	self.readAndCheckVersion(cr)
	self.constantPool = readConstantPool(cr)
	self.accessFlags = cr.readUint16()
	self.thisClass = cr.readUint16()
	self.superClass = cr.readUint16()
	self.interfaces = cr.readUint16s()
	self.fields = readFields(cr, self.constantPool)
	self.methods = readMembers(cr, self.constantPool)
	self.attributes = readAttributes(cr, self.constantPool)
}

func (self *ClassFile) readAndCheckMagic(cr *ClassReader) {
	magic := cr.readUint32()
	fmt.Printf("magic: %x\n",magic)
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError:magic!")
	}
}

func (self *ClassFile) readAndCheckVersion(cr *ClassReader) {
	self.minorVersion = cr.readUint16()
	self.majorVersion = cr.readUint16()
	fmt.Printf("minorVersion: %d\n",self.minorVersion)
	fmt.Printf("majorVersion: %d\n",self.majorVersion)
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)

	}
	return "" // 只有 java.lang.Object没有超类
}
func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, j := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(j)
	}
	return interfaceNames
}

func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}

func (self *ClassFile) Fields() []*FieldInfo {
	return self.fields
}

func (self *ClassFile) Methods() []*MethodInfo {
	return self.methods
}

func (self *ClassFile) Attributes() []AttributeInfo {
	return self.attributes
}
