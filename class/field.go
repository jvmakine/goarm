package class

import "github.com/jvmakine/goasm/classfile"

type Field struct {
	clazz      *Class
	memberInfo *classfile.MemberInfo
}

func (f *Field) AccessFlags() *FieldAccess {
	return &FieldAccess{&f.memberInfo.AccessFlags}
}

func (f *Field) Name() *String {
	return &String{f.clazz.file, f.memberInfo.NameIndex}
}

func (f *Field) SetName(to *String) {
	validateFilesEqual(f.clazz.file, to.file)
	f.memberInfo.NameIndex = to.index
}

func (f *Field) Descriptor() *String {
	return &String{f.clazz.file, f.memberInfo.DescriptionIndex}
}

func (f *Field) SetDescriptor(to *String) {
	validateFilesEqual(f.clazz.file, to.file)
	f.memberInfo.DescriptionIndex = to.index
}

func (c *Field) Attributes() *Attributes {
	return &Attributes{c.clazz, &c.memberInfo.Attributes}
}

type Fields struct {
	clazz *Class
}

func (c *Fields) List() []*Field {
	result := make([]*Field, len(c.clazz.file.Fields))
	for i, mi := range c.clazz.file.Fields {
		result[i] = &Field{c.clazz, mi}
	}
	return result
}

func (c *Fields) New(name, descriptor *String) *Field {
	validateFilesEqual(c.clazz.file, name.file)
	validateFilesEqual(c.clazz.file, descriptor.file)
	c.clazz.file.Fields = append(c.clazz.file.Fields, &classfile.MemberInfo{
		NameIndex:        name.index,
		DescriptionIndex: descriptor.index,
	})
	return &Field{c.clazz, c.clazz.file.Fields[len(c.clazz.file.Fields)-1]}
}
