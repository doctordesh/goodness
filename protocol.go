package goodness

import (
	"encoding/binary"
	"fmt"
	"strings"
)

const _HEADER_SIZE = 12

type Message struct {
	Data []byte
}

// ==================================================
//
// Header (12 bytes)
//
// ==================================================

func (self Message) ID() int {
	return int(binary.BigEndian.Uint16(self.Data[0:2]))
}

func (self Message) QR() QR {
	return QR(self.Data[2] & 0b10000000 >> 7)
}

func (self Message) OPCODE() OPCODE {
	v := (self.Data[2] & 0b01111000) >> 3
	return OPCODE(v)
}

func (self Message) AA() bool {
	if (self.Data[2]&0b00000100)>>2 == 1 {
		return true
	}
	return false
}

func (self Message) TC() bool {
	if (self.Data[2]&0b00000010)>>1 == 1 {
		return true
	}
	return false
}

func (self Message) RD() bool {
	if (self.Data[2] & 0b00000001) == 1 {
		return true
	}
	return false
}
func (self Message) RA() bool {
	if (self.Data[3]&0b10000000)>>7 == 1 {
		return true
	}
	return false
}

func (self Message) Z() int {
	return int((self.Data[3] & 0b01110000) >> 4)
}

func (self Message) RCODE() RCODE {
	return RCODE((self.Data[3] & 0b00001111))
}

func (self Message) QDCOUNT() int {
	return int(binary.BigEndian.Uint16(self.Data[4:6]))
}

func (self Message) ANCOUNT() int {
	return int(binary.BigEndian.Uint16(self.Data[6:8]))
}

func (self Message) NSCOUNT() int {
	return int(binary.BigEndian.Uint16(self.Data[8:10]))
}

func (self Message) ARCOUNT() int {
	return int(binary.BigEndian.Uint16(self.Data[10:12]))
}

// ==================================================
//
// Sections
//
// ==================================================

func (self Message) Questions() []Question {
	var questions []Question

	offset := 0
	for i := 0; i < self.QDCOUNT(); i++ {
		qlength := 0
		for {
			length := int(self.Data[_HEADER_SIZE+offset+qlength])
			if length == 0 {
				break
			}
			qlength += length + 1
		}

		questions = append(questions, Question{self.Data[_HEADER_SIZE+offset : _HEADER_SIZE+offset+qlength+1+2+2]})
		offset += qlength
	}

	return questions
}

// ==================================================
//
// Setters
//
// ==================================================

func (self Message) SetQR(qr QR) {
	self.Data[2] = self.Data[2] ^ (byte(qr) << 7)
}

func (self Message) SetRCODE(rcode RCODE) {
	// clear the first 4 bits
	self.Data[3] = self.Data[3] >> 4
	self.Data[3] = self.Data[3] << 4

	// set the new bits
	self.Data[3] = self.Data[3] ^ byte(rcode)
}

// ==================================================
//
// Helpers
//
// ==================================================

func (self Message) String() string {
	parts := []string{}

	parts = append(parts, fmt.Sprintf("ID: %v", self.ID()))
	parts = append(parts, fmt.Sprintf("QR: %v", self.QR()))
	parts = append(parts, fmt.Sprintf("OPCODE: %v", self.OPCODE()))
	parts = append(parts, fmt.Sprintf("AA: %v", self.AA()))
	parts = append(parts, fmt.Sprintf("TC: %v", self.TC()))
	parts = append(parts, fmt.Sprintf("RD: %v", self.RD()))
	parts = append(parts, fmt.Sprintf("RA: %v", self.RA()))
	parts = append(parts, fmt.Sprintf("Z: %v", self.Z()))
	parts = append(parts, fmt.Sprintf("RCODE: %v", self.RCODE()))
	parts = append(parts, fmt.Sprintf("QDCOUNT: %v", self.QDCOUNT()))
	parts = append(parts, fmt.Sprintf("ANCOUNT: %v", self.ANCOUNT()))
	parts = append(parts, fmt.Sprintf("NSCOUNT: %v", self.NSCOUNT()))
	parts = append(parts, fmt.Sprintf("ARCOUNT: %v", self.ARCOUNT()))

	return strings.Join(parts, "\n")
}
