package instr23

import (
	"testing"

	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
	"github.com/LCRERGO/GO8EM/pkg/utils/lcg"
	"github.com/stretchr/testify/assert"
)

func TestExec(t *testing.T) {

	tests := []struct {
		name  string
		chip8 *chip8.Chip8
		args  *argument.OpcodeArguments
		seed  int

		wantChip8State *chip8.Chip8
	}{
		{
			name: "null seed null mask",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				randGen := lcg.New(0x00)
				chip8.AddRandGen(state, randGen)

				return state
			}(),
			args: argument.ParseArguments(0xC100),
			seed: 0x00,

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				randGen := lcg.New(0x0B)
				chip8.AddRandGen(state, randGen)
				state.Registers.PC += 2

				return state
			}(),
		},
		{
			name: "null seed full mask",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				randGen := lcg.New(0x00)
				chip8.AddRandGen(state, randGen)

				return state
			}(),
			args: argument.ParseArguments(0xC1FF),
			seed: 0x00,

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				randGen := lcg.New(0x0B)
				chip8.AddRandGen(state, randGen)
				state.Registers.V[0x1] = 0x0B
				state.Registers.PC += 2

				return state
			}(),
		},
		{
			name: "full seed null mask",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				randGen := lcg.New(0xFF)
				chip8.AddRandGen(state, randGen)

				return state
			}(),
			args: argument.ParseArguments(0xC100),
			seed: 0xFF,

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				randGen := lcg.New(0x5D90DF9869E)
				chip8.AddRandGen(state, randGen)
				state.Registers.PC += 2

				return state
			}(),
		},
		{
			name: "full seed full mask",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				randGen := lcg.New(0xFF)
				chip8.AddRandGen(state, randGen)

				return state
			}(),
			args: argument.ParseArguments(0xC1FF),
			seed: 0xFF,

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				randGen := lcg.New(0x5D90DF9869E)
				chip8.AddRandGen(state, randGen)
				state.Registers.V[0x1] = 0x9E
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
