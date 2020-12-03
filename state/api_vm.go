package state

func (self *luaState) PC() int {
	return self.pc
}

// 指令跳转
func (self *luaState) AddPC(n int) {
	self.pc += n
}

// 获取指令
func (self *luaState) Fetch() uint32 {
	i := self.proto.Code[self.pc]
	self.pc++
	return i
}

func (self *luaState) GetConst(idx int) {
	c := self.proto.Constants[idx]
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
