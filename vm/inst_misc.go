package vm

import "luago/api"

// R(A) := R(B)
func move(i Instruction, vm api.LuaVM) {
	// 寄存器索引
	a, b, _ := i.ABC()
	// 寄存器索引 + 1 == 栈索引
	a += 1
	b += 1
	vm.Copy(b, a)
}

func jmp(i Instruction, vm api.LuaVM) {
	a, sBx := i.AsBx()
	vm.AddPC(sBx)
	if a != 0 {
		panic("todo!")
	}
}
