package video

import (
	"log"

	"github.com/LCRERGO/GO8EM/pkg/chip8/screen"
	"github.com/LCRERGO/GO8EM/pkg/constants"
	"github.com/veandco/go-sdl2/sdl"
)

type VideoSubsystem struct {
	Window   *sdl.Window
	Renderer *sdl.Renderer
}

func New() *VideoSubsystem {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		log.Fatal("new_video_subsystem: ", err)
	}

	window, err := sdl.CreateWindow(constants.WindowTitle,
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		constants.ScreenWidth*constants.ScreenScaleFactor,
		constants.ScreenHeight*constants.ScreenScaleFactor,
		sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal("new_video_subsystem: ", err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.TEXTUREACCESS_TARGET)
	if err != nil {
		log.Fatal("new_video_subsystem: ", err)
	}

	return &VideoSubsystem{
		Window:   window,
		Renderer: renderer,
	}
}

func Render(video *VideoSubsystem, screenData *screen.Screen) {
	err := video.Renderer.SetDrawColor(0, 0, 0, sdl.ALPHA_OPAQUE)
	if err != nil {
		log.Fatal("render_video_subsystem: ", err)
	}
	err = video.Renderer.Clear()
	if err != nil {
		log.Fatal("render_video_subsystem: ", err)
	}
	err = video.Renderer.SetDrawColor(255, 255, 255, sdl.ALPHA_OPAQUE)
	if err != nil {
		log.Fatal("render_video_subsystem: ", err)
	}

	for x := 0; x < constants.ScreenWidth; x++ {
		for y := 0; y < constants.ScreenHeight; y++ {
			if screen.IsSet(screenData, x, y) {
				r := sdl.Rect{
					X: int32(x * constants.ScreenScaleFactor),
					Y: int32(y * constants.ScreenScaleFactor),
					W: constants.ScreenScaleFactor,
					H: constants.ScreenScaleFactor,
				}

				video.Renderer.FillRect(&r)
			}
		}
	}

	video.Renderer.Present()
}

func Destroy(video *VideoSubsystem) {
	video.Renderer.Destroy()
	video.Window.Destroy()
}
