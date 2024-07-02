package goodness

import (
	"encoding/binary"
	"fmt"
	"strings"
)

type Type uint16

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
	TYPE_AAAA  = Type(28)
	TYPE_HTTPS = Type(65)
)

func (t Type) String() string {
	switch t {
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
	case TYPE_AAAA:
		return "AAAA"
	case TYPE_HTTPS:
		return "HTTPS"
	}

	panic(fmt.Sprintf("invalid value for 'Type': %d", int(t)))
}

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

type Question struct {
	Data []byte
}

func (q Question) NameLabels() []string {
	offset := 0
	parts := []string{}
	for {
		length := int(q.Data[offset])
		if length == 0 {
			break
		}

		parts = append(parts, string(q.Data[offset+1:offset+1+length]))
		offset = offset + length + 1
	}
	return parts
}

func (q Question) nameOffset() int {
	// Assumption, the first zero byte is the end of the name labels
	for i := 0; i < len(q.Data); i++ {
		if q.Data[i] == 0 {
			return i + 1
		}
	}
	panic("Question type needs a zero byte to split the name labels and type + class")
}

func (q Question) Type() Type {
	offset := q.nameOffset()
	return Type(binary.BigEndian.Uint16(q.Data[offset : offset+2]))
}

func (q Question) Class() Class {
	offset := q.nameOffset() + 2
	return Class(binary.BigEndian.Uint16(q.Data[offset : offset+2]))
}

func (q Question) String() string {
	return fmt.Sprintf("%T<%s, Type: %s, Class: %s>", q, strings.Join(q.NameLabels(), "."), q.Type(), q.Class())
}
