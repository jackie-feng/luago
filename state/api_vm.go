package state

func (self *luaState) PC() int {
	return self.stack.pc
}

// 指令跳转
func (self *luaState) AddPC(n int) {
	self.stack.pc += n
}

// 获取指令
func (self *luaState) Fetch() uint32 {
	i := self.stack.closure.proto.Code[self.stack.pc]
	self.stack.pc++
	return i
}

func (self *luaState) GetConst(idx int) {
	c := self.stack.closure.proto.Constants[idx]
	self.stack.push(c)
}

func (self *luaState) GetRK(rk int) {
	if rk > 0xff { // constant
		self.GetConst(rk & 0xff)
	} else { //register
		// 寄存器索引 + 1 == 栈索引
		self.PushValue(rk + 1)
	}
}

func (self *luaState) RegisterCount() int {
	return int(self.stack.closure.proto.MaxStackSize)
}

func (self *luaState) LoadVararg(n int) {
	if n < 0 {
		n = len(self.stack.varargs)
	}
	self.stack.check(n)
	self.stack.pushN(self.stack.varargs, n)
}

func (self *luaState) LoadProto(idx int) {
	// 从子函数原型表获取函数
	proto := self.stack.closure.proto.Protos[idx]
	closure := newLuaClosure(proto)
	self.stack.push(closure)
}
