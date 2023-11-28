// Package memory provides functions for memory manipulation.
package memory

import (
	"fmt"

	"github.com/LCRERGO/GO8EM/pkg/config"
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
	data := make([]byte, constants.MemorySize)
	copy(data[0x000:0x1FF], DefaultSprites)
	return &Memory{
		data: data,
	}
}

// Load a rom into Memory at adress 0x200.
func LoadROM(memory *Memory, rom []byte, size uint) {
	baseAddress := uint(constants.ProgramStartAddress)

	copy(memory.data[baseAddress:baseAddress+size], rom)
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
func Set(memory *Memory, index int, val uint8) {
	if !isValidIndex(index) {
		config.GetLogger(config.GetInstance()).
			Fatal("memory_set: invalid index")
	}
	memory.data[index] = val
}

// Get a single byte from a Memory given an index.
func Get(memory *Memory, index int) uint8 {
	if !isValidIndex(index) {
		config.GetLogger(config.GetInstance()).
			Fatal("memory_get: invalid index")
	}
	return memory.data[index]
}

// Get two bytes from a Memory given a starting index.
func Get16(memory *Memory, index int) uint16 {
	if !isValidIndex(index) {
		config.GetLogger(config.GetInstance()).
			Fatal("memory_get16: invalid index")
	}
	var fstByte uint16 = uint16(Get(memory, index))
	var sndByte uint16 = uint16(Get(memory, index+1))

	return fstByte<<8 | sndByte
}

// Fetch a slice given a starting index and a size
func FetchSprite(memory *Memory, index int, size int) []byte {
	data := make([]byte, size)
	copy(data, memory.data[index:index+size+1])

	return data
}

// ToString returns the string representation of a Memory.
func ToString(memory *Memory) string {
	var str string
	for i, v := range memory.data {
		if i%0x10 == 0 {
			str += "\n"
		}
		str += fmt.Sprintf("%02X ", v)
	}

	return str
}

func isValidIndex(index int) bool {
	return index >= 0x00 && index < constants.MemorySize
}
