package function

import (
	"fmt"
	"io/ioutil"
	"luago/binchunk"
	"luago/state"
	"luago/test/support"
	"luago/vm"
	"os"
	"testing"
)

func TestFunction(t *testing.T) {
	if len(os.Args) > 1 {
		data, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		ls := state.New(20, nil)
		ls.Load(data, os.Args[1], "b")
		ls.Call(0, 0)
	}

}

func TestVM(t *testing.T) {
	data, err := ioutil.ReadFile("function_test.luac")
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
