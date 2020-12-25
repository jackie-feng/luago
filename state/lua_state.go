package state

import (
	"fmt"
	"luago/binchunk"
	"strings"
)

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

func (self *luaState) popLuaStack() {
	stack := self.stack
	self.stack = stack.prev
	stack.prev = nil
}

func (self *luaState) String() string {
	stack := self.stack
	l := 0
	s := make([]string, 0)
	s = append(s, "==============state begin===============")
	for {
		if stack == nil {
			break
		}

		tabs := ""
		for i := 0; i < l; i++ {
			tabs = fmt.Sprintf("# %s", tabs)
		}
		stackString := stack.String()
		lines := strings.Split(stackString, "\n")
		newLines := make([]string, 0, len(lines))
		for _, line := range lines {
			newLines = append(newLines, fmt.Sprintf("%s%s", tabs, line))
		}
		s = append(s, strings.Join(newLines, "\n"))
		stack = stack.prev
		l = l + 1
	}
	s = append(s, "==============state   end===============")
	return strings.Join(s, "\n")
}
