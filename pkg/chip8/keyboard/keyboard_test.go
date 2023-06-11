package keyboard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyStatus(t *testing.T) {
	tests := []struct {
		name     string
		keyboard *Keyboard
		key      int

		wantKeyboard *Keyboard
		want         bool
	}{
		{
			name:     "get first key up",
			keyboard: New(),
			key:      0x00,

			wantKeyboard: New(),
			want:         false,
		},
		{
			name: "get first key down",
			keyboard: func() *Keyboard {
				state := New()
				state.data[0x00] = true

				return state
			}(),
			key: 0x00,

			wantKeyboard: func() *Keyboard {
				state := New()
				state.data[0x00] = true

				return state
			}(),

			want: true,
		},
		{
			name:     "get last key up",
			keyboard: New(),
			key:      0x0F,

			wantKeyboard: New(),
			want:         false,
		},
		{
			name: "get last key down",
			keyboard: func() *Keyboard {
				state := New()
				state.data[0x0F] = true

				return state
			}(),
			key: 0x0F,

			wantKeyboard: func() *Keyboard {
				state := New()
				state.data[0x0F] = true

				return state
			}(),

			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := KeyStatus(tt.keyboard, tt.key)
			assert.Equal(t, tt.wantKeyboard, tt.keyboard)
			assert.Equal(t, tt.want, value)
		})
	}
}

func TestSetKeyDown(t *testing.T) {
	tests := []struct {
		name     string
		keyboard *Keyboard
		key      int

		wantKeyboard *Keyboard
	}{
		{
			name: "set first key down (already set)",
			keyboard: func() *Keyboard {
				state := New()
				state.data[0x00] = true

				return state
			}(),
			key: 0x00,

			wantKeyboard: func() *Keyboard {
				state := New()
				state.data[0x00] = true

				return state
			}(),
		},
		{
			name:     "set first key down (not set)",
			keyboard: New(),
			key:      0x00,

			wantKeyboard: func() *Keyboard {
				state := New()
				state.data[0x00] = true

				return state
			}(),
		},
		{
			name: "set last key down (already set)",
			keyboard: func() *Keyboard {
				state := New()
				state.data[0x0F] = true

				return state
			}(),
			key: 0x0F,

			wantKeyboard: func() *Keyboard {
				state := New()
				state.data[0x0F] = true

				return state
			}(),
		},
		{
			name:     "set last key down (not set)",
			keyboard: New(),
			key:      0x0F,

			wantKeyboard: func() *Keyboard {
				state := New()
				state.data[0x0F] = true

				return state
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetKeyDown(tt.keyboard, tt.key)
			assert.Equal(t, tt.wantKeyboard, tt.keyboard)
		})
	}

}

func TestSetKeyUp(t *testing.T) {
	tests := []struct {
		name     string
		keyboard *Keyboard
		key      int

		wantKeyboard *Keyboard
	}{
		{
			name: "set first key up (already set)",
			keyboard: func() *Keyboard {
				state := New()
				state.data[0x00] = true

				return state
			}(),
			key: 0x00,

			wantKeyboard: New(),
		},
		{
			name:     "set first key up (not set)",
			keyboard: New(),
			key:      0x00,

			wantKeyboard: New(),
		},
		{
			name: "set last key up (already set)",
			keyboard: func() *Keyboard {
				state := New()
				state.data[0x0F] = true

				return state
			}(),
			key: 0x0F,

			wantKeyboard: New(),
		},
		{
			name:     "set last key down (not set)",
			keyboard: New(),
			key:      0x0F,

			wantKeyboard: New(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetKeyUp(tt.keyboard, tt.key)
			assert.Equal(t, tt.wantKeyboard, tt.keyboard)
		})
	}

}
