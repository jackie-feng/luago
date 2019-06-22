package state

type luaState struct {
	stack *luaStack
}

func New() *luaState {
	return &LuaState{
		stack: newLuaStack(20),
	}
}
