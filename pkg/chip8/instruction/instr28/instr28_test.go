package instr28

import (
	"testing"

	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/handler"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
	"github.com/stretchr/testify/assert"
)

func TestExec(t *testing.T) {
	tests := []struct {
		name    string
		chip8   *chip8.Chip8
		args    *argument.OpcodeArguments
		handler *handler.Handler

		wantChip8State *chip8.Chip8
	}{
		{
			name:  "load null key",
			chip8: chip8.New(),
			args:  argument.ParseArguments(0xF10A),
			handler: handler.New(func() (int, error) {
				return 0x0, nil
			}),

			wantChip8State: chip8.New(),
		},
		{
			name:  "load last key",
			chip8: chip8.New(),
			args:  argument.ParseArguments(0xF10A),
			handler: handler.New(func() (int, error) {
				return 0xF, nil
			}),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x1] = 0xF

				return state
			}(),
		},
		{
			name:  "load middle key",
			chip8: chip8.New(),
			args:  argument.ParseArguments(0xF10A),
			handler: handler.New(func() (int, error) {
				return 0x8, nil
			}),

			wantChip8State: func() *chip8.Chip8 {
				state := chip8.New()
				state.Registers.V[0x1] = 0x8

				return state
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chip8.AddHandler(tt.chip8, tt.handler)
			chip8.AddHandler(tt.wantChip8State, tt.handler)

			Exec(tt.chip8, tt.args)
			assert.Equal(t, tt.wantChip8State, tt.chip8)
		})
	}
}
