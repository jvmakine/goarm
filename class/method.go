package class

import "github.com/jvmakine/goasm/classfile"

type Method struct {
	clazz      *Class
	memberInfo *classfile.MemberInfo
}

func (f *Method) AccessFlags() *MethodAccess {
	return &MethodAccess{&f.memberInfo.AccessFlags}
}

func (f *Method) Name() *String {
	return &String{f.clazz.file, f.memberInfo.NameIndex}
}

func (f *Method) SetName(to *String) {
	validateFilesEqual(f.clazz.file, to.file)
	f.memberInfo.NameIndex = to.index
}

func (f *Method) Descriptor() *String {
	return &String{f.clazz.file, f.memberInfo.DescriptionIndex}
}

func (f *Method) SetDescriptor(to *String) {
	validateFilesEqual(f.clazz.file, to.file)
	f.memberInfo.DescriptionIndex = to.index
}

func (f *Method) Attributes() *Attributes {
	return &Attributes{f.clazz, &f.memberInfo.Attributes}
}

type Methods struct {
	clazz *Class
}

func (c *Methods) List() []*Method {
	result := make([]*Method, len(c.clazz.file.Methods))
	for i, mi := range c.clazz.file.Methods {
		result[i] = &Method{c.clazz, mi}
	}
	return result
}

func (c *Methods) New(name, descriptor *String) *Method {
	validateFilesEqual(c.clazz.file, name.file)
	validateFilesEqual(c.clazz.file, descriptor.file)
	c.clazz.file.Fields = append(c.clazz.file.Methods, &classfile.MemberInfo{
		NameIndex:        name.index,
		DescriptionIndex: descriptor.index,
	})
	return &Method{c.clazz, c.clazz.file.Methods[len(c.clazz.file.Methods)-1]}
}
