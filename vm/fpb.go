package vm

/*
如果某个字节用二进制表示成 eeeeexxx, 那么当 eeeee == 0 时, 改字节表示的整数为 xxx,
否则表示的整数是(1xxx)*2^(eeeee-1)
*/

func Int2fb(x int) int {
	// TODO:
	e := 0
	if x < 8 {
		return x
	}
	//x >= 0x80
	for x >= (8 << 4) {
		// x = ceil(x / 16)
		x = (x + 0xf) >> 4
		e += 4
	}
	for x >= (8 << 1) {
		x = (x + 1) >> 1
		e++
	}
	return ((e + 1) << 3) | (x - 8)
}

func Fb2int(x int) int {
	if x < 8 {
		return x
	} else {
		return ((x & 7) + 8) << uint((x>>3)-1)
	}
}
