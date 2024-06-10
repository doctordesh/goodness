package goodness

type RCODE int

const (
	RCODE_NOERROR = RCODE(iota)
	RCODE_FORMAT_ERROR
	RCODE_SERVER_FAILURE
	RCODE_NAME_ERROR
	RCODE_NOT_IMPLEMENTED
	RCODE_REFUSED
)

func (self RCODE) String() string {
	if self == RCODE_NOERROR {
		return "NOERROR"
	}
	if self == RCODE_FORMAT_ERROR {
		return "FORMAT_ERROR"
	}
	if self == RCODE_SERVER_FAILURE {
		return "SERVER_FAILURE"
	}
	if self == RCODE_NAME_ERROR {
		return "NAME_ERROR"
	}
	if self == RCODE_NOT_IMPLEMENTED {
		return "NOT_IMPLEMENTED"
	}
	if self == RCODE_REFUSED {
		return "REFUSED"
	}
	panic("invalid RCODE value")
}
