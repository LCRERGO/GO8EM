package lcg

var seed int

func LCG() int {
	a := 25214903917
	m := 281474976710656
	c := 11
	seed = (a*seed + c) % m

	return seed
}

func SeedLCG(value int) {
	seed = value
}
