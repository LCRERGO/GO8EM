package keyboard

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var subsystemMap map[sdl.Keycode]int = map[sdl.Keycode]int{
	sdl.SCANCODE_0: 0x00,
	sdl.SCANCODE_1: 0x01,
	sdl.SCANCODE_2: 0x02,
	sdl.SCANCODE_3: 0x03,
	sdl.SCANCODE_4: 0x04,
	sdl.SCANCODE_5: 0x05,
	sdl.SCANCODE_6: 0x06,
	sdl.SCANCODE_7: 0x07,
	sdl.SCANCODE_8: 0x08,
	sdl.SCANCODE_9: 0x09,
	sdl.SCANCODE_A: 0x0A,
	sdl.SCANCODE_B: 0x0B,
	sdl.SCANCODE_C: 0x0C,
	sdl.SCANCODE_D: 0x0D,
	sdl.SCANCODE_E: 0x0E,
	sdl.SCANCODE_F: 0x0F,
}

type KeyboardSubsystem struct {
	keyState map[int]bool
}

func New() *KeyboardSubsystem {
	return &KeyboardSubsystem{
		keyState: map[int]bool{
			0x00: false,
			0x01: false,
			0x02: false,
			0x03: false,
			0x04: false,
			0x05: false,
			0x06: false,
			0x07: false,
			0x08: false,
			0x09: false,
			0x0A: false,
			0x0B: false,
			0x0C: false,
			0x0D: false,
			0x0E: false,
			0x0F: false,
		},
	}
}

func Destroy(keyboard *KeyboardSubsystem) {
	keyboard = nil
}

func GetKeyRepr(keyboard *KeyboardSubsystem, keySymbol sdl.Keycode) (int, error) {
	for k, v := range subsystemMap {
		if k == keySymbol {
			return v, nil
		}
	}

	return 0xFF, fmt.Errorf("keyboard: invalid key")
}

func KeyStatus(keyboard *KeyboardSubsystem, key int) bool {
	return keyboard.keyState[key]
}

func KeyDown(keyboard *KeyboardSubsystem, key int) {
	keyboard.keyState[key] = true
}

func KeyUp(keyboard *KeyboardSubsystem, key int) {
	keyboard.keyState[key] = false
}
