package dns

import (
	"testing"

	"github.com/doctordesh/check"
)

func TestEncodeLabels(t *testing.T) {
	r := ResourceRecord{Domain: []string{"google", "com"}}
	d := r.EncodeLabels()

	check.Assert(t, d[0] == 6)
	check.Equals(t, byte('g'), d[1])
	check.Equals(t, byte('o'), d[2])
	check.Equals(t, byte('o'), d[3])
	check.Equals(t, byte('g'), d[4])
	check.Equals(t, byte('l'), d[5])
	check.Equals(t, byte('e'), d[6])

	check.Assert(t, d[7] == 3)
	check.Equals(t, byte('c'), d[8])
	check.Equals(t, byte('o'), d[9])
	check.Equals(t, byte('m'), d[10])

	check.Assert(t, len(d) == 11)
}

func TestString(t *testing.T) {
	r := ResourceRecord{
		Domain:     []string{"google", "com"},
		Type:       TYPE_A,
		Class:      CLASS_IN,
		TimeToLive: 3600,
	}

	check.Equals(t, "google.com: Type: A, Class: IN, TTL: 3600", r.String())
}

var test_response_for_google_dot_com = []byte{
	0xa9, 0xa9, 0x81, 0x80, // header
	0x00, 0x01, 0x00, 0x01, //
	0x00, 0x00, 0x00, 0x00, // end header
	// Start question
	0x06,                               // length of label
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, // 'google'
	0x03,             // length of label
	0x63, 0x6f, 0x6d, // 'com'
	0x00,       // end of labels (zero byte)
	0x00, 0x01, // qtype
	0x00, 0x01, // qclass
	// End question

	0xc0, 0x0c, // offset for label (12 in this case)
	0x00, 0x01, // type
	0x00, 0x01, // class
	0x00, 0x00, 0x00, 0xb5, // TTL
	0x00, 0x04, // Resource Data Length
	0x8e, 0xfa, 0x4a, 0x6e, // Resource Data (RDLength long)
}
