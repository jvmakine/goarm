package class

type AccessFlag uint16

const (
	ACC_PUBLIC     AccessFlag = 0x0001
	ACC_FINAL                 = 0x0010
	ACC_SUPER                 = 0x0020
	ACC_INTERFACE             = 0x0200
	ACC_ABSTRACT              = 0x0400
	ACC_SYNTHETIC             = 0x1000
	ACC_ANNOTATION            = 0x2000
	ACC_ENUM                  = 0x4000
)

type ClassAccess struct {
	field *uint16
}

func (a *ClassAccess) IsPublic() bool {
	return (*a.field)&uint16(ACC_PUBLIC) != 0
}

func (a *ClassAccess) SetPublic(is bool) {
	*a.field = *a.field | uint16(ACC_PUBLIC)
}

func (a *ClassAccess) IsFinal() bool {
	return (*a.field)&uint16(ACC_FINAL) != 0
}

func (a *ClassAccess) SetFinal(is bool) {
	*a.field = *a.field | uint16(ACC_FINAL)
}

func (a *ClassAccess) IsSuper() bool {
	return (*a.field)&uint16(ACC_SUPER) != 0
}

func (a *ClassAccess) SetSuper(is bool) {
	*a.field = *a.field | uint16(ACC_SUPER)
}

func (a *ClassAccess) IsInterface() bool {
	return (*a.field)&uint16(ACC_INTERFACE) != 0
}

func (a *ClassAccess) SetInterface(is bool) {
	*a.field = *a.field | uint16(ACC_INTERFACE)
}

func (a *ClassAccess) IsAbstract() bool {
	return (*a.field)&uint16(ACC_ABSTRACT) != 0
}

func (a *ClassAccess) SetAbstract(is bool) {
	*a.field = *a.field | uint16(ACC_ABSTRACT)
}

func (a *ClassAccess) IsSynthetic() bool {
	return (*a.field)&uint16(ACC_SYNTHETIC) != 0
}

func (a *ClassAccess) SetSynthetic(is bool) {
	*a.field = *a.field | uint16(ACC_SYNTHETIC)
}

func (a *ClassAccess) IsAnnotation() bool {
	return (*a.field)&uint16(ACC_ANNOTATION) != 0
}

func (a *ClassAccess) SetAnnotation(is bool) {
	*a.field = *a.field | uint16(ACC_ANNOTATION)
}

func (a *ClassAccess) IsEnum() bool {
	return (*a.field)&uint16(ACC_ENUM) != 0
}

func (a *ClassAccess) SetEnum(is bool) {
	*a.field = *a.field | uint16(ACC_ENUM)
}
