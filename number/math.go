package number

import (
	"math"
)

func IFloorDiv(a, b int64) int64 {
	if (a > 0 && b > 0) || (a < 0 && b < 0) || (a % b == 0) {
		return a / b
	} else {
		// Go 中向 0 取整, -10 / 3 = -3.33333... => -3
		// Lua 中向下取整, -10 / 3 = -3.33333... => -4
		return a / b - 1
	}
}

func FFloorDiv(a, b float64) float64 {
	return math.Floor(a / b)
}

func IMod(a, b int64) int64 {
	return a - IFloorDiv(a, b) * b
}

func FMod(a, b float64) float64 {
	return a - math.Floor(a / b) * b
}

func ShiftLeft(a, n int64) int64 {
	if n >= 0 {
		return a << uint64(n)
	} else {
		return ShiftRight(a, -n)
	}
}

func ShiftRight(a, n int64) int64 {
	if n >= 0 {
		// golang 右移位为有符号数移位, 自动补1, 转为无符号数自动补0
		return int64(uint64(a) >> uint64(n))
	} else {
		return ShiftLeft(a, -n)
	}
}

func FloatToInteger(f float64) (int64, bool) {
	i := int64(f)
	return i, float64(i) == f
}
