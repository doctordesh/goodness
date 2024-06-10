package goodness

type QR int

const (
	QR_QUERY    = QR(0)
	QR_RESPONSE = QR(1)
)

func (self QR) String() string {
	if self == QR_QUERY {
		return "QUERY"
	}
	if self == QR_RESPONSE {
		return "RESPONSE"
	}
	panic("invalid QR value")
}
