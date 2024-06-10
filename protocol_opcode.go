package goodness

type OPCODE int

const (
	OPCODE_QUERY = OPCODE(iota)
	OPCODE_IQUERY
	OPCODE_STATUS
)

func (self OPCODE) String() string {
	if self == OPCODE_QUERY {
		return "QUERY"
	}
	if self == OPCODE_IQUERY {
		return "IQUERY"
	}
	if self == OPCODE_STATUS {
		return "STATUS"
	}
	panic("invalid OPCODE value")
}
