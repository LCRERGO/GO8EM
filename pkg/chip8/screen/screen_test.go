package screen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClear(t *testing.T) {
	tests := []struct {
		name   string
		screen *Screen

		wantScreenState *Screen
	}{
		{
			name:            "screen clear empty screen",
			screen:          New(),
			wantScreenState: New(),
		},
		{
			name: "screen clear with pixels set at borders",
			screen: func() *Screen {
				state := New()
				state.pixels[0][0] = true
				state.pixels[0][31] = true
				state.pixels[63][0] = true
				state.pixels[63][31] = true

				return state
			}(),
			wantScreenState: New(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Clear(tt.screen)

			assert.Equal(t, tt.wantScreenState, tt.screen)
		})
	}
}

func TestSet(t *testing.T) {
	tests := []struct {
		name   string
		screen *Screen
		x, y   int

		wantScreenState *Screen
	}{
		{
			name:   "set first pixel",
			screen: New(),
			x:      0,
			y:      0,

			wantScreenState: func() *Screen {
				state := New()
				state.pixels[0][0] = true

				return state
			}(),
		},
		{
			name:   "set last column pixel",
			screen: New(),
			x:      0,
			y:      31,

			wantScreenState: func() *Screen {
				state := New()
				state.pixels[0][31] = true

				return state
			}(),
		},
		{
			name:   "set last row pixel",
			screen: New(),
			x:      63,
			y:      0,

			wantScreenState: func() *Screen {
				state := New()
				state.pixels[63][0] = true

				return state
			}(),
		},
		{
			name:   "set last pixel",
			screen: New(),
			x:      63,
			y:      31,

			wantScreenState: func() *Screen {
				state := New()
				state.pixels[63][31] = true

				return state
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Set(tt.screen, tt.x, tt.y)

			assert.Equal(t, tt.wantScreenState, tt.screen)
		})
	}
}

func TestIsSet(t *testing.T) {
	tests := []struct {
		name   string
		screen *Screen
		x, y   int

		wantScreenState *Screen
		want            bool
	}{
		{
			name:   "check middle pixel on an empty screen",
			screen: New(),
			x:      31,
			y:      15,

			wantScreenState: New(),
			want:            false,
		},
		{
			name: "check middle pixel with pixel set",
			screen: func() *Screen {
				state := New()
				state.pixels[31][15] = true

				return state
			}(),
			x: 31,
			y: 15,

			wantScreenState: func() *Screen {
				state := New()
				state.pixels[31][15] = true

				return state
			}(), want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := IsSet(tt.screen, tt.x, tt.y)

			assert.Equal(t, tt.wantScreenState, tt.screen)
			assert.Equal(t, tt.want, value)
		})
	}
}

func TestDrawSprite(t *testing.T) {
	tests := []struct {
		name   string
		screen *Screen
		x, y   int
		sprite []byte
		size   int

		wantScreenState *Screen
		wantSpriteState []byte
		want            bool
	}{
		{
			name:   "Drawing one pixel at first position",
			screen: New(),
			x:      0,
			y:      0,
			sprite: []byte{0x80},
			size:   1,

			wantScreenState: func() *Screen {
				state := New()
				state.pixels[0][0] = true

				return state
			}(),
			wantSpriteState: []byte{0x80},
			want:            false,
		},
		{
			name: "Drawing one pixel on a previously written pixel",
			screen: func() *Screen {
				state := New()
				state.pixels[0][0] = true

				return state
			}(),
			x:      0,
			y:      0,
			sprite: []byte{0x80},
			size:   1,

			wantScreenState: New(),
			wantSpriteState: []byte{0x80},
			want:            true,
		},
		{
			name:   "Drawing two pixels at the horizontal limit of the screen",
			screen: New(),
			x:      0,
			y:      31,
			sprite: []byte{0xC0},
			size:   1,

			wantScreenState: func() *Screen {
				state := New()
				state.pixels[0][31] = true
				state.pixels[1][31] = true

				return state
			}(),
			wantSpriteState: []byte{0xC0},
			want:            false,
		},
		{
			name:   "Drawing sprite 0 at initial position",
			screen: New(),
			x:      0,
			y:      0,
			sprite: []byte{0xF0, 0x90, 0x90, 0x90, 0xF0},
			size:   5,

			wantScreenState: func() *Screen {
				state := New()
				copy(state.pixels[0][0:5], []bool{true, true, true, true, true})
				copy(state.pixels[1][0:5], []bool{true, false, false, false, true})
				copy(state.pixels[2][0:5], []bool{true, false, false, false, true})
				copy(state.pixels[3][0:5], []bool{true, true, true, true, true})

				return state
			}(),
			wantSpriteState: []byte{0xF0, 0x90, 0x90, 0x90, 0xF0},
			want:            false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := DrawSprite(tt.screen, tt.x, tt.y, tt.sprite, tt.size)

			assert.Equal(t, tt.wantScreenState, tt.screen)
			assert.Equal(t, tt.wantSpriteState, tt.sprite)
			assert.Equal(t, tt.want, value)
		})
	}
}
