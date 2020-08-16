package classfile

import "math"

// int
// CONSTANT_Integer_info { u1 tag; u4 bytes; }
type ConstantIntegerInfo struct {
	val int32
}

func (self *ConstantIntegerInfo) readInfo(cr *ClassReader)  {
	bytes := cr.readUint32()
	self.val = int32(bytes)
}

// long CONSTANT_Long_info { u1 tag; u4 high_bytes; u4 low_bytes; }
type ConstantLongInfo struct {
	val int64
}

func (self *ConstantLongInfo) readInfo(cr *ClassReader)  {
	bytes := cr.readUint64()
	self.val = int64(bytes)
}

// float CONSTANT_Float_info { u1 tag; u4 bytes; }
type ConstantFloatInfo struct {
	val float32
}

func (self *ConstantFloatInfo) readInfo(cr *ClassReader)  {
	bytes := cr.readUint32()
	self.val = math.Float32frombits(bytes)
}

//double  CONSTANT_Double_info { u1 tag; u4 high_bytes; u4 low_bytes; }
type ConstantDoubleInfo struct {
	val float64
}

func (self *ConstantDoubleInfo) readInfo(cr *ClassReader)  {
	bytes := cr.readUint64()
	self.val = math.Float64frombits(bytes)
}
