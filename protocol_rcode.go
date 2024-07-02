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

func (r RCODE) String() string {
	if r == RCODE_NOERROR {
		return "NOERROR"
	}
	if r == RCODE_FORMAT_ERROR {
		return "FORMAT_ERROR"
	}
	if r == RCODE_SERVER_FAILURE {
		return "SERVER_FAILURE"
	}
	if r == RCODE_NAME_ERROR {
		return "NAME_ERROR"
	}
	if r == RCODE_NOT_IMPLEMENTED {
		return "NOT_IMPLEMENTED"
	}
	if r == RCODE_REFUSED {
		return "REFUSED"
	}
	panic("invalid RCODE value")
}
