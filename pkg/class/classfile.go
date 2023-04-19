package class

import (
	"encoding/binary"
	"fmt"
	"io"
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

const magicNumber = 0xCAFEBABE

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

func Parse(from io.Reader) (*Classfile, error) {
	var magic uint32
	if err := binary.Read(from, order, &magic); err != nil {
		return nil, err
	}
	if magic != magicNumber {
		return nil, fmt.Errorf("unexpected magic number: %d", magic)
	}

	var minor, major uint16
	if err := binary.Read(from, order, &minor); err != nil {
		return nil, err
	}
	if err := binary.Read(from, order, &major); err != nil {
		return nil, err
	}

	constantPool, err := parseConstantPool(from)
	if err != nil {
		return nil, err
	}

	var accessFlags, thisClass, superClass uint16
	if err := binary.Read(from, order, &accessFlags); err != nil {
		return nil, err
	}
	if err := binary.Read(from, order, &thisClass); err != nil {
		return nil, err
	}
	if err := binary.Read(from, order, &superClass); err != nil {
		return nil, err
	}

	var interfacesLength uint16
	if err := binary.Read(from, order, &interfacesLength); err != nil {
		return nil, err
	}
	interfaces := make([]uint16, interfacesLength)
	if err := binary.Read(from, order, interfaces); err != nil {
		return nil, err
	}

	fields, err := parseMembers(from)
	if err != nil {
		return nil, err
	}

	methods, err := parseMembers(from)
	if err != nil {
		return nil, err
	}

	attributes, err := parseAttributes(from)
	if err != nil {
		return nil, err
	}

	return &Classfile{
		MinorVersion: minor,
		MajorVersion: major,
		ConstantPool: constantPool,
		AccessFlags:  accessFlags,
		ThisClass:    thisClass,
		SuperClass:   superClass,
		Interfaces:   interfaces,
		Fields:       fields,
		Methods:      methods,
		Attributes:   attributes,
	}, nil
}

func parseConstantPool(from io.Reader) ([]*ConstantInfo, error) {
	var length uint16
	if err := binary.Read(from, order, &length); err != nil {
		return nil, err
	}
	result := make([]*ConstantInfo, length)
	for i := 0; i < int(length)-1; i++ {
		r, err := parseConstantInfo(from)
		if err != nil {
			return nil, err
		}
		result[i] = r
	}
	return result, nil
}

func parseAttributes(from io.Reader) ([]*AttributeInfo, error) {
	var length uint16
	if err := binary.Read(from, order, &length); err != nil {
		return nil, err
	}
	result := make([]*AttributeInfo, length)
	for i := 0; i < int(length); i++ {
		r, err := parseAttributeInfo(from)
		if err != nil {
			return nil, err
		}
		result[i] = r
	}
	return result, nil
}

func parseMembers(from io.Reader) ([]*MemberInfo, error) {
	var length uint16
	if err := binary.Read(from, order, &length); err != nil {
		return nil, err
	}
	result := make([]*MemberInfo, length)
	for i := 0; i < int(length); i++ {
		r, err := parseMemberInfo(from)
		if err != nil {
			return nil, err
		}
		result[i] = r
	}
	return result, nil
}

func parseConstantInfo(from io.Reader) (*ConstantInfo, error) {
	var tag uint8
	if err := binary.Read(from, order, &tag); err != nil {
		return nil, err
	}
	switch tag {
	case CONSTANT_Utf8:
		var length uint16
		if err := binary.Read(from, order, &length); err != nil {
			return nil, err
		}
		data := make([]byte, length)
		if err := binary.Read(from, order, data); err != nil {
			return nil, err
		}
		lengthBytes := make([]byte, 2)
		data = append(lengthBytes, data...)
		return &ConstantInfo{Tag: tag, Info: data}, nil
	case CONSTANT_Class:
		fallthrough
	case CONSTANT_Double:
		fallthrough
	case CONSTANT_Fieldref:
		fallthrough
	case CONSTANT_Float:
		fallthrough
	case CONSTANT_Integer:
		fallthrough
	case CONSTANT_InterfaceMethodref:
		fallthrough
	case CONSTANT_InvokeDynamic:
		fallthrough
	case CONSTANT_Long:
		fallthrough
	case CONSTANT_MethodHandle:
		fallthrough
	case CONSTANT_MethodType:
		fallthrough
	case CONSTANT_Methodref:
		fallthrough
	case CONSTANT_NameAndType:
		fallthrough
	case CONSTANT_String:
		length := ConstantLengths[int(tag)]
		data := make([]byte, length)
		if err := binary.Read(from, order, data); err != nil {
			return nil, err
		}
		return &ConstantInfo{Tag: tag, Info: data}, nil
	default:
		return nil, fmt.Errorf("unknown constant tag: %d", tag)
	}
}

func parseMemberInfo(from io.Reader) (*MemberInfo, error) {
	var accessFlags, nameIndex, descriptorIndex uint16
	if err := binary.Read(from, order, &accessFlags); err != nil {
		return nil, err
	}
	if err := binary.Read(from, order, &nameIndex); err != nil {
		return nil, err
	}
	if err := binary.Read(from, order, &descriptorIndex); err != nil {
		return nil, err
	}
	attributess, err := parseAttributes(from)
	if err != nil {
		return nil, err
	}
	return &MemberInfo{
		AccessFlags:      accessFlags,
		NameIndex:        nameIndex,
		DescriptionIndex: descriptorIndex,
		Attributes:       attributess,
	}, nil
}

func parseAttributeInfo(from io.Reader) (*AttributeInfo, error) {
	var attributeNameIndex uint16
	if err := binary.Read(from, order, &attributeNameIndex); err != nil {
		return nil, err
	}
	var attributeLength uint32
	if err := binary.Read(from, order, &attributeLength); err != nil {
		return nil, err
	}
	data := make([]byte, attributeLength)
	if err := binary.Read(from, order, data); err != nil {
		return nil, err
	}
	return &AttributeInfo{
		AttributeNameIndex: attributeNameIndex,
		Info:               data,
	}, nil
}
