package video

import (
	"os"
	"testing"
	"time"

	"github.com/LCRERGO/GO8EM/pkg/chip8/screen"
	"github.com/LCRERGO/GO8EM/pkg/constants"
)

func TestDisplay(t *testing.T) {
	if os.Getenv("TEST_ALL") != "" ||
		os.Getenv("TEST_GRAPHICS") != "" {
		video := NewDisplay()
		defer Destroy(&video)

		screenData := screen.NewScreen()
		screen.DrawSprite(screenData, constants.ScreenWidth/2-2, constants.ScreenHeight/2-2,
			[]byte{0xF0, 0x90, 0x90, 0x90, 0xF0}, 5)
		Render(&video, screenData)

		time.Sleep(1 * time.Second)
	}
}
