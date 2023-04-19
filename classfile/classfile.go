package classfile

import (
	"encoding/binary"
)

const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

var ConstantLengths = map[int]uint16{
	CONSTANT_Class:              2,
	CONSTANT_Fieldref:           4,
	CONSTANT_Methodref:          4,
	CONSTANT_InterfaceMethodref: 4,
	CONSTANT_String:             2,
	CONSTANT_Integer:            4,
	CONSTANT_Float:              4,
	CONSTANT_Long:               8,
	CONSTANT_Double:             8,
	CONSTANT_NameAndType:        4,
	CONSTANT_MethodHandle:       3,
	CONSTANT_MethodType:         2,
	CONSTANT_InvokeDynamic:      4,
}

const magicNumber uint32 = 0xCAFEBABE

var order = binary.BigEndian

type Classfile struct {
	MinorVersion uint16
	MajorVersion uint16
	ConstantPool []*ConstantInfo // uint16 length
	AccessFlags  uint16
	ThisClass    uint16
	SuperClass   uint16
	Interfaces   []uint16         // uint16 length
	Fields       []*MemberInfo    // uint16 length
	Methods      []*MemberInfo    // uint16 length
	Attributes   []*AttributeInfo // uint16 length
}

type ConstantInfo struct {
	Tag  uint8
	Info []byte
}

type MemberInfo struct {
	AccessFlags      uint16
	NameIndex        uint16
	DescriptionIndex uint16
	Attributes       []*AttributeInfo // uint16 length
}

type AttributeInfo struct {
	AttributeNameIndex uint16
	Info               []byte // uint32 length
}
