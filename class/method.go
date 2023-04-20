package class

import "github.com/jvmakine/goasm/classfile"

type Method struct {
	file       *classfile.Classfile
	memberInfo *classfile.MemberInfo
}

func (f *Method) AccessFlags() *MethodAccess {
	return &MethodAccess{&f.memberInfo.AccessFlags}
}

func (f *Method) Name() *String {
	return &String{f.file, f.memberInfo.NameIndex}
}

func (f *Method) SetName(to *String) {
	validateFilesEqual(f.file, to.file)
	f.memberInfo.NameIndex = to.index
}

func (f *Method) Descriptor() *String {
	return &String{f.file, f.memberInfo.DescriptionIndex}
}

func (f *Method) SetDescriptor(to *String) {
	validateFilesEqual(f.file, to.file)
	f.memberInfo.DescriptionIndex = to.index
}

func (f *Method) Attributes() []*Attribute {
	result := make([]*Attribute, len(f.memberInfo.Attributes))
	for i, a := range f.memberInfo.Attributes {
		result[i] = &Attribute{f.file, a}
	}
	return result
}
