package main

import (
	"bufio"
	"fmt"
	"net"

	"github.com/doctordesh/goodness"
	"github.com/doctordesh/goodness/bits"
)

func main() {
	g := goodness.New()
	msg := g.Build()

	msg.Data = []byte{0xA9, 0xA9, 0x01, 0x20, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x06, 0x67, 0x6F, 0x6F, 0x67, 0x6C, 0x65, 0x03, 0x63, 0x6F, 0x6D, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x29, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	// conn, err := net.Dial("udp", "1.1.1.1:53")
	conn, err := net.Dial("udp", "localhost:53")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}

	fmt.Println("Connection established")
	fmt.Println("Sending message:")

	fmt.Println(msg)
	fmt.Println()

	n, err := conn.Write(msg.Data)
	if n != len(msg.Data) {
		panic("did not send full message")
	}

	fmt.Println("Reading response")

	resp := make([]byte, 512)
	n, err = bufio.NewReader(conn).Read(resp)
	if err != nil {
		panic(fmt.Sprintf("could not read: %s", err.Error()))
	}
	conn.Close()

	msg = goodness.Message{Data: resp[0:n]}
	err = msg.Parse()
	if err != nil {
		panic(err)
	}

	fmt.Println(msg)
	fmt.Printf("[]byte{")
	for i, v := range msg.Data {
		if i%4 == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("0x%02X, ", v)
	}
	fmt.Printf("}\n")
	fmt.Println(bits.String(0xC0))
	fmt.Println(bits.String(0x0C))
	// fmt.Printf("\n%#v\n\n", msg.Data)

	for _, q := range msg.Questions {
		fmt.Println(q.NameLabels(), q.Type(), q.Class())
	}

	fmt.Println(";;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;")
	mask := byte(0b11000000)
	value := byte(0b11010100)
	fmt.Println(bits.AND(mask, value))
	fmt.Println(bits.String(0b010100), 0b010100)
	fmt.Println((value << 2) >> 2)
	fmt.Println(value ^ 0b00111111)
	fmt.Println(bits.XOR(value, 0b11000000))

	b := []byte{97, 99, 104, 106}
	s := string(b)
	fmt.Println(s, b)
	b[2] = 97
	fmt.Println(s, b)
}
