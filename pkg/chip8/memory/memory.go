package memory

import "github.com/LCRERGO/GO8EM/pkg/constants"

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
// 4096 (0x1000) bytes in total
type Memory struct {
	data [constants.MemorySize]byte
}

func New() *Memory {
	return &Memory{}
}

func Destroy(memory *Memory) {
	memory = nil
}

func Set(mem *Memory, index int, val uint8) {
	mem.data[index] = val
}

func Get(mem Memory, index int) uint8 {
	return mem.data[index]
}

func Get16(mem Memory, index int) uint16 {
	var fstByte uint16 = uint16(Get(mem, index))
	var sndByte uint16 = uint16(Get(mem, index+1))

	return fstByte<<8 | sndByte
}
