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

var AllAccessFlags []AccessFlag = []AccessFlag{
	ACC_PUBLIC,
	ACC_FINAL,
	ACC_SUPER,
	ACC_INTERFACE,
	ACC_ABSTRACT,
	ACC_SYNTHETIC,
	ACC_ANNOTATION,
	ACC_ENUM,
}
