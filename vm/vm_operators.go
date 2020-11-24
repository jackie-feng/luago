package vm

import "luago/api"

// R(A) := RK(B) op RK(C)
// 二元运算指令
func _binaryArith(i Instruction, vm api.LuaVM, op api.ArithOp) {
	a, b, c := i.ABC()
	a += 1
	// 获取 b, c 索引指向到值, 推到栈顶
	vm.GetRK(b)
	vm.GetRK(c)
	// 弹出栈顶元素 计算结果推入栈顶
	vm.Arith(op)
	// 弹出栈顶元素 写入指定栈索引处
	vm.Replace(a)
}

// R(A) := op R(B)
// 一元运算指令
func _unaryArith(i Instruction, vm api.LuaVM, op api.ArithOp) {
	a, b, _ := i.ABC()
	a += 1
	b += 1
	vm.PushValue(b)
	vm.Arith(op)
	vm.Replace(a)
}

func add(i Instruction, vm api.LuaVM)  { _binaryArith(i, vm, api.LuaOpADD) }
func sub(i Instruction, vm api.LuaVM)  { _binaryArith(i, vm, api.LuaOpSUB) }
func mul(i Instruction, vm api.LuaVM)  { _binaryArith(i, vm, api.LuaOpMUL) }
func mod(i Instruction, vm api.LuaVM)  { _binaryArith(i, vm, api.LuaOpMOD) }
func pow(i Instruction, vm api.LuaVM)  { _binaryArith(i, vm, api.LuaOpPOW) }
func div(i Instruction, vm api.LuaVM)  { _binaryArith(i, vm, api.LuaOpDIV) }
func idiv(i Instruction, vm api.LuaVM) { _binaryArith(i, vm, api.LuaOpIDIV) }
func band(i Instruction, vm api.LuaVM) { _binaryArith(i, vm, api.LuaOpBAND) }
func bor(i Instruction, vm api.LuaVM)  { _binaryArith(i, vm, api.LuaOpBOR) }
func bxor(i Instruction, vm api.LuaVM) { _binaryArith(i, vm, api.LuaOpBXOR) }
func shl(i Instruction, vm api.LuaVM)  { _binaryArith(i, vm, api.LuaOpSHL) }
func shr(i Instruction, vm api.LuaVM)  { _binaryArith(i, vm, api.LuaOpSHR) }
func unm(i Instruction, vm api.LuaVM)  { _unaryArith(i, vm, api.LuaOpUNM) }
func bnot(i Instruction, vm api.LuaVM) { _unaryArith(i, vm, api.LuaOpBNOT) }

// R(A) := length of R(B)
func _len(i Instruction, vm api.LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	b += 1
	vm.Len(b)
	vm.Replace(a)
}

// R(A) := R(B).. ... ..R(C)
func concat(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1
	b += 1
	c += 1
	n := c - b + 1
	vm.CheckStack(n)
	for i := b; i <= c; i++ {
		vm.PushValue(i)
	}
	vm.Concat(n)
	vm.Replace(a)
}

// if((RK(B) op RK(C)) ~= A) then pc++
func _compare(i Instruction, vm api.LuaVM, op api.CompareOp) {
	a, b, c := i.ABC()
	vm.GetRK(b)
	vm.GetRK(c)
	if vm.Compare(-2, -1, op) == (a == 0) {
		vm.AddPC(1)
	}
	vm.Pop(2)
}

func eq(i Instruction, vm api.LuaVM) { _compare(i, vm, api.LuaOpEQ) } // ==
func lt(i Instruction, vm api.LuaVM) { _compare(i, vm, api.LuaOpLT) } // <
func le(i Instruction, vm api.LuaVM) { _compare(i, vm, api.LuaOpLE) } // <=

// R(A) := not R(B)
func not(i Instruction, vm api.LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	b += 1
	vm.PushBoolean(!vm.ToBoolean(b))
	vm.Replace(a)
}

// if (R(B) <=> C) the R(A) := R(B) else pc++
func testset(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1
	b += 1

	if vm.ToBoolean(b) == (c != 0) {
		vm.Copy(b, a)
	} else {
		vm.AddPC(1)
	}
}

// if not(R(A) <=> C) then pc ++
// [[false, 1], [true, 1], [false, 0], [true, 0]]
// [true, false, false, true]
func test(i Instruction, vm api.LuaVM) {
	a, _, c := i.ABC()
	a += 1
	if vm.ToBoolean(a) == (c == 0) {
		vm.AddPC(1)
	}
}

/*
for 循环

FORPREP:
	R(A) -= R(A+2); pc+=sBx

FORLOOP:
	R(A) += R(A+2)
	if R(A) <?= R(A+1) then {
		pc+=sBx
		R(A+3) = R(A)
	}
*/
func forPrep(i Instruction, vm api.LuaVM) {
	a, sBx := i.AsBx()
	a += 1

	// R(A) -= R(A+2)
	vm.PushValue(a)
	vm.PushValue(a + 2)
	vm.Arith(api.LuaOpSUB)
	vm.Replace(a)
	// pc += sBx
	vm.AddPC(sBx)
}

func forLoop(i Instruction, vm api.LuaVM) {
	a, sBx := i.AsBx()
	a += 1

	// R(A) := R(A+2)
	vm.PushValue(a + 2)
	vm.PushValue(a)
	vm.Arith(api.LuaOpADD)
	vm.Replace(a)

	// R(A) <?= R(A+1)
	isPositiveStep := vm.ToNumber(a+2) >= 0 // 步数是否为正数
	if isPositiveStep && vm.Compare(a, a+1, api.LuaOpLE) || !isPositiveStep && vm.Compare(a+1, a, api.LuaOpLE) {
		vm.AddPC(sBx)   // pc += sBx, sBx为负数
		vm.Copy(a, a+3) // R(A+3) = R(A)
	}
}
