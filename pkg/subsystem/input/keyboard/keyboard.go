package keyboard

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var subsystemMap map[sdl.Scancode]int = map[sdl.Scancode]int{
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

func GetKeyRepr(keySymbol sdl.Keycode) (int, error) {
	for k, v := range subsystemMap {
		if sdl.GetKeyFromScancode(k) == keySymbol {
			return v, nil
		}
	}

	return 0xFF, fmt.Errorf("keyboard: invalid key")
}
