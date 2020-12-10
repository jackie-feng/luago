package state

import "luago/binchunk"

type luaState struct {
	stack *luaStack
}

func New(stackSize int, proto *binchunk.Prototype) *luaState {
	return &luaState{
		stack: newLuaStack(stackSize),
	}
}

//
//func New() *luaState {
//	return &luaState{
//		stack: newLuaStack(20),
//	}
//}

func (self *luaState) pushLuaStack(stack *luaStack) {
	stack.prev = self.stack
	self.stack = stack
}

func (self *luaState) popLuaState() {
	stack := self.stack
	self.stack = stack.prev
	stack.prev = nil
}
