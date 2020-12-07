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
}
