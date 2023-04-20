package class

import "github.com/jvmakine/goasm/classfile"

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

func (c *Field) Attributes() *Attributes {
	return &Attributes{c.file, c.memberInfo.Attributes}
}

type Fields struct {
	file *classfile.Classfile
}

func (c *Fields) List() []*Field {
	result := make([]*Field, len(c.file.Fields))
	for i, mi := range c.file.Fields {
		result[i] = &Field{c.file, mi}
	}
	return result
}

func (c *Fields) New(name, descriptor *String) *Field {
	validateFilesEqual(c.file, name.file)
	validateFilesEqual(c.file, descriptor.file)
	c.file.Fields = append(c.file.Fields, &classfile.MemberInfo{
		NameIndex:        name.index,
		DescriptionIndex: descriptor.index,
	})
	return &Field{c.file, c.file.Fields[len(c.file.Fields)-1]}
}
