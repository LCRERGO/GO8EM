package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	tests := []struct {
		name   string
		memory Memory
		index  int
		value  uint8

		wantMemoryState Memory
	}{
		{
			name:   "set at begin",
			memory: Memory{},
			index:  0x000,
			value:  0x42,

			wantMemoryState: Memory{
				data: [0x1000]byte{0x42},
			},
		},
		{
			name:   "set at end",
			memory: Memory{},
			index:  0xFFF,
			value:  0x42,

			wantMemoryState: Memory{
				data: [0x1000]byte{0xFFF: 0x42},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Set(&tt.memory, tt.index, tt.value)
			assert.Equal(t, tt.wantMemoryState, tt.memory)
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name   string
		memory Memory
		index  int

		wantMemoryState Memory
		want            uint8
	}{
		{
			name: "get at begin",
			memory: Memory{
				data: [0x1000]byte{0x42},
			},
			index: 0x000,

			wantMemoryState: Memory{
				data: [0x1000]byte{0x42},
			},
			want: 0x42,
		},
		{
			name: "get at begin",
			memory: Memory{
				data: [0x1000]byte{0xFFF: 0x42},
			},
			index: 0xFFF,

			wantMemoryState: Memory{
				data: [0x1000]byte{0xFFF: 0x42},
			},
			want: 0x42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := Get(tt.memory, tt.index)

			assert.Equal(t, tt.wantMemoryState, tt.memory)
			assert.Equal(t, tt.want, value)
		})
	}
}

func TestGet16(t *testing.T) {
	tests := []struct {
		name   string
		memory Memory
		index  int

		wantMemoryState Memory
		want            uint16
	}{
		{
			name: "get at begin",
			memory: Memory{
				data: [0x1000]byte{0xDE, 0xAD},
			},
			index: 0x000,

			wantMemoryState: Memory{
				data: [0x1000]byte{0xDE, 0xAD},
			},
			want: 0xDEAD,
		},
		{
			name: "get at begin",
			memory: Memory{
				data: [0x1000]byte{0xFFE: 0xBE, 0xFFF: 0xEF},
			},
			index: 0xFFE,

			wantMemoryState: Memory{
				data: [0x1000]byte{0xFFE: 0xBE, 0xFFF: 0xEF},
			},
			want: 0xBEEF,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := Get16(tt.memory, tt.index)

			assert.Equal(t, tt.memory, tt.wantMemoryState)
			assert.Equal(t, tt.want, value)
		})
	}

}
