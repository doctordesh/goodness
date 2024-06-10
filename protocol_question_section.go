package goodness

import "encoding/binary"

type Type int

const (
	TYPE_A     = Type(1)
	TYPE_NS    = Type(2)
	TYPE_MD    = Type(3)
	TYPE_MF    = Type(4)
	TYPE_CNAME = Type(5)
	TYPE_SOA   = Type(6)
	TYPE_MB    = Type(7)
	TYPE_MG    = Type(8)
	TYPE_MR    = Type(9)
	TYPE_NULL  = Type(10)
	TYPE_WKS   = Type(11)
	TYPE_PTR   = Type(12)
	TYPE_HINFO = Type(13)
	TYPE_MINFO = Type(14)
	TYPE_MX    = Type(15)
	TYPE_TXT   = Type(16)
)

func (self Type) String() string {
	switch self {
	case TYPE_A:
		return "A"
	case TYPE_NS:
		return "NS"
	case TYPE_MD:
		return "MD"
	case TYPE_MF:
		return "MF"
	case TYPE_CNAME:
		return "CNAME"
	case TYPE_SOA:
		return "SOA"
	case TYPE_MB:
		return "MB"
	case TYPE_MG:
		return "MG"
	case TYPE_MR:
		return "MR"
	case TYPE_NULL:
		return "NULL"
	case TYPE_WKS:
		return "WKS"
	case TYPE_PTR:
		return "PTR"
	case TYPE_HINFO:
		return "HINFO"
	case TYPE_MINFO:
		return "MINFO"
	case TYPE_MX:
		return "MX"
	case TYPE_TXT:
		return "TXT"
	}

	panic("invalid value for 'Type'")
}

type Class int

const (
	CLASS_IN = Class(1)
	CLASS_CS = Class(2)
	CLASS_CH = Class(3)
	CLASS_HS = Class(4)
)

func (self Class) String() string {
	switch self {
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

type Question struct {
	Data []byte
}

func (self Question) NameLabels() []string {
	offset := 0
	parts := []string{}
	for {
		length := int(self.Data[offset])
		if length == 0 {
			break
		}

		parts = append(parts, string(self.Data[offset+1:offset+1+length]))
		offset = offset + length + 1
	}
	return parts
}

func (self Question) nameOffset() int {
	// Assumption, the first zero byte is the end of the name labels
	for i := 0; i < len(self.Data); i++ {
		if self.Data[i] == 0 {
			return i + 1
		}
	}
	panic("Question type needs a zero byte to split the name labels and type + class")
}

func (self Question) Type() Type {
	offset := self.nameOffset()
	return Type(binary.BigEndian.Uint16(self.Data[offset : offset+2]))
}

func (self Question) Class() Class {
	offset := self.nameOffset() + 2
	return Class(binary.BigEndian.Uint16(self.Data[offset : offset+2]))
}
