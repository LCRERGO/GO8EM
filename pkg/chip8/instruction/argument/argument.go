// Package of instruction argument types.
package argument

// Agregate the arguments of an opcode, which may be in the form:
// nnn or addr - A 12-bit value, the lowest 12 bits of the instruction
// n or nibble - A 4-bit value, the lowest 4 bits of the instruction
// x - A 4-bit value, the lower 4 bits of the high byte of the instruction
// y - A 4-bit value, the upper 4 bits of the low byte of the instruction
// kk or byte - An 8-bit value, the lowest 8 bits of the instruction
type OpcodeArguments struct {
	NNN  uint16
	N    uint8
	X, Y uint8
	KK   uint8
}

// Parse OpcodeArguments given an opcode.
func ParseArguments(opcode uint16) *OpcodeArguments {
	return &OpcodeArguments{
		NNN: opcode & 0x0FFF,
		N:   uint8(opcode & 0x000F),
		X:   uint8((opcode >> 8) & 0x000F),
		Y:   uint8((opcode >> 4) & 0x000F),
		KK:  uint8(opcode & 0x00FF),
	}
}
