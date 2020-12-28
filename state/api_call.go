package state

import (
	"luago/binchunk"
	"luago/vm"
)

func (self *luaState) Load(chunk []byte, chunkName, mode string) int {
	proto := binchunk.Undump(chunk)
	c := newLuaClosure(proto)
	self.stack.push(c)
	return 0
}

func (self *luaState) Call(nArgs, nResults int) {
	val := self.stack.get(-(nArgs + 1))
	if c, ok := val.(*closure); ok {
		self.callLuaClosure(nArgs, nResults, c)
	} else {
		panic("not function!")
	}
}

// nResults: 主调帧需要的返回值数量
func (self *luaState) callLuaClosure(nArgs int, nResults int, c *closure) {
	// 寄存器数量
	nRegs := int(c.proto.MaxStackSize)
	// 函数固定参数个数
	nParams := int(c.proto.NumParams)
	isVararg := c.proto.IsVararg == 1
	newStack := newLuaStack(nRegs + 20)
	newStack.closure = c

	// 主调栈弹出参数, 多退少补写入被调栈
	funcAndArgs := self.stack.popN(nArgs + 1)
	newStack.pushN(funcAndArgs[1:], nParams)
	// TODO: top = nRegs ?
	newStack.top = nRegs
	if nArgs > nParams && isVararg {
		newStack.varargs = funcAndArgs[nParams+1:]
	}

	self.pushLuaStack(newStack)
	self.runLuaClosure()
	self.popLuaStack()

	if nResults != 0 {
		// TODO: newstack.top - nRegs
		results := newStack.popN(newStack.top - nRegs)
		self.stack.check(len(results))
		self.stack.pushN(results, nResults)
	}
}

func (self *luaState) runLuaClosure() {
	for {
		inst := vm.Instruction(self.Fetch())
		inst.Execute(self)

		if inst.Opcode() == vm.OpReturn {
			break
		}
	}
}
