package class

import "github.com/jvmakine/goasm/classfile"

type Attribute struct {
	file      *classfile.Classfile
	attribute *classfile.AttributeInfo
}

func (a *Attribute) Name() *String {
	return &String{a.file, a.attribute.AttributeNameIndex}
}

func (a *Attribute) SetName(to *String) {
	validateFilesEqual(a.file, to.file)
	a.attribute.AttributeNameIndex = to.index
}

func (a *Attribute) Info() []byte {
	return a.attribute.Info
}

func (a *Attribute) SetInfo(to []byte) {
	a.attribute.Info = to
}

type Attributes struct {
	clazz *Class
	slice *[]*classfile.AttributeInfo
}

func (c *Attributes) List() []*Attribute {
	result := make([]*Attribute, len(*c.slice))
	for i, a := range *c.slice {
		result[i] = &Attribute{c.clazz.file, a}
	}
	return result
}

func (c *Attributes) New(name *String, info []byte) *Attribute {
	validateFilesEqual(c.clazz.file, name.file)
	newSlice := append(*c.slice, &classfile.AttributeInfo{
		AttributeNameIndex: name.index,
		Info:               info,
	})
	*c.slice = newSlice
	return &Attribute{c.clazz.file, (*c.slice)[len(*c.slice)-1]}
}

func (c *Attributes) NewOrReplace(name string, info []byte) *Attribute {
	for i, a := range c.List() {
		if a.Name().Text() == name {
			a.SetInfo(info)
			return &Attribute{c.clazz.file, (*c.slice)[i]}
		}
	}
	str := c.clazz.Constants().NewString(name)
	return c.New(str, info)
}
