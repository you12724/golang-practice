package popcount

import "sync"

var loadPCOnce sync.Once
var pc [256]byte

func loadPC() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PC() [256]byte {
	loadPCOnce.Do(loadPC)
	return pc
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
