package classfile

import (
	"encoding/binary"
	"io"
)

func Write(class *Classfile, to io.Writer) error {
	if err := binary.Write(to, Order, magicNumber); err != nil {
		return err
	}
	if err := binary.Write(to, Order, class.MinorVersion); err != nil {
		return err
	}
	if err := binary.Write(to, Order, class.MajorVersion); err != nil {
		return err
	}
	if err := writeConstantPool(to, class.ConstantPool); err != nil {
		return err
	}
	if err := binary.Write(to, Order, class.AccessFlags); err != nil {
		return err
	}
	if err := binary.Write(to, Order, class.ThisClass); err != nil {
		return err
	}
	if err := binary.Write(to, Order, class.SuperClass); err != nil {
		return err
	}
	if err := writeInterfaces(to, class.Interfaces); err != nil {
		return err
	}
	if err := writeMemberInfos(to, class.Fields); err != nil {
		return err
	}
	if err := writeMemberInfos(to, class.Methods); err != nil {
		return err
	}
	if err := writeAttributes(to, class.Attributes); err != nil {
		return err
	}
	return nil
}

func writeAttributes(to io.Writer, attributeInfo []*AttributeInfo) error {
	if err := binary.Write(to, Order, uint16(len(attributeInfo))); err != nil {
		return err
	}
	for _, attribute := range attributeInfo {
		if err := binary.Write(to, Order, attribute.AttributeNameIndex); err != nil {
			return err
		}
		if err := binary.Write(to, Order, uint32(len(attribute.Info))); err != nil {
			return err
		}
		if err := binary.Write(to, Order, attribute.Info); err != nil {
			return err
		}
	}
	return nil
}

func writeMemberInfos(to io.Writer, memberInfo []*MemberInfo) error {
	if err := binary.Write(to, Order, uint16(len(memberInfo))); err != nil {
		return err
	}
	for _, member := range memberInfo {
		if err := binary.Write(to, Order, member.AccessFlags); err != nil {
			return err
		}
		if err := binary.Write(to, Order, member.NameIndex); err != nil {
			return err
		}
		if err := binary.Write(to, Order, member.DescriptionIndex); err != nil {
			return err
		}
		if err := writeAttributes(to, member.Attributes); err != nil {
			return err
		}
	}
	return nil
}

func writeInterfaces(to io.Writer, interfaces []uint16) error {
	if err := binary.Write(to, Order, uint16(len(interfaces))); err != nil {
		return err
	}
	for _, interf := range interfaces {
		if err := binary.Write(to, Order, interf); err != nil {
			return err
		}
	}
	return nil
}

func writeConstantPool(to io.Writer, constantInfo []*ConstantInfo) error {
	if err := binary.Write(to, Order, uint16(len(constantInfo))+1); err != nil {
		return err
	}
	for _, info := range constantInfo {
		if err := binary.Write(to, Order, info.Tag); err != nil {
			return err
		}
		if err := binary.Write(to, Order, info.Info); err != nil {
			return err
		}
	}
	return nil
}
