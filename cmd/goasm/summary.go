package main

import (
	"github.com/jvmakine/goasm/class"
	"github.com/jvmakine/goasm/classfile"
)

type ClassSummary struct {
	Name       string              `yaml:"name"`
	Super      string              `yaml:"super"`
	Interfaces []string            `yaml:"interfaces"`
	Attributes []*AttributeSummary `yaml:"attributes"`
	Constants  []*ConstantSummary  `yaml:"constants"`
	Methods    []*MethodSummary    `yaml:"methods"`
	Fields     []*FieldSummary     `yaml:"fields"`
}

type AttributeSummary struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type MethodSummary struct {
	Name       string              `yaml:"name"`
	Descriptor string              `yaml:"descriptor"`
	Attributes []*AttributeSummary `yaml:"attributes"`
}

type FieldSummary struct {
	Name       string              `yaml:"name"`
	Descriptor string              `yaml:"descriptor"`
	Attributes []*AttributeSummary `yaml:"attributes"`
}

type ConstantSummary struct {
	Tag   int    `yaml:"tag"`
	Value string `yaml:"value"`
}

func SummaryFrom(clazz *class.Class, file *classfile.Classfile) *ClassSummary {
	return &ClassSummary{
		Name:       clazz.ThisClass().Name().Text(),
		Super:      clazz.SuperClass().Name().Text(),
		Interfaces: interfacesSummary(clazz.Interfaces().List()),
		Attributes: attributesSummary(clazz.Attributes().List()),
		Constants:  constantSummary(file.ConstantPool),
		Methods:    methodSummary(clazz.Methods().List()),
		Fields:     fieldSummary(clazz.Fields().List()),
	}
}

func constantSummary(constants []*classfile.ConstantInfo) []*ConstantSummary {
	result := make([]*ConstantSummary, len(constants))
	for i, c := range constants {
		result[i] = &ConstantSummary{
			Tag:   int(c.Tag),
			Value: string(c.Info),
		}
	}
	return result
}

func methodSummary(methods []*class.Method) []*MethodSummary {
	result := make([]*MethodSummary, len(methods))
	for i, m := range methods {
		result[i] = &MethodSummary{
			Name:       m.Name().Text(),
			Descriptor: m.Descriptor().Text(),
			Attributes: attributesSummary(m.Attributes().List()),
		}
	}
	return result
}

func fieldSummary(fields []*class.Field) []*FieldSummary {
	result := make([]*FieldSummary, len(fields))
	for i, m := range fields {
		result[i] = &FieldSummary{
			Name:       m.Name().Text(),
			Descriptor: m.Descriptor().Text(),
			Attributes: attributesSummary(m.Attributes().List()),
		}
	}
	return result
}

func interfacesSummary(interfaces []*class.ClassInfo) []string {
	result := make([]string, len(interfaces))
	for i, a := range interfaces {
		result[i] = a.Name().Text()
	}
	return result
}

func attributesSummary(attributes []*class.Attribute) []*AttributeSummary {
	result := make([]*AttributeSummary, len(attributes))
	for i, a := range attributes {
		result[i] = &AttributeSummary{
			Name:  a.Name().Text(),
			Value: string(a.Info()),
		}
	}
	return result
}
