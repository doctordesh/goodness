package know

type ResourceRecord struct {
	Data  []byte
	Start int
}

func Do(s string) []byte {
	// s += "sdf"
	// return []byte(s)
	// b := []byte{0, 1, 2, 3, 4, 5, 6, 0, 1, 2, 3, 4, 5, 6, 0, 1, 2, 3, 4, 5, 6}
	// s := string(b)
	// b[2] = 99

	// r := ResourceRecord{Data: b[0:17], Start: 4}
	// _, _ = r, s

	// b = make([]byte, 4)
	// sdf := new(ResourceRecord)
	// _ = sdf
	// return b
}
