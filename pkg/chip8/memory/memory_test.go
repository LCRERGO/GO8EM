package memory

import (
	"testing"

	"github.com/LCRERGO/GO8EM/pkg/constants"
	"github.com/stretchr/testify/assert"
)

func TestLoadROM(t *testing.T) {
	tests := []struct {
		name   string
		memory *Memory
		rom    []byte
		size   uint

		wantMemoryState *Memory
	}{
		{
			name:   "load an empty rom",
			memory: New(),
			rom:    []byte{},
			size:   0,

			wantMemoryState: New(),
		},
		{
			name:   "load a rom of 1 byte",
			memory: New(),
			rom:    []byte{0xFF},
			size:   1,

			wantMemoryState: func() *Memory {
				baseAddress := uint(constants.ProgramStartAddress)
				state := New()
				copy(state.data[baseAddress:baseAddress+1], []byte{0xFF})

				return state
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LoadROM(tt.memory, tt.rom, tt.size)
			assert.Equal(t, tt.wantMemoryState, tt.memory)
		})
	}
}

func TestSet(t *testing.T) {
	tests := []struct {
		name   string
		memory *Memory
		index  int
		value  uint8

		wantMemoryState *Memory
	}{
		{
			name:   "set at begin",
			memory: New(),
			index:  0x000,
			value:  0x42,

			wantMemoryState: func() *Memory {
				state := New()
				copy(state.data, []byte{0x42})

				return state
			}(),
		},
		{
			name:   "set at end",
			memory: New(),
			index:  0xFFF,
			value:  0x42,

			wantMemoryState: func() *Memory {
				state := New()
				copy(state.data[0xFFF:0x1000], []byte{0x42})

				return state
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Set(tt.memory, tt.index, tt.value)
			assert.Equal(t, tt.wantMemoryState, tt.memory)
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name   string
		memory *Memory
		index  int

		wantMemoryState *Memory
		want            uint8
	}{
		{
			name: "get at begin",
			memory: func() *Memory {
				state := New()
				copy(state.data, []byte{0x42})

				return state
			}(),
			index: 0x000,

			wantMemoryState: func() *Memory {
				state := New()
				copy(state.data, []byte{0x42})

				return state
			}(),
			want: 0x42,
		},
		{
			name: "get at end",
			memory: func() *Memory {
				state := New()
				copy(state.data[0xFFF:0x1000], []byte{0x42})

				return state
			}(),
			index: 0xFFF,

			wantMemoryState: func() *Memory {
				state := New()
				copy(state.data[0xFFF:0x1000], []byte{0x42})

				return state
			}(),
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
		memory *Memory
		index  int

		wantMemoryState *Memory
		want            uint16
	}{
		{
			name: "get at begin",
			memory: func() *Memory {
				state := New()
				copy(state.data, []byte{0xDE, 0xAD})

				return state
			}(),
			index: 0x000,

			wantMemoryState: func() *Memory {
				state := New()
				copy(state.data, []byte{0xDE, 0xAD})

				return state
			}(),
			want: 0xDEAD,
		},
		{
			name: "get at end",
			memory: func() *Memory {
				state := New()
				copy(state.data[0xFFE:0x1000], []byte{0xBE, 0xEF})

				return state
			}(),
			index: 0xFFE,

			wantMemoryState: func() *Memory {
				state := New()
				copy(state.data[0xFFE:0x1000], []byte{0xBE, 0xEF})

				return state
			}(),
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
