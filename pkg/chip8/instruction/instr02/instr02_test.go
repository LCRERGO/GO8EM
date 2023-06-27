package instr02

import (
	"testing"

	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
	"github.com/LCRERGO/GO8EM/pkg/chip8/screen"
	"github.com/LCRERGO/GO8EM/pkg/constants"
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
			name:  "clear an empty screen",
			chip8: chip8.New(),
			args:  argument.ParseArguments(0x00E0),

			wantChip8State: chip8.New(),
		},
		{
			name: "clear an screen with 1 pixel",
			chip8: func() *chip8.Chip8 {
				state := chip8.New()
				screen.DrawSprite(state.Screen,
					constants.ScreenWidth/2, constants.ScreenHeight/2,
					[]byte{0x80}, 1)

				return state
			}(),
			args: argument.ParseArguments(0x00E0),

			wantChip8State: chip8.New(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Exec(tt.chip8, tt.args)
			assert.Equal(t, tt.wantChip8State, tt.chip8)
		})
	}
}
