// Package to aggregate default constants.
package constants

const (
	WindowTitle         = "GO8EM" // name title of the executable
	MemorySize          = 0x1000  // upper memory limit
	ProgramStartAddress = 0x200   // Chip-8 default starting address
	ScreenWidth         = 64      // screen pixel width
	ScreenHeight        = 32      // screen pixel height
	ScreenScaleFactor   = 10      // default pixel unit scaling factor

	SpriteSetStartAddress = 0x00  // first address of sprite set
	DefaultSpriteHeight   = 5     // default height of sprite to serve as index
	ROMMaxSize            = 0xE00 // maximum ROM size to check limits

	ClockFrequency = 700 // default emulator clock frequency
)
