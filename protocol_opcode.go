package goodness

type OPCODE int

const (
	OPCODE_QUERY = OPCODE(iota)
	OPCODE_IQUERY
	OPCODE_STATUS
)

func (o OPCODE) String() string {
	if o == OPCODE_QUERY {
		return "QUERY"
	}
	if o == OPCODE_IQUERY {
		return "IQUERY"
	}
	if o == OPCODE_STATUS {
		return "STATUS"
	}
	panic("invalid OPCODE value")
}
