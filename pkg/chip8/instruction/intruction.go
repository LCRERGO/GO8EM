// Package of the instruction mapper.
package instruction

import (
	"github.com/LCRERGO/GO8EM/pkg/chip8"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/argument"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr00"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr01"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr02"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr03"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr04"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr05"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr06"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr07"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr08"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr09"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr10"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr11"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr12"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr13"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr14"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr15"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr16"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr17"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr18"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr19"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr20"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr21"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr22"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr23"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr24"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr25"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr26"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr27"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr28"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr29"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr30"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr31"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr32"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr33"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr34"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/instr35"
	"github.com/LCRERGO/GO8EM/pkg/chip8/instruction/opcode_part"
)

type ExecInstruction func(*chip8.Chip8, *argument.OpcodeArguments)

// An Instruction consists of a Tag which is the symbol of an
// Instruction, it's Repr which is how it is represented in
// assembly form, and it's ExecFunc which is what operation
// should be performed when the instruction is encountered.
type Instruction struct {
	Tag      string
	Repr     string
	ExecFunc ExecInstruction
}

// A mapper of Instructions.
var instructionMap []Instruction = []Instruction{
	{"ERROR", "error", instr00.Exec},
	{"SYS", "0nnn - SYS addr", instr01.Exec},
	{"CLS", "00E0 - CLS", instr02.Exec},
	{"RET", "00EE - RET", instr03.Exec},
	{"JPADDR", "1nnn - JP addr", instr04.Exec},
	{"CALL", "2nnn - CALL addr", instr05.Exec},
	{"SEVXBYTE", "3xkk - SE Vx, byte", instr06.Exec},
	{"SNEVXBYTE", "4xkk - SNE Vx, byte", instr07.Exec},
	{"SEVXVY", "5xy0 - SE Vx, Vy", instr08.Exec},
	{"LDVXBYTE", "6xkk - LD Vx, byte", instr09.Exec},
	{"ADDVXBYTE", "7xkk - ADD Vx, byte", instr10.Exec},
	{"LDVXVY", "8xy0 - LD Vx, Vy", instr11.Exec},
	{"ORVXVY", "8xy1 - OR Vx, Vy", instr12.Exec},
	{"ANDVXVY", "8xy2 - AND Vx, Vy", instr13.Exec},
	{"XORVXVY", "8xy3 - XOR Vx, Vy", instr14.Exec},
	{"ADDVXVY", "8xy4 - ADD Vx, Vy", instr15.Exec},
	{"SUBVXVY", "8xy5 - SUB Vx, Vy", instr16.Exec},
	{"SHRVXVY", "8xy6 - SHR Vx {, Vy}", instr17.Exec},
	{"SUBNVXVY", "8xy7 - SUBN Vx, Vy", instr18.Exec},
	{"SHLVXVY", "8xyE - SHL Vx {, Vy}", instr19.Exec},
	{"SNEVXVY", "9xy0 - SNE Vx, Vy", instr20.Exec},
	{"LDIADDR", "Annn - LD I, addr", instr21.Exec},
	{"JPV0ADDR", "Bnnn - JP V0, addr", instr22.Exec},
	{"RNDVXBYTE", "Cxkk - RND Vx, byte", instr23.Exec},
	{"DRWVXVYNIBBLE", "Dxyn - DRW Vx, Vy, nibble", instr24.Exec},
	{"SKPVX", "Ex9E - SKP Vx", instr25.Exec},
	{"SKNPVX", "ExA1 - SKNP Vx", instr26.Exec},
	{"LDVXDT", "Fx07 - LD Vx, DT", instr27.Exec},
	{"LDVXK", "Fx0A - LD Vx, K", instr28.Exec},
	{"LDDTVX", "Fx15 - LD DT, Vx", instr29.Exec},
	{"LDSTVX", "Fx18 - LD ST, Vx", instr30.Exec},
	{"ADDIVX", "Fx1E - ADD I, Vx", instr31.Exec},
	{"LDFVX", "Fx29 - LD F, Vx", instr32.Exec},
	{"LDBVX", "Fx33 - LD B, Vx", instr33.Exec},
	{"LDIVX", "Fx55 - LD [I], Vx", instr34.Exec},
	{"LDVXI", "Fx65 - LD Vx, [I]", instr35.Exec},
}

// Parse an Instruction given an opcode.
func ParseInstruction(opcode uint16) Instruction {
	var instruction Instruction
	parts := opcode_part.ParseOpcodeParts(opcode)

	switch parts.MSB {
	case 0x0:
		switch parts.Tail {
		case 0x0E0:
			instruction = instructionMap[2]
		case 0x0EE:
			instruction = instructionMap[3]
		default:
			// error
			instruction = instructionMap[1]
		}
	case 0x1:
		instruction = instructionMap[4]
	case 0x2:
		instruction = instructionMap[5]
	case 0x3:
		instruction = instructionMap[6]
	case 0x4:
		instruction = instructionMap[7]
	case 0x5:
		instruction = instructionMap[8]
	case 0x6:
		instruction = instructionMap[9]
	case 0x7:
		instruction = instructionMap[10]
	case 0x8:
		switch parts.LSB {
		case 0x0:
			instruction = instructionMap[11]
		case 0x1:
			instruction = instructionMap[12]
		case 0x2:
			instruction = instructionMap[13]
		case 0x3:
			instruction = instructionMap[14]
		case 0x4:
			instruction = instructionMap[15]
		case 0x5:
			instruction = instructionMap[16]
		case 0x6:
			instruction = instructionMap[17]
		case 0x7:
			instruction = instructionMap[18]
		case 0xE:
			instruction = instructionMap[19]
		default:
			// error
			instruction = instructionMap[0]
		}
	case 0x9:
		switch parts.LSB {
		case 0x0:
			instruction = instructionMap[20]
		default:
			// error
			instruction = instructionMap[0]
		}
	case 0xA:
		instruction = instructionMap[21]
	case 0xB:
		instruction = instructionMap[22]
	case 0xC:
		instruction = instructionMap[23]
	case 0xD:
		instruction = instructionMap[24]
	case 0xE:
		switch parts.LowByte {
		case 0x9E:
			instruction = instructionMap[25]
		case 0xA1:
			instruction = instructionMap[26]
		default:
			// error
			instruction = instructionMap[0]
		}
	case 0xF:
		switch parts.LowByte {
		case 0x07:
			instruction = instructionMap[27]
		case 0x0A:
			instruction = instructionMap[28]
		case 0x15:
			instruction = instructionMap[29]
		case 0x18:
			instruction = instructionMap[30]
		case 0x1E:
			instruction = instructionMap[31]
		case 0x29:
			instruction = instructionMap[32]
		case 0x33:
			instruction = instructionMap[33]
		case 0x55:
			instruction = instructionMap[34]
		case 0x65:
			instruction = instructionMap[35]
		default:
			// error
			instruction = instructionMap[0]
		}
	default:
		// error
		instruction = instructionMap[0]
	}

	return instruction
}
