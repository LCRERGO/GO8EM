package instr31

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
			name: "add null value",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.I = 0x42

				return state
			}(),
			args: argument.ParseArguments(0xF11E),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.I = 0x42

				return state
			}(),
		},
		{
			name: "add last value",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.I = 0xFFF
				state.Registers.V[0x1] = 0x42

				return state
			}(),
			args: argument.ParseArguments(0xF11E),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.I = 0x041
				state.Registers.V[0x1] = 0x42

				return state
			}(),
		},
		{
			name: "add middle value",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.I = 0x7FF
				state.Registers.V[0x1] = 0x42

				return state

			}(),
			args: argument.ParseArguments(0xF11E),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.I = 0x841
				state.Registers.V[0x1] = 0x42

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
