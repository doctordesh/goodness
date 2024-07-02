package dns

import (
	"fmt"
	"strings"
)

type ResourceRecord struct {
	Domain     []string
	Type       Type
	Class      Class
	TimeToLive int32
	Data       []byte
}

func (r ResourceRecord) EncodeLabels() []byte {
	data := []byte{}
	for _, part := range r.Domain {
		data = append(data, byte(len(part)))
		data = append(data, []byte(part)...)
	}

	return data
}

func (r ResourceRecord) String() string {
	return fmt.Sprintf("%s: Type: %s, Class: %s, TTL: %d", strings.Join(r.Domain, "."), r.Type, r.Class, r.TimeToLive)
}
