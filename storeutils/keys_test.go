package storeutils

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tendermint/dex-demo/testutil/testflags"
)

func TestPrefixKey(t *testing.T) {
	testflags.UnitTest(t)
	out1 := PrefixKeyString("fooprefix")
	assert.Equal(t, "fooprefix", string(out1))
	out2 := PrefixKeyString("fooprefix", []byte("sub1"), []byte("sub2"))
	assert.Equal(t, "fooprefix/sub1/sub2", string(out2))
}

func TestUint64AndInt64Subkey(t *testing.T) {
	tests := []struct {
		in  int64
		out []byte
	}{
		{
			0,
			[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
		{
			1000,
			[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xe8},
		},
		{
			math.MaxInt64,
			[]byte{0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.out, Int64Subkey(tt.in))
		assert.Equal(t, tt.out, Uint64Subkey(uint64(tt.in)))
	}

	assert.Equal(t, []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, Uint64Subkey(math.MaxUint64))
}
