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

func (f *Method) Attributes() *Attributes {
	return &Attributes{f.file, f.memberInfo.Attributes}
}

type Methods struct {
	file *classfile.Classfile
}

func (c *Methods) List() []*Method {
	result := make([]*Method, len(c.file.Methods))
	for i, mi := range c.file.Methods {
		result[i] = &Method{c.file, mi}
	}
	return result
}

func (c *Methods) New(name, descriptor *String) *Method {
	validateFilesEqual(c.file, name.file)
	validateFilesEqual(c.file, descriptor.file)
	c.file.Fields = append(c.file.Methods, &classfile.MemberInfo{
		NameIndex:        name.index,
		DescriptionIndex: descriptor.index,
	})
	return &Method{c.file, c.file.Methods[len(c.file.Methods)-1]}
}
