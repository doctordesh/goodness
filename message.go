package goodness

import (
	"encoding/binary"
	"fmt"
	"strings"
)

const _HEADER_SIZE = 12
const _OFFSET_MASK = 0b11000000

type ResourceRecord struct {
	Name  []string
	Type  Type
	Class Class
	TTL   int
	Data  []byte
}

type Message struct {
	Data              []byte
	Questions         []Question
	AnswerRecords     []ResourceRecord
	AuthorityRecords  []ResourceRecord
	AdditionalRecords []ResourceRecord
}

// Message parses the message header and fields to be able to later find all the
// different resources regardless of access order
func (m *Message) Parse() error {
	offset := _HEADER_SIZE

	// Parse Questions
	for i := 0; i < m.QDCOUNT(); i++ {
		qlength := 0
		for {
			length := int(m.Data[offset+qlength])
			if length == 0 {
				break
			}
			qlength += length + 1
		}

		m.Questions = append(m.Questions, Question{m.Data[offset : offset+qlength+1+2+2]})
		offset += qlength + 1 + 2 + 2
	}

	// Parse Resource Records

	// if _OFFSET_MASK&m.Data[offset] == _OFFSET_MASK {
	// 	pointer := binary.BigEndian.Uint16([]byte{m.Data[offset] ^ _OFFSET_MASK, m.Data[offset+1]})
	// 	fmt.Println("Pointer: ", pointer)
	// 	labels := readLabels(m.Data[pointer:])
	// 	fmt.Println("Labels: ", labels)
	// }
	// fmt.Println("--------------------------------------------------")
	// fmt.Printf("0x%02X 0x%02X 0x%02X 0x%02X \n", 0x8e, 0xfa, 0x4a, 0x6e)
	// fmt.Printf("%d %d %d %d\n", 0x8e, 0xfa, 0x4a, 0x6e)
	// for _, b := range m.Data[offset:] {
	// 	fmt.Printf("0x%02x (%d), ", b, b)
	// }
	// fmt.Println("\n--------------------------------------------------")
	// fmt.Println("Answer Count:     ", m.ANCOUNT())
	// fmt.Println("Name server Count:", m.NSCOUNT())
	// fmt.Println("Authority Count:  ", m.ARCOUNT())

	// fmt.Println(offset)

	return nil
}

func readLabels(data []byte) []string {
	labels := []string{}
	offset := 0
	length := 0
	for {
		length = int(data[offset])
		if length == 0 {
			break
		}
		labels = append(labels, string(data[offset+1:offset+1+length]))
		offset += 1 + length
	}

	return labels
}

// ==================================================
//
// Header (12 bytes)
//
// ==================================================

func (m Message) ID() int {
	return int(binary.BigEndian.Uint16(m.Data[0:2]))
}

func (m Message) QR() QR {
	return QR(m.Data[2] & 0b10000000 >> 7)
}

func (m Message) OPCODE() OPCODE {
	v := (m.Data[2] & 0b01111000) >> 3
	return OPCODE(v)
}

func (m Message) AA() bool {
	if (m.Data[2]&0b00000100)>>2 == 1 {
		return true
	}
	return false
}

func (m Message) TC() bool {
	if (m.Data[2]&0b00000010)>>1 == 1 {
		return true
	}
	return false
}

func (m Message) RD() bool {
	if (m.Data[2] & 0b00000001) == 1 {
		return true
	}
	return false
}
func (m Message) RA() bool {
	if (m.Data[3]&0b10000000)>>7 == 1 {
		return true
	}
	return false
}

func (m Message) Z() int {
	return int((m.Data[3] & 0b01110000) >> 4)
}

func (m Message) RCODE() RCODE {
	return RCODE((m.Data[3] & 0b00001111))
}

func (m Message) QDCOUNT() int {
	return int(binary.BigEndian.Uint16(m.Data[4:6]))
}

func (m Message) ANCOUNT() int {
	return int(binary.BigEndian.Uint16(m.Data[6:8]))
}

func (m Message) NSCOUNT() int {
	return int(binary.BigEndian.Uint16(m.Data[8:10]))
}

func (m Message) ARCOUNT() int {
	return int(binary.BigEndian.Uint16(m.Data[10:12]))
}

// ==================================================
//
// Sections
//
// ==================================================

// func (m Message) Questions() []Question {
// 	var questions []Question

// 	offset := 0
// 	for i := 0; i < m.QDCOUNT(); i++ {
// 		qlength := 0
// 		for {
// 			length := int(m.Data[_HEADER_SIZE+offset+qlength])
// 			if length == 0 {
// 				break
// 			}
// 			qlength += length + 1
// 		}

// 		questions = append(questions, Question{m.Data[_HEADER_SIZE+offset : _HEADER_SIZE+offset+qlength+1+2+2]})
// 		offset += qlength
// 	}

// 	return questions
// }

// ==================================================
//
// Setters
//
// ==================================================

func (m Message) SetQR(qr QR) {
	m.Data[2] = m.Data[2] ^ (byte(qr) << 7)
}

func (m Message) SetRCODE(rcode RCODE) {
	// clear the first 4 bits
	m.Data[3] = m.Data[3] >> 4
	m.Data[3] = m.Data[3] << 4

	// set the new bits
	m.Data[3] = m.Data[3] ^ byte(rcode)
}

// ==================================================
//
// Helpers
//
// ==================================================

func (m Message) String() string {
	parts := []string{}

	parts = append(parts, fmt.Sprintf("ID:                        %v", m.ID()))
	parts = append(parts, fmt.Sprintf("QueryResponse (QR):        %v", m.QR()))
	parts = append(parts, fmt.Sprintf("OPCODE:                    %v", m.OPCODE()))
	parts = append(parts, fmt.Sprintf("Authoritative Answer (AA): %v", m.AA()))
	parts = append(parts, fmt.Sprintf("TrunCation (TC):           %v", m.TC()))
	parts = append(parts, fmt.Sprintf("Recursion Desired (RD):    %v", m.RD()))
	parts = append(parts, fmt.Sprintf("Recursion Available (RA):  %v", m.RA()))
	parts = append(parts, fmt.Sprintf("Z:                         %v", m.Z()))
	parts = append(parts, fmt.Sprintf("Response code (RCODE):     %v", m.RCODE()))
	parts = append(parts, fmt.Sprintf("QDCOUNT:                   %v", m.QDCOUNT()))
	parts = append(parts, fmt.Sprintf("ANCOUNT:                   %v", m.ANCOUNT()))
	parts = append(parts, fmt.Sprintf("NSCOUNT:                   %v", m.NSCOUNT()))
	parts = append(parts, fmt.Sprintf("ARCOUNT:                   %v", m.ARCOUNT()))

	return strings.Join(parts, "\n")
}
