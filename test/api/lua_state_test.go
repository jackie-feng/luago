package api

import (
	"luago/state"
	"luago/test/support"
	"testing"
)

func TestLuaState(t *testing.T) {
	ls := state.New(20, nil)
	ls.PushBoolean(true)
	support.PrintStack(ls)
	ls.PushInteger(10)
	support.PrintStack(ls)
	ls.PushNil()
	support.PrintStack(ls)
	ls.PushString("hello")
	support.PrintStack(ls)
	ls.PushValue(-4)
	support.PrintStack(ls)
	ls.Replace(3)
	support.PrintStack(ls)
	ls.SetTop(6)
	support.PrintStack(ls)
	ls.Remove(-3)
	support.PrintStack(ls)
	ls.SetTop(-5)
	support.PrintStack(ls)
}
