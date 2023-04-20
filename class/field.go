package class

import "github.com/jvmakine/goarm/classfile"

type Field struct {
	file       *classfile.Classfile
	memberInfo *classfile.MemberInfo
}

func (f *Field) AccessFlags() *FieldAccess {
	return &FieldAccess{&f.memberInfo.AccessFlags}
}

func (f *Field) Name() *String {
	return &String{f.file, f.memberInfo.NameIndex}
}

func (f *Field) SetName(to *String) {
	validateFilesEqual(f.file, to.file)
	f.memberInfo.NameIndex = to.index
}

func (f *Field) Descriptor() *String {
	return &String{f.file, f.memberInfo.DescriptionIndex}
}

func (f *Field) SetDescriptor(to *String) {
	validateFilesEqual(f.file, to.file)
	f.memberInfo.DescriptionIndex = to.index
}
