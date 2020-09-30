package reversevariant

import "io"

// WriteUint32 encodes a uint32 with variant-length encoding (values less than 0x80 take only one byte).
func WriteUint32(w io.Writer, x uint32) int {
	var buf [5]byte
	buf[0] = 0x7F & byte(x)
	n := 1
	for x >= 0x80 {
		x >>= 7
		buf[n] = 0x80 | byte(x)
		n++
	}
	w.Write(buf[:n])
	return n
}

// ReadUint32 reads a uint32 value encoded using WriteUint32 from the END of s.
// Returns the decoded uint32 and the number of bytes read, or panics otherwise.
func ReadUint32(s string) (uint32, int) {
	n := 1
	var x uint32
	lenS := len(s)
	b := s[lenS - n]
	for b >= 0x80 {
		x |= uint32(b) & 0x7F
		x <<= 7
		n++
		b = s[lenS - n]
	}
	x |= uint32(b)
	return x, n
}