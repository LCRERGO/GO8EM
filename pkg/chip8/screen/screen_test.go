package screen

import (
	"testing"

	"github.com/LCRERGO/GO8EM/pkg/constants"
	"github.com/stretchr/testify/assert"
)

func TestClear(t *testing.T) {
	tests := []struct {
		name   string
		screen Screen

		wantScreenState Screen
	}{
		{
			name:            "screen clear empty screen",
			screen:          Screen{},
			wantScreenState: Screen{},
		},
		{
			name: "screen clear with pixels set at borders",
			screen: Screen{
				pixels: [constants.ScreenWidth][constants.ScreenHeight]bool{
					0:  {0: true, 31: true},
					63: {0: true, 31: true},
				},
			},
			wantScreenState: Screen{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Clear(&tt.screen)

			assert.Equal(t, tt.wantScreenState, tt.screen)
		})
	}
}

func TestSet(t *testing.T) {
	tests := []struct {
		name   string
		screen Screen
		x, y   int

		wantScreenState Screen
	}{
		{
			name:   "set first pixel",
			screen: Screen{},
			x:      0,
			y:      0,

			wantScreenState: Screen{
				pixels: [constants.ScreenWidth][constants.ScreenHeight]bool{{true}},
			},
		},
		{
			name:   "set last column pixel",
			screen: Screen{},
			x:      0,
			y:      31,

			wantScreenState: Screen{
				pixels: [constants.ScreenWidth][constants.ScreenHeight]bool{{31: true}},
			},
		},
		{
			name:   "set last row pixel",
			screen: Screen{},
			x:      63,
			y:      0,

			wantScreenState: Screen{
				pixels: [constants.ScreenWidth][constants.ScreenHeight]bool{63: {true}},
			},
		},
		{
			name:   "set last pixel",
			screen: Screen{},
			x:      63,
			y:      31,

			wantScreenState: Screen{
				pixels: [constants.ScreenWidth][constants.ScreenHeight]bool{63: {31: true}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Set(&tt.screen, tt.x, tt.y)

			assert.Equal(t, tt.wantScreenState, tt.screen)
		})
	}
}

func TestIsSet(t *testing.T) {
	tests := []struct {
		name   string
		screen Screen
		x, y   int

		wantScreenState Screen
		want            bool
	}{
		{
			name:   "check middle pixel on an empty screen",
			screen: Screen{},
			x:      31,
			y:      15,

			wantScreenState: Screen{},
			want:            false,
		},
		{
			name: "check middle pixel with pixel set",
			screen: Screen{
				pixels: [constants.ScreenWidth][constants.ScreenHeight]bool{31: {15: true}},
			},
			x: 31,
			y: 15,

			wantScreenState: Screen{
				pixels: [constants.ScreenWidth][constants.ScreenHeight]bool{31: {15: true}},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := IsSet(&tt.screen, tt.x, tt.y)

			assert.Equal(t, tt.wantScreenState, tt.screen)
			assert.Equal(t, tt.want, value)
		})
	}
}

func TestDrawSprite(t *testing.T) {
	tests := []struct {
		name   string
		screen Screen
		x, y   int
		sprite []byte
		size   int

		wantScreenState Screen
		wantSpriteState []byte
		want            bool
	}{
		{
			name:   "Drawing one pixel at first position",
			screen: Screen{},
			x:      0,
			y:      0,
			sprite: []byte{0x80},
			size:   1,

			wantScreenState: Screen{
				pixels: [constants.ScreenWidth][constants.ScreenHeight]bool{
					{true},
				},
			},
			wantSpriteState: []byte{0x80},
			want:            false,
		},
		{
			name: "Drawing one pixel on a previously written pixel",
			screen: Screen{
				pixels: [constants.ScreenWidth][constants.ScreenHeight]bool{
					{true},
				},
			},
			x:      0,
			y:      0,
			sprite: []byte{0x80},
			size:   1,

			wantScreenState: Screen{
				pixels: [constants.ScreenWidth][constants.ScreenHeight]bool{
					{false},
				},
			},
			wantSpriteState: []byte{0x80},
			want:            true,
		},
		{
			name:   "Drawing two pixels at the horizontal limit of the screen",
			screen: Screen{},
			x:      0,
			y:      31,
			sprite: []byte{0xC0},
			size:   1,

			wantScreenState: Screen{
				pixels: [constants.ScreenWidth][constants.ScreenHeight]bool{
					0: {31: true},
					1: {31: true},
				},
			},
			wantSpriteState: []byte{0xC0},
			want:            false,
		},
		{
			name:   "Drawing sprite 0 at initial position",
			screen: Screen{},
			x:      0,
			y:      0,
			sprite: []byte{0xF0, 0x90, 0x90, 0x90, 0xF0},
			size:   5,

			wantScreenState: Screen{
				pixels: [constants.ScreenWidth][constants.ScreenHeight]bool{
					{
						true, true, true, true, true,
					},
					{
						true, false, false, false, true,
					},
					{
						true, false, false, false, true,
					},
					{
						true, true, true, true, true,
					},
				},
			},
			wantSpriteState: []byte{0xF0, 0x90, 0x90, 0x90, 0xF0},
			want:            false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := DrawSprite(&tt.screen, tt.x, tt.y, tt.sprite, tt.size)

			assert.Equal(t, tt.wantScreenState, tt.screen)
			assert.Equal(t, tt.wantSpriteState, tt.sprite)
			assert.Equal(t, tt.want, value)
		})
	}
}
