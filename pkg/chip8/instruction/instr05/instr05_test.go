package instr05

import (
	"testing"

	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
	"github.com/LCRERGO/GO8EM/pkg/chip8/stack"
	"github.com/stretchr/testify/assert"
)

func TestExec(t *testing.T) {
	tests := []struct {
		name  string
		chip8 *chip8.Chip8
		args  *argument.OpcodeArguments

		wantChip8State *chip8.Chip8
	}{
		{
			name: "call first address",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.PC = 0x342

				return state
			}(),
			args: argument.ParseArguments(0x2200),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				stack.Push(state.Stack, state.Registers, 0x344)
				state.Registers.PC = 0x200

				return state
			}(),
		},
		{
			name: "call last address",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.PC = 0x342

				return state
			}(),
			args: argument.ParseArguments(0x2FFF),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				stack.Push(state.Stack, state.Registers, 0x344)
				state.Registers.PC = 0xFFF

				return state
			}(),
		},
		{
			name: "call middle address",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.PC = 0x342

				return state
			}(),
			args: argument.ParseArguments(0x26FF),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				stack.Push(state.Stack, state.Registers, 0x344)
				state.Registers.PC = 0x6FF

				return state
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Exec(tt.chip8, tt.args)
			assert.Equal(t, tt.wantChip8State, tt.chip8)
		})
	}
}
