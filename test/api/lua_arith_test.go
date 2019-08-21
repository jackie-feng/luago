package api

import (
	"luago/api"
	"luago/state"
	"testing"
)

func TestArith(t *testing.T) {
	ls := state.New()

	ls.PushInteger(1)
	ls.PushString("2.0")
	ls.PushString("3.0")
	ls.PushNumber(4.0)
	printStack(ls)

	ls.Arith(api.LuaOpADD)
	printStack(ls)
	ls.Arith(api.LuaOpBNOT)
	printStack(ls)
	ls.Len(2)
	printStack(ls)
	ls.Concat(3)
	printStack(ls)
	ls.PushBoolean(ls.Compare(1, 2, api.LuaOpEQ))
	printStack(ls)
}
