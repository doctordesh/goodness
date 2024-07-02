package goodness

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type Goodness struct {
	Host string
	Port int
}

func New(host string, port int) Goodness {
	return Goodness{Host: host, Port: port}
}

func (g Goodness) ConnectionString() string {
	return fmt.Sprintf("%s:%d", g.Host, g.Port)
}

// Serve launches the DNS cache server
func (g Goodness) Serve() error {
	var err error
	addr, err := net.ResolveUDPAddr("udp", g.ConnectionString())
	if err != nil {
		return err
	}

	sock, err := net.ListenUDP("udp", addr)
	if err != nil {
		return fmt.Errorf("could not setup socket for %v: %w", addr, err)
	}
	sock.SetReadBuffer(100_000)

	jobs := make(chan udpMessage)
	go sender(sock, jobs)

	log.Printf("starting server, listening on %v\n", addr)
	for {
		buf := make([]byte, 512)
		n, clientAddr, err := sock.ReadFromUDP(buf)
		if err != nil {
			panic(err)
		}

		go g.handleRequest(buf[0:n], clientAddr, jobs)
	}
}

func (g Goodness) handleRequest(payload []byte, client *net.UDPAddr, replyChan chan<- udpMessage) {
	// In order to have this running always, there must exist some
	// heuristics that tell if this is a request that 'goodness' can handle
	// or if the request should just be piped through. Over time when more
	// features are implemented these rules will change.

	// Parse messge
	msg := Message{Data: payload}
	err := msg.Parse()
	if err != nil {
		panic(err)
	}

	if true { // very simple heuristic
		g.pipeThrough(payload, client, replyChan)
	} else {
		// - Parse the message
		// - Find the 'Question' in question
		// - Does the domain exist in the local store?
		// - If no, build a new request and as another DNS server for the details
		//   - Store it in the local store
		// - Reply with the info

	}

	for _, q := range msg.Questions {
		fmt.Printf("Type %-8s request for %s: %s\n", q.Type(), q.Class(), strings.Join(q.NameLabels(), "."))
	}

	// for i := 0; i < len(payload); i++ {
	// 	fmt.Printf("0x%02X, ", payload[i])
	// }

	// fmt.Println()

	// m := Message{Data: payload}
	// m.Parse()
	// fmt.Println(m)
	// m.SetQR(QR_RESPONSE)
	// m.SetRCODE(RCODE_NOT_IMPLEMENTED)

	// // for _, q := range m.Questions() {
	// // 	fmt.Println(q.NameLabels(), q.Type(), q.Class())
	// // }

	// // fmt.Println(m)
	// for _, q := range m.Questions {
	// 	fmt.Println(q.NameLabels(), q.Type(), q.Class())
	// }

	// replyChan <- udpMessage{receiver: client, payload: m.Data}
}

type udpMessage struct {
	receiver *net.UDPAddr
	payload  []byte
}

func sender(sock *net.UDPConn, jobs <-chan udpMessage) {
	log.Printf("starting the sender goroutine")
	for job := range jobs {
		// log.Printf("replying to %v", job.receiver)
		sock.WriteToUDP(job.payload, job.receiver)
	}
}

func (g Goodness) pipeThrough(payload []byte, client *net.UDPAddr, replyChan chan<- udpMessage) {
	var err error
	var n int

	conn, err := net.Dial("udp", "1.1.1.1:53")
	if err != nil {
		panic(err)
	}

	n, err = conn.Write(payload)
	if n != len(payload) {
		panic(fmt.Sprintf("did not send full message: %d/%d", n, len(payload)))
	}

	resp := make([]byte, 512)
	n64, err := conn.Read(resp)
	if err != nil {
		panic(fmt.Sprintf("could not read: %s", err.Error()))
	}
	conn.Close()

	replyChan <- udpMessage{client, resp[0:n64]}
}
