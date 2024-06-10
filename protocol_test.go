package goodness

import (
	"testing"

	"github.com/doctordesh/check"
)

// func TestProtocolParsing(t *testing.T) {
// 	b := []byte{116, 237, 1, 32, 0, 1, 0, 0, 0, 0, 0, 1, 6, 103, 111, 111, 103, 108, 101, 3, 99, 111, 109, 0, 0, 1, 0, 1, 0, 0, 41, 16, 0, 0, 0, 0, 0, 0}

// 	m := Message{Data: b}

// 	check.Equals(t, 29933, m.ID())
// 	check.Equals(t, QR_QUERY, m.QR())
// 	check.Equals(t, OPCODE_QUERY, m.OPCODE())
// 	check.Equals(t, false, m.AA())
// 	check.Equals(t, false, m.TC())
// 	check.Equals(t, false, m.RD())
// }

func TestProtocolParsingHeader(t *testing.T) {
	b := []byte{0xFF, 0x00}
	b = append(b, 0b10010111)
	b = append(b, 0b10000101)
	b = append(b, 0x00, 0xF0)
	b = append(b, 0x00, 0xF1)
	b = append(b, 0x00, 0xF2)
	b = append(b, 0x00, 0xF3)

	m := Message{Data: b}

	check.Equals(t, 65280, m.ID())
	check.Equals(t, QR_RESPONSE, m.QR())
	check.Equals(t, OPCODE_STATUS, m.OPCODE())
	check.Equals(t, true, m.AA())
	check.Equals(t, true, m.TC())
	check.Equals(t, true, m.RD())
	check.Equals(t, true, m.RA())
	check.Equals(t, 0, m.Z())
	check.Equals(t, RCODE_REFUSED, m.RCODE())

	check.Equals(t, 240, m.QDCOUNT())
	check.Equals(t, 241, m.ANCOUNT())
	check.Equals(t, 242, m.NSCOUNT())
	check.Equals(t, 243, m.ARCOUNT())
}

func TestSomething(t *testing.T) {
	t.Logf("\n%s", bitsAND(0b10010111, 0b01110000))
	t.Logf("\n%s", bitsOR(0b10010111, 0b01110000))
	t.Logf("\n%s", bitsXOR(0b10010111, 0b01110000))
	t.Logf("\n%s", bitsRSHIFT(0b01111000, 3))
	t.Logf("\n%s", bitsLSHIFT(0b00001111, 4))
}

func TestMessageSetQR(t *testing.T) {
	m := Message{[]byte{0, 0, 0, 0}}
	m.SetQR(QR_RESPONSE)
	check.Equals(t, byte(0b10000000), m.Data[2])

	m.SetRCODE(RCODE_REFUSED)
	check.Equals(t, byte(0x05), m.Data[3])

	// Make sure the first 4 bits are preserved
	m.Data[3] = 0xFF
	m.SetRCODE(RCODE_REFUSED)
	check.Equals(t, byte(0xF5), m.Data[3])
}
