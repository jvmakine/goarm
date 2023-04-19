package classfile

import (
	"encoding/binary"
	"io"
)

func Write(class *Classfile, to io.Writer) error {
	if err := binary.Write(to, order, magicNumber); err != nil {
		return err
	}
	if err := binary.Write(to, order, class.MinorVersion); err != nil {
		return err
	}
	if err := binary.Write(to, order, class.MajorVersion); err != nil {
		return err
	}
	if err := writeConstantPool(to, class.ConstantPool); err != nil {
		return err
	}
	if err := binary.Write(to, order, class.AccessFlags); err != nil {
		return err
	}
	if err := binary.Write(to, order, class.ThisClass); err != nil {
		return err
	}
	if err := binary.Write(to, order, class.SuperClass); err != nil {
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
	if err := binary.Write(to, order, uint16(len(attributeInfo))); err != nil {
		return err
	}
	for _, attribute := range attributeInfo {
		if err := binary.Write(to, order, attribute.AttributeNameIndex); err != nil {
			return err
		}
		if err := binary.Write(to, order, uint32(len(attribute.Info))); err != nil {
			return err
		}
		if err := binary.Write(to, order, attribute.Info); err != nil {
			return err
		}
	}
	return nil
}

func writeMemberInfos(to io.Writer, memberInfo []*MemberInfo) error {
	if err := binary.Write(to, order, uint16(len(memberInfo))); err != nil {
		return err
	}
	for _, member := range memberInfo {
		if err := binary.Write(to, order, member.AccessFlags); err != nil {
			return err
		}
		if err := binary.Write(to, order, member.NameIndex); err != nil {
			return err
		}
		if err := binary.Write(to, order, member.DescriptionIndex); err != nil {
			return err
		}
		if err := writeAttributes(to, member.Attributes); err != nil {
			return err
		}
	}
	return nil
}

func writeInterfaces(to io.Writer, interfaces []uint16) error {
	if err := binary.Write(to, order, uint16(len(interfaces))); err != nil {
		return err
	}
	for _, interf := range interfaces {
		if err := binary.Write(to, order, interf); err != nil {
			return err
		}
	}
	return nil
}

func writeConstantPool(to io.Writer, constantInfo []*ConstantInfo) error {
	if err := binary.Write(to, order, uint16(len(constantInfo))+1); err != nil {
		return err
	}
	for _, info := range constantInfo {
		if err := binary.Write(to, order, info.Tag); err != nil {
			return err
		}
		if err := binary.Write(to, order, info.Info); err != nil {
			return err
		}
	}
	return nil
}
