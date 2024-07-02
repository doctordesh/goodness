package dns

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
	}

	panic("invalid value for 'Type'")
}
