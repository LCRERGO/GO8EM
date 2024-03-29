package instr07

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
			name: "skip not null byte",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x1] = 0x42

				return state
			}(),
			args: argument.ParseArguments(0x4100),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x1] = 0x42
				state.Registers.PC += 4

				return state
			}(),
		},
		{
			name:  "no skip null byte",
			chip8: chip8.New(),
			args:  argument.ParseArguments(0x4100),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.PC += 2

				return state
			}(),
		},
		{
			name:  "skip full byte",
			chip8: chip8.New(),
			args:  argument.ParseArguments(0x41FF),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.PC += 4

				return state
			}(),
		},
		{
			name: "no skip full byte",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x1] = 0xFF

				return state
			}(),
			args: argument.ParseArguments(0x41FF),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x1] = 0xFF
				state.Registers.PC += 2

				return state
			}(),
		},
		{
			name:  "skip half byte",
			chip8: chip8.New(),
			args:  argument.ParseArguments(0x417F),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.PC += 4

				return state
			}(),
		},
		{
			name: "no skip half byte",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x1] = 0x7F

				return state
			}(),
			args: argument.ParseArguments(0x417F),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
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
