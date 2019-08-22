package vm

import (
	"fmt"
	"io/ioutil"
	"luago/binchunk"
	"luago/state"
	"luago/test/support"
	"luago/vm"
	"testing"
)

func TestVM(t *testing.T) {
	data, err := ioutil.ReadFile("luac.out")
	if err != nil {
		panic(err)
	}
	proto := binchunk.Undump(data)
	luaMain(proto)
}

func luaMain(proto *binchunk.Prototype) {
	nRegs := int(proto.MaxStackSize)
	ls := state.New(nRegs+8, proto)
	ls.SetTop(nRegs)
	for {
		pc := ls.PC()
		inst := vm.Instruction(ls.Fetch())
		if inst.Opcode() != vm.OpReturn {
			inst.Execute(ls)
			fmt.Printf("[%02d]\t%s\t", pc+1, inst.OpName())
			support.PrintStack(ls)
		} else {
			break
		}
	}
}
