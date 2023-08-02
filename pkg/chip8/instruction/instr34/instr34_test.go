package instr34

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
			name: "store single register",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.I = 0x400
				state.Registers.V[0x0] = 0x42

				return state
			}(),
			args: argument.ParseArguments(0xF055),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.I = 0x400
				state.Registers.V[0x0] = 0x42
				memory.Set(state.Memory,
					int(state.Registers.I),
					uint8(state.Registers.V[0x0]))
				state.Registers.PC += 2

				return state
			}(),
		},
		{
			name: "store all registers",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.I = 0x400
				for i := range state.Registers.V {
					state.Registers.V[i] = uint16(i)
				}

				return state
			}(),
			args: argument.ParseArguments(0xFF55),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.I = 0x400
				for i := range state.Registers.V {
					state.Registers.V[i] = uint16(i)
					memory.Set(state.Memory, 0x400+i, uint8(i))
				}
				state.Registers.PC += 2

				return state
			}(),
		},
		{
			name: "add middle value",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.I = 0x400
				for i := 0; i <= 0x7; i++ {
					state.Registers.V[i] = uint16(i)
				}

				return state

			}(),
			args: argument.ParseArguments(0xF755),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.I = 0x400
				for i := 0; i <= 0x7; i++ {
					state.Registers.V[i] = uint16(i)
					memory.Set(state.Memory, 0x400+i, uint8(i))
				}
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
