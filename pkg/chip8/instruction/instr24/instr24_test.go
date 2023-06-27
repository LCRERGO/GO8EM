package instr24

import (
	"testing"

	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
	"github.com/LCRERGO/GO8EM/pkg/chip8/memory"
	"github.com/LCRERGO/GO8EM/pkg/chip8/screen"
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
			name: "draw null byte",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.I = 0x400

				return state
			}(),
			args: argument.ParseArguments(0xD121),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.I = 0x400

				return state
			}(),
		},
		{
			name: "draw unit byte",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.I = 0x400
				memory.Set(state.Memory, 0x400, 0x80)

				return state
			}(),
			args: argument.ParseArguments(0xD121),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.I = 0x400
				memory.Set(state.Memory, 0x400, 0x80)
				screen.Set(state.Screen, 0, 0)

				return state
			}(),
		},
		{
			name: "draw full byte",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.I = 0x400
				memory.Set(state.Memory, 0x400, 0xFF)

				return state
			}(),
			args: argument.ParseArguments(0xD121),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.I = 0x400
				memory.Set(state.Memory, 0x400, 0xFF)
				screen.DrawSprite(state.Screen, 0, 0, []byte{0xFF}, 1)

				return state
			}(),
		},
		{
			name: "draw half byte",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.I = 0x400
				memory.Set(state.Memory, 0x400, 0x0F)

				return state
			}(),
			args: argument.ParseArguments(0xD121),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.I = 0x400
				memory.Set(state.Memory, 0x400, 0x0F)
				screen.DrawSprite(state.Screen, 0, 0, []byte{0x0F}, 1)

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
