package instr21

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
			name:  "load first address",
			chip8: chip8.New(),
			args:  argument.ParseArguments(0xA200),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.I = 0x200
				state.Registers.PC += 2

				return state
			}(),
		},
		{
			name:  "load last address",
			chip8: chip8.New(),
			args:  argument.ParseArguments(0xAFFF),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.I = 0xFFF
				state.Registers.PC += 2

				return state
			}(),
		},
		{
			name:  "load middle address",
			chip8: chip8.New(),
			args:  argument.ParseArguments(0xA7FF),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.I = 0x7FF
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
