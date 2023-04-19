package class

import (
	"fmt"

	"github.com/jvmakine/goarm/classfile"
)

type String struct {
	file  *classfile.Classfile
	index uint16
}

func (s *String) Text() string {
	entry := s.file.ConstantPool[s.index-1]
	if entry.Tag != classfile.CONSTANT_Utf8 {
		panic(fmt.Errorf("incorrect string constant type: %d", entry.Tag))
	}
	length := classfile.Order.Uint16(entry.Info)
	if len(entry.Info) != int(length)+2 {
		panic(fmt.Errorf("incorrect string length: %d, expected %d", len(entry.Info)-2, length))
	}
	return string(entry.Info[2:])
}

func (s *String) SetText(to string) {
	entry := s.file.ConstantPool[s.index-1]
	if entry.Tag != classfile.CONSTANT_Utf8 {
		panic(fmt.Errorf("incorrect string constant type: %d", entry.Tag))
	}
	bytes := []byte(to)

	lengthBytes := make([]byte, 2)
	classfile.Order.PutUint16(lengthBytes, uint16(len(bytes)))

	entry.Info = append(lengthBytes, bytes...)
}
