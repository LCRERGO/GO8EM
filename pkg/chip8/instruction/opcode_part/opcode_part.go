// Package opcode_part manipulate the opcode divide opcode in parts
// for the Instruction.
package opcode_part

// HighByte, LowByte,
// most significant bit (MSB), least significant bit (LSB)
// and Tail (the opcode excluding it's MSB)
// are needed to identify an instruction.
type OpcodeParts struct {
	HighByte, LowByte uint8
	MSB, LSB          uint8
	Tail              uint16
}

// Parse the OpcodeParts given an opcode.
func ParseOpcodeParts(opcode uint16) OpcodeParts {
	return OpcodeParts{
		HighByte: uint8((opcode >> 8) & 0x00FF),
		LowByte:  uint8(opcode & 0x00FF),
		MSB:      uint8((opcode >> 12) & 0x000F),
		LSB:      uint8(opcode & 0x000F),
		Tail:     opcode & 0x0FFF,
	}
}
