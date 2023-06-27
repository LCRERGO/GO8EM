package instr22

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
			name:  "load first address (V0 not loaded)",
			chip8: chip8.New(),
			args:  argument.ParseArguments(0xB200),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.PC = 0x200

				return state
			}(),
		},
		{
			name:  "load last address (V0 not loaded)",
			chip8: chip8.New(),
			args:  argument.ParseArguments(0xBFFF),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.PC = 0xFFF

				return state
			}(),
		},
		{
			name:  "load middle address (V0 not loaded)",
			chip8: chip8.New(),
			args:  argument.ParseArguments(0xB7FF),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.PC = 0x7FF

				return state
			}(),
		},
		{
			name: "load first address (V0 loaded)",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x00] = 0x42

				return state
			}(),
			args: argument.ParseArguments(0xB200),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x00] = 0x42
				state.Registers.PC = 0x242

				return state
			}(),
		},
		{
			name: "load last address (V0 loaded)",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x00] = 0x42

				return state
			}(),
			args: argument.ParseArguments(0xBFFF),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x00] = 0x42
				state.Registers.PC = 0x41

				return state
			}(),
		},
		{
			name: "load middle address (V0 loaded)",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x00] = 0x42

				return state
			}(),
			args: argument.ParseArguments(0xB7FF),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x00] = 0x42
				state.Registers.PC = 0x841

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
