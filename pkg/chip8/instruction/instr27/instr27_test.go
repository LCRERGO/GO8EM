package instr27

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
			name: "load null byte",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.DT = 0x00

				return state
			}(),
			args: argument.ParseArguments(0xF107),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.DT = 0x00
				state.Registers.V[0x1] = 0x00
				state.Registers.PC += 2

				return state
			}(),
		},
		{
			name: "load last byte",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.DT = 0xFF

				return state
			}(),
			args: argument.ParseArguments(0xF107),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.DT = 0xFF
				state.Registers.V[0x1] = 0xFF
				state.Registers.PC += 2

				return state
			}(),
		},
		{
			name: "load middle byte",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.DT = 0x7F

				return state
			}(),
			args: argument.ParseArguments(0xF107),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.DT = 0x7F
				state.Registers.V[0x1] = 0x7F
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
