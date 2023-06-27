package instr29

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
			name:  "load null byte",
			chip8: chip8.New(),
			args:  argument.ParseArguments(0xF115),

			wantChip8State: chip8.New(),
		},
		{
			name: "load last byte",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x1] = 0xFF

				return state
			}(),
			args: argument.ParseArguments(0xF115),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x1] = 0xFF
				state.Registers.DT = 0xFF

				return state
			}(),
		},
		{
			name: "load half byte",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x1] = 0x0F

				return state
			}(),
			args: argument.ParseArguments(0xF115),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x1] = 0x0F
				state.Registers.DT = 0x0F

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
