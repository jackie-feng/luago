package api

import (
	"luago/api"
	"luago/state"
	"luago/test/support"
	"testing"
)

func TestArith(t *testing.T) {
	ls := state.New(20, nil)

	ls.PushInteger(1)
	ls.PushString("2.0")
	ls.PushString("3.0")
	ls.PushNumber(4.0)
	support.PrintStack(ls)

	ls.Arith(api.LuaOpADD)
	support.PrintStack(ls)
	ls.Arith(api.LuaOpBNOT)
	support.PrintStack(ls)
	ls.Len(2)
	support.PrintStack(ls)
	ls.Concat(3)
	support.PrintStack(ls)
	ls.PushBoolean(ls.Compare(1, 2, api.LuaOpEQ))
	support.PrintStack(ls)
}
