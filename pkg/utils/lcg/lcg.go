// Package of linear congruential generator util.
package lcg

type LCG struct {
	seed int
}

func New(value int) *LCG {
	return &LCG{
		seed: value,
	}
}

func Destroy(lcg *LCG) {
	lcg.seed = 0
	lcg = nil
}

func RandInt(lcg *LCG) int {
	a := 25214903917
	m := 281474976710656
	c := 11
	lcg.seed = (a*lcg.seed + c) % m

	return lcg.seed
}
