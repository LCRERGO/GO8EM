package instr33

import (
	"testing"

	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
	"github.com/LCRERGO/GO8EM/pkg/chip8/memory"
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
			name: "store null byte",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.I = 0x400

				return state
			}(),
			args: argument.ParseArguments(0xF133),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.I = 0x400

				return state
			}(),
		},
		{
			name: "store full byte",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x1] = 0xFF
				state.Registers.I = 0x400

				return state
			}(),
			args: argument.ParseArguments(0xF133),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x1] = 0xFF
				state.Registers.I = 0x400

				memory.Set(state.Memory, int(state.Registers.I), 0x02)
				memory.Set(state.Memory, int(state.Registers.I+1), 0x05)
				memory.Set(state.Memory, int(state.Registers.I+2), 0x05)

				return state
			}(),
		},
		{
			name: "store half byte",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x1] = 0x0F
				state.Registers.I = 0x400

				return state
			}(),
			args: argument.ParseArguments(0xF133),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x1] = 0x0F
				state.Registers.I = 0x400

				memory.Set(state.Memory, int(state.Registers.I), 0x00)
				memory.Set(state.Memory, int(state.Registers.I+1), 0x01)
				memory.Set(state.Memory, int(state.Registers.I+2), 0x05)

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
