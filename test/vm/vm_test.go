package vm

import (
	"io/ioutil"
	"luago/state"
	"testing"
)

func TestVM(t *testing.T) {
	data, err := ioutil.ReadFile("sum.luac")
	if err != nil {
		panic(err)
	}
	luaMain(data)
}

func luaMain(chunk []byte) {
	ls := state.New(20, nil)
	ls.Load(chunk, "main", "b")
	ls.Call(0, 0)
	//for {
	//	pc := ls.PC()
	//	inst := vm.Instruction(ls.Fetch())
	//	if inst.Opcode() != vm.OpReturn {
	//		inst.Execute(ls)
	//		fmt.Printf("[%02d]\t%s\t", pc+1, inst.OpName())
	//		support.PrintStack(ls)
	//	} else {
	//		break
	//	}
	//}
}
