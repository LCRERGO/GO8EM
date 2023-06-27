package utils

import (
	"reflect"
	"runtime"
	"unsafe"

	"golang.org/x/exp/constraints"
)

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func byte2Short(dataSlice []byte) uint16 {
	return uint16(dataSlice[1])<<8 + uint16(dataSlice[0])
}

func Byte2BCD(value byte) (units, tens, hundreds uint8) {
	hundreds = value / 100
	tens = (value / 10) - (hundreds * 10)
	units = value - (hundreds*100 + tens*10)

	return
}

func Num2Bin[T constraints.Integer](n T) string {
	var str string
	size := unsafe.Sizeof(n) * 8

	for i := 0; i < int(size); i++ {
		if (n>>i)&0x1 == 0x1 {
			str += "1"
		} else {
			str += "0"
		}
	}

	return str
}
