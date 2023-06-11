// Package register provides functions for registers manipulation.
package register

import "github.com/LCRERGO/GO8EM/pkg/constants"

// RegisterFile entity provides access to all the
// registers available in the emulator.
type RegisterFile struct {
	V      []uint16 // General Porpuse registers
	I      uint16   // Memory register
	ST, DT uint8    // Sound and Delay registers
	SP     uint16   // Stack Pointer
	PC     uint16   // Program Counter
}

// Create a new RegisterFile.
func New() *RegisterFile {
	return &RegisterFile{
		V:  make([]uint16, 0x10),
		PC: constants.ProgramStartAddress,
	}
}

// Deep Copy a RegisterFile.
func Copy(registers *RegisterFile) *RegisterFile {
	v := make([]uint16, 0x10)
	copy(v, registers.V)

	return &RegisterFile{
		V:  v,
		I:  registers.I,
		ST: registers.ST,
		DT: registers.DT,
		SP: registers.SP,
		PC: registers.PC,
	}
}

// Destroy a RegisterFile.
func Destroy(registers *RegisterFile) {
	registers.V = nil
	registers = nil
}

// Increments PC register from a RegisterFile
func NextInstruction(registers *RegisterFile) {
	registers.PC += 2
}

// Decrement the DT register from a RegisterFile.
func DecDT(registers *RegisterFile) {
	registers.DT--
}

// Reset the ST register from a RegisterFile.
func ResetST(registers *RegisterFile) {
	registers.ST = 0
}
