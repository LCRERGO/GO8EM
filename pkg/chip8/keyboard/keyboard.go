// Package memory provides functions for keyboard map manipulation.
package keyboard

import "log"

type Keyboard struct {
	data map[int]bool
}

// Create a new keyboard.
func New() *Keyboard {
	return &Keyboard{
		data: map[int]bool{
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

// Deep Copy a Keyboard.
func Copy(keyboard *Keyboard) *Keyboard {
	data := make(map[int]bool)
	for k, v := range keyboard.data {
		data[k] = v
	}

	return &Keyboard{
		data: data,
	}
}

// Destroy a Keyboard.
func Destroy(keyboard *Keyboard) {
	keyboard.data = nil
	keyboard = nil
}

// Get Key current state from a Keyboard.
func KeyStatus(keyboard *Keyboard, key int) bool {
	if !isValidKey(keyboard, key) {
		log.Fatal("keyboard: invalid key")
	}
	return keyboard.data[key]
}

// Check if a key on a Keyboard is up.
func IsUp(keyboard *Keyboard, key int) bool {
	return KeyStatus(keyboard, key)
}

// Check if a key on a Keyboard is down.
func IsDown(keyboard *Keyboard, key int) bool {
	return !KeyStatus(keyboard, key)
}

// Set Key down in a Keyboard.
func SetKeyDown(keyboard *Keyboard, key int) {
	if !isValidKey(keyboard, key) {
		log.Fatal("keyboard: invalid key")
	}
	keyboard.data[key] = true
}

// Set Key up in a Keyboard.
func SetKeyUp(keyboard *Keyboard, key int) {
	if !isValidKey(keyboard, key) {
		log.Fatal("keyboard: invalid key")
	}
	keyboard.data[key] = false
}

func isValidKey(keyboard *Keyboard, key int) bool {
	for k, _ := range keyboard.data {
		if k == key {
			return true
		}
	}

	return false
}
