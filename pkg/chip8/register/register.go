package register

type RegisterFile struct {
	V      [16]uint16 // General Porpuse registers
	I      uint16     // Memory register
	ST, DT uint8      // Sound and Delay registers
	SP     uint8      // Stack Pointer
	PC     uint16     // Program Counter
}

func New() *RegisterFile {
	return &RegisterFile{}
}

func Destroy(registerFile *RegisterFile) {
	registerFile = nil
}
