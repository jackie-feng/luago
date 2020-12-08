package vm

import "luago/api"

const LfieldsPerFlush = 50

// NEWTABLE
// R(A) := {} (size = B, C)
func newTable(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1
	vm.CreateTable(Fb2int(b), Fb2int(c))
	vm.Replace(a)
}

// GETTABLE
// R(A) = R(B)[RK(C)]
func getTable(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1
	b += 1
	vm.GetRK(c)
	vm.GetTable(b)
	vm.Replace(a)
}

// SETTABLE
// R(A)[RK(B)] := RK(C)
func setTable(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1
	vm.GetRK(b)
	vm.GetRK(c)
	vm.SetTable(a)
}

// SETLIST
// R(A)[(C-1)*FPF + i] := R(A+i), 1 <= i <= B
// 用于按索引批量设置数组元素
// 其中数组位于寄存器中，索引由操作数 A 指定
// 需要写入的一系列值也在寄存器中，紧挨着数组，数量由操作数 B 指定
// 数组起始索引由操作数 C 指定
func setList(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1
	if c > 0 {
		c = c - 1
	} else {
		c = Instruction(vm.Fetch()).Ax()
	}
	idx := int64(c * LfieldsPerFlush)
	for j := 1; j <= b; j++ {
		idx++
		vm.PushValue(a + j)
		vm.SetI(a, idx)
	}

	bIsZero := b == 0
	if bIsZero {
		b = int(vm.ToInteger(-1)) - a - 1
		vm.Pop(1)
	}
	if bIsZero {
		for j := vm.RegisterCount() + 1; j <= vm.GetTop(); j++ {
			idx++
			vm.PushValue(j)
			vm.SetI(a, idx)
		}
		vm.SetTop(vm.RegisterCount())
	}
}
