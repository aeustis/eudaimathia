package reversevariant_test

import (
	"fmt"
	"math/bits"
	"strings"
	"testing"

	"github.com/eudaimathia/src/reversevariant"
	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	expectedSize := func(x uint32) int {
		if x == 0 {
			return 1
		}
		return 1 + (31-bits.LeadingZeros32(x))/7
	}
	tests := []uint32{0, 1, 2, 3, 4, 5, 0x7F, 0x80, 0x2000, 1E3, 1E6, 1E9, 1<<32 - 1}
	for _, x := range tests {
		t.Run(fmt.Sprintf("%d", x), func(t *testing.T) {
			var buf strings.Builder
			n := reversevariant.WriteUint32(&buf, x)
			assert.Equal(t, expectedSize(x), n)
			s := buf.String()
			assert.Len(t, s, n)
			x2, n2 := reversevariant.ReadUint32("garbage" + s)
			assert.Equal(t, x, x2)
			assert.Equal(t, n, n2)
		})
	}
}
