// Package screen provides functions for screen manipulation.
package screen

import "github.com/LCRERGO/GO8EM/pkg/constants"

// Screen entity is the logical representation of the pixels
// in the emulator.
type Screen struct {
	pixels [][]bool
}

// Create a new Screen.
func New() *Screen {
	pixels := make([][]bool, constants.ScreenWidth)
	for i := range pixels {
		pixels[i] = make([]bool, constants.ScreenHeight)
	}

	return &Screen{
		pixels: pixels,
	}
}

// Deep Copy a Screen.
func Copy(screen *Screen) *Screen {
	pixels := make([][]bool, constants.ScreenWidth)
	for i := range pixels {
		pixels[i] = make([]bool, constants.ScreenHeight)
		copy(pixels[i], screen.pixels[i])
	}

	return &Screen{
		pixels: pixels,
	}
}

// Destroy a Screen.
func Destroy(screen *Screen) {
	for i := range screen.pixels {
		screen.pixels[i] = nil
	}
	screen.pixels = nil
	screen = nil
}

// Clear a Screen.
func Clear(screen *Screen) {
	for i := 0; i < constants.ScreenWidth; i++ {
		for j := 0; j < constants.ScreenHeight; j++ {
			screen.pixels[i][j] = false
		}
	}
	// screen.pixels = [constants.ScreenWidth][constants.ScreenHeight]bool{}
}

// Set a pixel on a Screen given it's x and y coordinates.
func Set(screen *Screen, x, y int) {
	screen.pixels[x][y] = true
}

// Check if a pixel on a Screen
// given it's x and y coordinates is set.
func IsSet(screen *Screen, x, y int) bool {
	return screen.pixels[x][y]
}

// Draw a sprite into a Screen at starting position with
// coordinates x,y, and the sprite size.
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

// ToString returns the string representation of a Screen.
func ToString(screen *Screen) string {
	var str string

	for x := range screen.pixels {
		str += "\n"
		for y := range screen.pixels {
			if screen.pixels[x][y] {
				str += "█ "
			} else {
				// screen.pixels[x][y]
				str += "_ "
			}
		}
	}

	return str
}
