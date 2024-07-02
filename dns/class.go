package dns

type Class uint16

const (
	CLASS_IN = Class(1)
	CLASS_CS = Class(2)
	CLASS_CH = Class(3)
	CLASS_HS = Class(4)
)

func (c Class) String() string {
	switch c {
	case CLASS_IN:
		return "IN"
	case CLASS_CS:
		return "CS"
	case CLASS_CH:
		return "CH"
	case CLASS_HS:
		return "HS"
	}
	panic("invalid value for 'Class'")
}
