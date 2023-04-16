package stack

import (
	"testing"

	"github.com/LCRERGO/GO8EM/pkg/chip8/register"
	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	tests := []struct {
		name      string
		stack     Stack
		registers register.RegisterFile
		address   uint16

		wantStackState     Stack
		wantRegistersState register.RegisterFile
	}{
		{
			name:  "push empty stack",
			stack: Stack{},
			registers: register.RegisterFile{
				SP: 0x00,
			},
			address: 0x0042,
			wantStackState: Stack{
				data: []uint16{0x0042},
			},
			wantRegistersState: register.RegisterFile{
				SP: 0x01,
			},
		},
		{
			name: "push a single address",
			stack: Stack{
				data: []uint16{0x0100},
			},
			registers: register.RegisterFile{
				SP: 0x01,
			},
			address: 0x0042,
			wantStackState: Stack{
				data: []uint16{
					0x0100,
					0x0042,
				},
			},
			wantRegistersState: register.RegisterFile{
				SP: 0x02,
			},
		},
		{
			name: "push the last address",
			stack: Stack{
				data: []uint16{
					0x0100,
					0x0200,
					0x0300,
					0x0400,
					0x0500,
					0x0600,
					0x0700,
					0x0800,
					0x0900,
					0x0A00,
					0x0B00,
					0x0C00,
					0x0D00,
					0x0E00,
					0x0F00,
				},
			},
			registers: register.RegisterFile{
				SP: 0x0F,
			},
			address: 0x0042,
			wantStackState: Stack{
				data: []uint16{
					0x0100,
					0x0200,
					0x0300,
					0x0400,
					0x0500,
					0x0600,
					0x0700,
					0x0800,
					0x0900,
					0x0A00,
					0x0B00,
					0x0C00,
					0x0D00,
					0x0E00,
					0x0F00,
					0x0042,
				},
			},
			wantRegistersState: register.RegisterFile{
				SP: 0x10,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Push(&tt.stack, &tt.registers, tt.address)

			assert.Equal(t, tt.wantStackState, tt.stack)
			assert.Equal(t, tt.wantRegistersState, tt.registers)
		})
	}
}

func TestPop(t *testing.T) {
	tests := []struct {
		name      string
		stack     Stack
		registers register.RegisterFile

		wantStackState     Stack
		wantRegistersState register.RegisterFile
		want               uint16
	}{
		{
			name: "pop the last address",
			stack: Stack{
				data: []uint16{
					0x0100,
				},
			},
			registers: register.RegisterFile{
				SP: 0x01,
			},
			wantStackState: Stack{
				data: make([]uint16, 0, 0x10),
			},
			wantRegistersState: register.RegisterFile{
				SP: 0x00,
			},
			want: 0x0100,
		},
		{
			name: "pop a single address",
			stack: Stack{
				data: []uint16{
					0x0100,
					0x0200,
				},
			},
			registers: register.RegisterFile{
				SP: 0x02,
			},
			wantStackState: Stack{
				data: []uint16{0x0100},
			},
			wantRegistersState: register.RegisterFile{
				SP: 0x01,
			},
			want: 0x0200,
		},
		{
			name: "pop full stack",
			stack: Stack{
				data: []uint16{
					0x0100,
					0x0200,
					0x0300,
					0x0400,
					0x0500,
					0x0600,
					0x0700,
					0x0800,
					0x0900,
					0x0A00,
					0x0B00,
					0x0C00,
					0x0D00,
					0x0E00,
					0x0F00,
					0x0042,
				},
			},
			registers: register.RegisterFile{
				SP: 0x10,
			},
			wantStackState: Stack{
				data: []uint16{
					0x0100,
					0x0200,
					0x0300,
					0x0400,
					0x0500,
					0x0600,
					0x0700,
					0x0800,
					0x0900,
					0x0A00,
					0x0B00,
					0x0C00,
					0x0D00,
					0x0E00,
					0x0F00,
				},
			},
			wantRegistersState: register.RegisterFile{
				SP: 0x0F,
			},
			want: 0x0042,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := Pop(&tt.stack, &tt.registers)

			assert.Equal(t, tt.wantRegistersState, tt.registers)
			assert.Equal(t, tt.wantStackState, tt.stack)
			assert.Equal(t, tt.want, value)
		})
	}
}
