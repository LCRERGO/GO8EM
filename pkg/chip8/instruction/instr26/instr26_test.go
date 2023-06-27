package instr26

import (
	"testing"

	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
	"github.com/LCRERGO/GO8EM/pkg/chip8/keyboard"
	"github.com/LCRERGO/GO8EM/pkg/chip8/register"
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
			name: "skip instruction",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				key := state.Registers.V[0x1]
				keyboard.SetKeyDown(state.Keyboard, int(key))

				return state
			}(),
			args: argument.ParseArguments(0xE1A1),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				key := state.Registers.V[0x1]
				keyboard.SetKeyDown(state.Keyboard, int(key))

				return state
			}(),
		},
		{
			name:  "no skip instruction",
			chip8: chip8.New(),
			args:  argument.ParseArguments(0xE1A1),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				register.NextInstruction(state.Registers)

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
