package goodness

import (
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
)

type Goodness struct{}

func New() Goodness {
	return Goodness{}
}

func (self Goodness) Build() {
	b := []byte{}

	// ID
	id := uint16(10_000)
	b = append(b, 0, 0)
	binary.BigEndian.PutUint16(b[0:2], id)

	// QR, Opcode, AA, TC, RD
	l := byte(0) // left
	// RA, Z, RCode
	r := byte(0) // right

	b = append(b, l)
	b = append(b, r)

	fmt.Println(b)
}

// Server ...
func (self Goodness) Serve() {
	var err error
	addr, err := net.ResolveUDPAddr("udp", ":53")
	if err != nil {
		panic(err)
	}

	sock, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Listening on %v\n", addr)

	sock.SetReadBuffer(100_000)

	i := 0
	for {
		i++
		buf := make([]byte, 512)
		n, clientAddr, err := sock.ReadFromUDP(buf)
		if err != nil {
			panic(err)
		}

		m := Message{buf[0:n]}
		fmt.Println(m)
		m.SetQR(QR_RESPONSE)
		m.SetRCODE(RCODE_NOT_IMPLEMENTED)

		for _, q := range m.Questions() {
			fmt.Println(q.NameLabels(), q.Type(), q.Class())
		}

		sock.WriteToUDP(m.Data, clientAddr)
	}
}

func bits(b byte) string {
	return fmt.Sprintf("%08s", strconv.FormatInt(int64(b), 2))
}

func bitsAND(b byte, op byte) string {
	bs := bits(b)
	ops := bits(op)
	res := bits(b & op)
	return fmt.Sprintf("   %8s\n+  %8s\n=  %8s\n", bs, ops, res)
}

func bitsOR(b byte, op byte) string {
	bs := bits(b)
	ops := bits(op)
	res := bits(b | op)
	return fmt.Sprintf("   %8s\n|  %8s\n=  %8s\n", bs, ops, res)
}

func bitsXOR(b byte, op byte) string {
	bs := bits(b)
	ops := bits(op)
	res := bits(b ^ op)
	return fmt.Sprintf("   %8s\n^  %8s\n=  %8s\n", bs, ops, res)
}

func bitsRSHIFT(b byte, op byte) string {
	bs := bits(b)
	res := bits(b >> op)
	return fmt.Sprintf("   %8s\n>%d %8s\n", bs, op, res)
}

func bitsLSHIFT(b byte, op byte) string {
	bs := bits(b)
	res := bits(b << op)
	return fmt.Sprintf("   %8s\n<%d %8s\n", bs, op, res)
}
