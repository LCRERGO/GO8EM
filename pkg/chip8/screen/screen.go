package screen

import "github.com/LCRERGO/GO8EM/pkg/constants"

type Screen struct {
	pixels [constants.ScreenWidth][constants.ScreenHeight]bool
}

func NewScreen() *Screen {
	return &Screen{}
}

func Clear(screen *Screen) {
	screen.pixels = [constants.ScreenWidth][constants.ScreenHeight]bool{}
}

func Set(screen *Screen, x, y int) {
	screen.pixels[x][y] = true
}

func IsSet(screen *Screen, x, y int) bool {
	return screen.pixels[x][y]
}

func DrawSprite(screen *Screen, x, y int, sprite []byte, size int) bool {
	var pixelCollision bool

	for cy := 0; cy < size; cy++ {
		sb := sprite[cy]
		for cx := 0; cx < 8; cx++ {
			if (sb & (0x80 >> cx)) == 0 {
				continue
			}

			currentPixel := screen.pixels[(cx+x)%constants.ScreenWidth][(cy+y)%constants.ScreenHeight]
			if currentPixel {
				pixelCollision = true
			}
			screen.pixels[(cx+x)%constants.ScreenWidth][(cy+y)%constants.ScreenHeight] =
				(true || currentPixel) && !(true && currentPixel)
		}
	}

	return pixelCollision
}
