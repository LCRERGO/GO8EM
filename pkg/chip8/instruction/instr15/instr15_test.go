package instr15

import (
	"testing"

	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
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
			name: "add null byte",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x1] = 0x42
				state.Registers.V[0x2] = 0x00

				return state
			}(),
			args: argument.ParseArguments(0x8124),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x1] = 0x42
				state.Registers.V[0x2] = 0x00
				state.Registers.PC += 2

				return state
			}(),
		},
		{
			name: "add last byte",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x1] = 0x42
				state.Registers.V[0x2] = 0xFF

				return state
			}(),
			args: argument.ParseArguments(0x8124),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x1] = 0x41
				state.Registers.V[0x2] = 0xFF
				state.Registers.V[0xF] = 0x01
				state.Registers.PC += 2

				return state
			}(),
		},
		{
			name: "add half byte",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x1] = 0x42
				state.Registers.V[0x2] = 0x0F

				return state

			}(),
			args: argument.ParseArguments(0x8124),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x1] = 0x51
				state.Registers.V[0x2] = 0x0F
				state.Registers.PC += 2

				return state
			}(),
		},
		{
			name: "add dijunct byte",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x1] = 0x42
				state.Registers.V[0x2] = 0x09

				return state

			}(),
			args: argument.ParseArguments(0x8124),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x1] = 0x4B
				state.Registers.V[0x2] = 0x09
				state.Registers.PC += 2

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
