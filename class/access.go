package class

type AccessFlag uint16

const (
	ACC_PUBLIC     AccessFlag = 0x0001
	ACC_PRIVATE               = 0x0002
	ACC_PROTECTED             = 0x0004
	ACC_STATIC                = 0x0008
	ACC_FINAL                 = 0x0010
	ACC_VOLATILE              = 0x0040
	ACC_SUPER                 = 0x0020
	ACC_TRANSIENT             = 0x0080
	ACC_INTERFACE             = 0x0200
	ACC_ABSTRACT              = 0x0400
	ACC_SYNTHETIC             = 0x1000
	ACC_ANNOTATION            = 0x2000
	ACC_ENUM                  = 0x4000
	// method specific
	ACC_SYNCHRONIZED = 0x00200
	ACC_BRIDGE       = 0x0040
	ACC_VARARGS      = 0x0080
	ACC_NATIVE       = 0x0100
	ACC_STRICT       = 0x0800
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

type FieldAccess struct {
	field *uint16
}

func (a *FieldAccess) IsPublic() bool {
	return (*a.field)&uint16(ACC_PUBLIC) != 0
}

func (a *FieldAccess) SetPublic(is bool) {
	*a.field = *a.field | uint16(ACC_PUBLIC)
}

func (a *FieldAccess) IsFinal() bool {
	return (*a.field)&uint16(ACC_FINAL) != 0
}

func (a *FieldAccess) SetFinal(is bool) {
	*a.field = *a.field | uint16(ACC_FINAL)
}

func (a *FieldAccess) IsSynthetic() bool {
	return (*a.field)&uint16(ACC_SYNTHETIC) != 0
}

func (a *FieldAccess) SetSynthetic(is bool) {
	*a.field = *a.field | uint16(ACC_SYNTHETIC)
}

func (a *FieldAccess) IsEnum() bool {
	return (*a.field)&uint16(ACC_ENUM) != 0
}

func (a *FieldAccess) SetEnum(is bool) {
	*a.field = *a.field | uint16(ACC_ENUM)
}

func (a *FieldAccess) IsProtected() bool {
	return (*a.field)&uint16(ACC_PROTECTED) != 0
}

func (a *FieldAccess) SetProtected(is bool) {
	*a.field = *a.field | uint16(ACC_PROTECTED)
}

func (a *FieldAccess) IsPrivate() bool {
	return (*a.field)&uint16(ACC_PRIVATE) != 0
}

func (a *FieldAccess) SetPrivate(is bool) {
	*a.field = *a.field | uint16(ACC_PRIVATE)
}

func (a *FieldAccess) IsStatic() bool {
	return (*a.field)&uint16(ACC_STATIC) != 0
}

func (a *FieldAccess) SetStatic(is bool) {
	*a.field = *a.field | uint16(ACC_STATIC)
}

func (a *FieldAccess) IsVolatile() bool {
	return (*a.field)&uint16(ACC_VOLATILE) != 0
}

func (a *FieldAccess) SetVolatile(is bool) {
	*a.field = *a.field | uint16(ACC_VOLATILE)
}

func (a *FieldAccess) IsTransient() bool {
	return (*a.field)&uint16(ACC_TRANSIENT) != 0
}

func (a *FieldAccess) SetTransient(is bool) {
	*a.field = *a.field | uint16(ACC_TRANSIENT)
}

type MethodAccess struct {
	field *uint16
}

func (a *MethodAccess) IsPublic() bool {
	return (*a.field)&uint16(ACC_PUBLIC) != 0
}

func (a *MethodAccess) SetPublic(is bool) {
	*a.field = *a.field | uint16(ACC_PUBLIC)
}

func (a *MethodAccess) IsFinal() bool {
	return (*a.field)&uint16(ACC_FINAL) != 0
}

func (a *MethodAccess) SetFinal(is bool) {
	*a.field = *a.field | uint16(ACC_FINAL)
}

func (a *MethodAccess) IsSynthetic() bool {
	return (*a.field)&uint16(ACC_SYNTHETIC) != 0
}

func (a *MethodAccess) SetSynthetic(is bool) {
	*a.field = *a.field | uint16(ACC_SYNTHETIC)
}

func (a *MethodAccess) IsProtected() bool {
	return (*a.field)&uint16(ACC_PROTECTED) != 0
}

func (a *MethodAccess) SetProtected(is bool) {
	*a.field = *a.field | uint16(ACC_PROTECTED)
}

func (a *MethodAccess) IsPrivate() bool {
	return (*a.field)&uint16(ACC_PRIVATE) != 0
}

func (a *MethodAccess) SetPrivate(is bool) {
	*a.field = *a.field | uint16(ACC_PRIVATE)
}

func (a *MethodAccess) IsStatic() bool {
	return (*a.field)&uint16(ACC_STATIC) != 0
}

func (a *MethodAccess) SetStatic(is bool) {
	*a.field = *a.field | uint16(ACC_STATIC)
}

func (a *MethodAccess) IsSynchronized() bool {
	return (*a.field)&uint16(ACC_SYNCHRONIZED) != 0
}

func (a *MethodAccess) SetSynchronized(is bool) {
	*a.field = *a.field | uint16(ACC_SYNCHRONIZED)
}

func (a *MethodAccess) IsBridge() bool {
	return (*a.field)&uint16(ACC_BRIDGE) != 0
}

func (a *MethodAccess) SetBridge(is bool) {
	*a.field = *a.field | uint16(ACC_BRIDGE)
}

func (a *MethodAccess) IsVarArgs() bool {
	return (*a.field)&uint16(ACC_VARARGS) != 0
}

func (a *MethodAccess) SetVarArgs(is bool) {
	*a.field = *a.field | uint16(ACC_VARARGS)
}

func (a *MethodAccess) IsNative() bool {
	return (*a.field)&uint16(ACC_NATIVE) != 0
}

func (a *MethodAccess) SetNative(is bool) {
	*a.field = *a.field | uint16(ACC_NATIVE)
}

func (a *MethodAccess) IsAbstract() bool {
	return (*a.field)&uint16(ACC_ABSTRACT) != 0
}

func (a *MethodAccess) SetAbstract(is bool) {
	*a.field = *a.field | uint16(ACC_ABSTRACT)
}

func (a *MethodAccess) IsStrict() bool {
	return (*a.field)&uint16(ACC_STRICT) != 0
}

func (a *MethodAccess) SetStrict(is bool) {
	*a.field = *a.field | uint16(ACC_STRICT)
}
