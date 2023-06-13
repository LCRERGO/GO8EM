// Package register provides functions for registers manipulation.
package register

import (
	"fmt"

	"github.com/LCRERGO/GO8EM/pkg/constants"
)

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

// ToString returns the string representation of a RegisterFile.
func ToString(registers *RegisterFile) string {
	var vRepr string

	for _, v := range registers.V {
		vRepr += fmt.Sprintf("0x%04X, ", v)
	}

	return fmt.Sprintf(`
	{
		V: %s
		I: 0x%04X
		ST: 0x%02X, DT: 0x%02X

		SP: 0x%04X
		PC: 0x%04X
	}
`,
		vRepr, registers.I,
		registers.ST, registers.DT,
		registers.SP, registers.PC)
}
