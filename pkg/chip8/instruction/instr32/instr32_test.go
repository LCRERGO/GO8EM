package instr32

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
			name: "load first sprite",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x1] = 0x00

				return state
			}(),
			args: argument.ParseArguments(0xF129),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.I = 0x000
				state.Registers.V[0x1] = 0x00

				return state
			}(),
		},
		{
			name: "load last sprite",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x1] = 0xF

				return state
			}(),
			args: argument.ParseArguments(0xF129),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.I = 0x4B
				state.Registers.V[0x1] = 0xF

				return state
			}(),
		},
		{
			name: "add middle value",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x1] = 0x7

				return state

			}(),
			args: argument.ParseArguments(0xF129),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.I = 0x23
				state.Registers.V[0x1] = 0x7

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
