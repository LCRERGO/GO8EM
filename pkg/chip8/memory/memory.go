// Package memory provides functions for memory manipulation.
package memory

import (
	"log"

	"github.com/LCRERGO/GO8EM/pkg/constants"
)

// +---------------+= 0xFFF (4095) End of Chip-8 RAM
// |               |
// |               |
// |               |
// |               |
// |               |
// | 0x200 to 0xFFF|
// |     Chip-8    |
// | Program / Data|
// |     Space     |
// |               |
// |               |
// |               |
// +- - - - - - - -+= 0x600 (1536) Start of ETI 660 Chip-8 programs
// |               |
// |               |
// |               |
// +---------------+= 0x200 (512) Start of most Chip-8 programs
// | 0x000 to 0x1FF|
// | Reserved for  |
// |  interpreter  |
// +---------------+= 0x000 (0) Start of Chip-8 RAM
//
// Memory size on CHIP-8 goes from
// 0x000 to 0xFFF, thus having
// 4096 (0x1000) bytes in total.
type Memory struct {
	data []byte
}

// Create a new Memory of size constants.MemorySize.
func New() *Memory {
	return &Memory{
		data: make([]byte, constants.MemorySize),
	}
}

// Deep Copy a Memory.
func Copy(memory *Memory) *Memory {
	data := make([]byte, constants.MemorySize)
	copy(data, memory.data)

	return &Memory{
		data: data,
	}
}

// Destroy a Memory.
func Destroy(memory *Memory) {
	memory.data = nil
	memory = nil
}

// Set a single byte from a Memory given an index
// for the position and a value to be set.
func Set(mem *Memory, index int, val uint8) {
	if !isValidIndex(index) {
		log.Fatal("memory_set: invalid index")
	}
	mem.data[index] = val
}

// Get a single byte from a Memory given an index.
func Get(mem *Memory, index int) uint8 {
	if !isValidIndex(index) {
		log.Fatal("memory_get: invalid index")
	}
	return mem.data[index]
}

// Get two bytes from a Memory given a starting index
func Get16(mem *Memory, index int) uint16 {
	if !isValidIndex(index) {
		log.Fatal("memory_get16: invalid index")
	}
	var fstByte uint16 = uint16(Get(mem, index))
	var sndByte uint16 = uint16(Get(mem, index+1))

	return fstByte<<8 | sndByte
}

func isValidIndex(index int) bool {
	return index >= 0x00 && index < constants.MemorySize
}
