package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	r := 0
	for i := 1; i < 9; i++ {
		r += int(pc[byte(x>>(i*8))])
	}
	return r
}
