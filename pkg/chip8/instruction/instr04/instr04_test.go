package instr04

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
			name:  "jump to first address",
			chip8: chip8.New(),
			args:  argument.ParseArguments(0x1200),

			wantChip8State: chip8.New(),
		},
		{
			name:  "jump to last address",
			chip8: chip8.New(),
			args:  argument.ParseArguments(0x1FFF),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.PC = 0xFFF

				return state
			}(),
		},
		{
			name:  "jump to middle address",
			chip8: chip8.New(),
			args:  argument.ParseArguments(0x16FF),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
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
