package stack

import (
	"testing"

	"github.com/LCRERGO/GO8EM/pkg/chip8/register"
	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	tests := []struct {
		name      string
		stack     *Stack
		registers *register.RegisterFile
		address   uint16

		wantStackState     *Stack
		wantRegistersState *register.RegisterFile
	}{
		{
			name:      "push empty stack",
			stack:     New(),
			registers: register.New(),
			address:   0x0042,
			wantStackState: func() *Stack {
				state := New()
				copy(state.data, []uint16{0x0042})

				return state
			}(),
			wantRegistersState: func() *register.RegisterFile {
				registersState := register.New()
				registersState.SP = 0x01

				return registersState
			}(),
		},
		{
			name: "push single address",
			stack: func() *Stack {
				state := New()
				copy(state.data, []uint16{0x0100})

				return state
			}(),
			registers: func() *register.RegisterFile {
				registerState := register.New()
				registerState.SP = 0x01

				return registerState
			}(),
			address: 0x0042,
			wantStackState: func() *Stack {
				state := New()
				copy(state.data, []uint16{0x0100, 0x0042})

				return state
			}(),
			wantRegistersState: func() *register.RegisterFile {
				registerState := register.New()
				registerState.SP = 0x02

				return registerState
			}(),
		},
		{
			name: "push last address",
			stack: func() *Stack {
				state := New()
				copy(state.data, []uint16{
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
				})

				return state
			}(),

			registers: func() *register.RegisterFile {
				registerState := register.New()
				registerState.SP = 0x0F

				return registerState
			}(),
			address: 0x0042,
			wantStackState: func() *Stack {
				state := New()
				copy(state.data, []uint16{
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
				})

				return state
			}(),
			wantRegistersState: func() *register.RegisterFile {
				registerState := register.New()
				registerState.SP = 0x10

				return registerState
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Push(tt.stack, tt.registers, tt.address)

			assert.Equal(t, tt.wantStackState, tt.stack)
			assert.Equal(t, tt.wantRegistersState, tt.registers)
		})
	}
}

func TestPop(t *testing.T) {
	tests := []struct {
		name      string
		stack     *Stack
		registers *register.RegisterFile

		wantStackState     *Stack
		wantRegistersState *register.RegisterFile
		want               uint16
	}{
		{
			name: "pop last address",
			stack: func() *Stack {
				state := New()
				copy(state.data, []uint16{0x0100})

				return state
			}(),
			registers: func() *register.RegisterFile {
				registerState := register.New()
				registerState.SP = 0x01

				return registerState
			}(),
			wantStackState:     New(),
			wantRegistersState: register.New(),
			want:               0x0100,
		},
		{
			name: "pop single address",
			stack: func() *Stack {
				state := New()
				copy(state.data, []uint16{0x0100, 0x0200})

				return state
			}(),
			registers: func() *register.RegisterFile {
				registerState := register.New()
				registerState.SP = 0x02

				return registerState
			}(),
			wantStackState: func() *Stack {
				state := New()
				copy(state.data, []uint16{0x0100})

				return state
			}(),
			wantRegistersState: func() *register.RegisterFile {
				registerState := register.New()
				registerState.SP = 0x01

				return registerState
			}(),
			want: 0x0200,
		},
		{
			name: "pop full stack",
			stack: func() *Stack {
				state := New()
				copy(state.data, []uint16{
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
				})

				return state
			}(),
			registers: func() *register.RegisterFile {
				registerState := register.New()
				registerState.SP = 0x10

				return registerState
			}(),
			wantStackState: func() *Stack {
				state := New()
				copy(state.data, []uint16{
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
				})

				return state
			}(),
			wantRegistersState: func() *register.RegisterFile {
				registerState := register.New()
				registerState.SP = 0x0F

				return registerState
			}(),
			want: 0x0042,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := Pop(tt.stack, tt.registers)

			assert.Equal(t, tt.wantRegistersState, tt.registers)
			assert.Equal(t, tt.wantStackState, tt.stack)
			assert.Equal(t, tt.want, value)
		})
	}
}
