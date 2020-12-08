package vm

import (
	"luago/api"
)

// CLOSURE
// 把当前 Lua 函数的子函数原型实例化为闭包，放入由操作数 A 指定的寄存器中。
// 子函数原型来自当前函数原型的子函数原型表，索引由操作数 Bx 指定。
// R(A) := closure(KPROTO[Bx])
func closure(i Instruction, vm api.LuaVM) {
	a, bx := i.ABx()
	a += 1
	vm.LoadProto(bx)
	vm.Replace(a)
}

// CALL
// 调用 Lua 函数。其中被调函数位于寄存器中，索引由操作数 A 指定。
// 需要传递给被调函数的参数值也在寄存器中，紧挨着被调函数，数量由操作数 B 指定。
// 函数调用结束后，原先存放在函数和参数值的寄存器会被返回值占据，具体多少个返回值由操作数 C 指定。
// R(A), ..., R(A+C-2) := R(A)( R(A+1), ... , R(A+B-1) )
func call(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1

	nArgs := _pushFuncAndArgs(a, b, vm)
	vm.Call(nArgs, c-1)
	_popResults(a, c, vm)
}

// RETURN
// 把存放在连续多个寄存器的值返回给主调函数。
// 其中第一个寄存器的索引由操作数 A 指定，寄存器数量由操作数 B 指定，操作数 C 没用。
// return R(A), ... , R(A+B-2)
func _return(i Instruction, vm api.LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	if b == 1 {
		// no return values
	} else if b > 1 {
		// b-1 return values
		vm.CheckStack(b - 1)
		for i := a; i <= a+b-2; i++ {
			vm.PushValue(i)
		}
	} else { // b == 0
		_fixStack(a, vm)
	}
}

// VARARG
// 把传递给当前函数的变长参数加载到连续多个寄存器中。
// 其中第一个寄存器的索引由操作数 A 指定，寄存器数量由操作数 B 指定，操作数 C 没有用。
// R(A), R(A+1), ... , R(A+B-2) = vararg
func vararg(i Instruction, vm api.LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	if b != 1 {
		vm.LoadVararg(b - 1)
		_popResults(a, b, vm)
	}
}

// TAILCALL
// 尾递归优化
// return R(A)( R(A+1), ... , R(A+B-1) )
func tailCall(i Instruction, vm api.LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	c := 0

	nArgs := _pushFuncAndArgs(a, b, vm)
	vm.Call(nArgs, c-1)
	_popResults(a, c, vm)
}

// SELF
// 把对象和方法拷贝到相邻到两个目标寄存器中。
// 对象在寄存器中，索引由操作数 B 指定。
// 方法名在常量表里，索引由操作数 C 指定。
// 目标寄存器索引由操作数 A 指定。
// R(A+1) := R(B);
// R(A) := R(B)[RK(C)]
func self(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1
	b += 1
	// 复制对象到 a+1
	vm.Copy(b, a+1)
	// 读取方法名到栈顶
	vm.GetRK(c)
	// 从栈顶弹出方法名, 并从 b 处的 table 弹出具体方法到栈顶
	vm.GetTable(b)
	// 把栈顶的方法写回到 a 处
	vm.Replace(a)
}

func _popResults(a int, c int, vm api.LuaVM) {
	if c == 1 {
		// no results
	} else if c > 1 { // c-1 results
		for i := a + c - 2; i >= a; i-- {
			vm.Replace(i)
		}
	} else {
		// c == 0 需要把被调函数的返回值全部返回
		vm.CheckStack(1)
		vm.PushInteger(int64(a))
	}
}

func _pushFuncAndArgs(a int, b int, vm api.LuaVM) int {
	if b >= 1 { // b - 1 args
		vm.CheckStack(b) // 扩容以安置 closure 及参数
		for i := a; i < a+b; i++ {
			vm.PushValue(i)
		}
		return b - 1
	} else {
		_fixStack(a, vm)
		return vm.GetTop() - vm.RegisterCount() - 1
	}
}

func _fixStack(a int, vm api.LuaVM) {
	x := int(vm.ToInteger(-1))
	vm.Pop(1)
	vm.CheckStack(x - a)
	for i := a; i < x; i++ {
		vm.PushValue(i)
	}
	vm.Rotate(vm.RegisterCount()+1, x-a)
}
