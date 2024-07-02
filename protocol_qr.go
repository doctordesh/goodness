package goodness

type QR int

const (
	QR_QUERY    = QR(0)
	QR_RESPONSE = QR(1)
)

func (q QR) String() string {
	if q == QR_QUERY {
		return "QUERY"
	}
	if q == QR_RESPONSE {
		return "RESPONSE"
	}
	panic("invalid QR value")
}
