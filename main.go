package main

import (
	"fmt"
	"io/ioutil"
	"luago/binchunk"
	"luago/vm"
	"os"
)

func main() {
	fmt.Println("hello, luago")
	if len(os.Args) > 1 {
		data, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		proto := binchunk.Undump(data)
		list(proto)
	}
}

func list(f *binchunk.Prototype) {
	printHeader(f)
	printCode(f)
	printDetail(f)
	for _, p := range f.Protos {
		list(p)
	}
}

func printHeader(f *binchunk.Prototype) {
	funcType := "main"
	if f.LineDefined > 0 {
		funcType = "function"
	}

	varargFlag := ""
	if f.IsVararg > 0 {
		varargFlag = "+"
	}

	fmt.Printf("\n%s <%s:%d,%d> (%d instructions)\n", funcType, f.Source, f.LineDefined, f.LastLineDefined, len(f.Code))
	fmt.Printf("%d%s params, %d slots, %d upvalues,", f.NumParams, varargFlag, f.MaxStackSize, len(f.Upvalues))
	fmt.Printf("%d locals, %d constants, %d functions\n", len(f.LocVars), len(f.Constants), len(f.Protos))
}

func printCode(f *binchunk.Prototype) {
	for pc, c := range f.Code {
		line := "-"
		if len(f.LineInfo) > 0 {
			line = fmt.Sprintf("%d", f.LineInfo[pc])
		}

		i := vm.Instruction(c)
		fmt.Printf("\t%d\t[%s]\t%s\t", pc+1, line, i.OpName())
		printOperands(i)
		fmt.Printf("\n")
	}
}

func printOperands(i vm.Instruction) {
	switch i.OpMode() {
	case vm.IABC:
		a, b, c := i.ABC()
		fmt.Printf("%d", a)
		if i.BMode() != vm.OpArgN {
			if b > 0xFF {
				fmt.Printf(" %d", -1-b&0xFF)
			} else {
				fmt.Printf(" %d", b)
			}
		}
		if i.CMode() != vm.OpArgN {
			if c > 0xFF {
				fmt.Printf(" %d", -1-c&0xFF)
			} else {
				fmt.Printf(" %d", c)
			}
		}
	case vm.IABx:
		a, bx := i.ABx()
		fmt.Printf("%d", a)
		if i.BMode() == vm.OpArgK {
			fmt.Printf(" %d", -1-bx)
		} else if i.BMode() == vm.OpArgU {
			fmt.Printf(" %d", bx)
		}
	case vm.IAsBx:
		a, sbx := i.AsBx()
		fmt.Printf("%d %d", a, sbx)
	case vm.IAx:
		ax := i.Ax()
		fmt.Printf("%d", -1-ax)
	}
}

func printDetail(f *binchunk.Prototype) {
	fmt.Printf("constants (%d):\n", len(f.Constants))
	for i, k := range f.Constants {
		fmt.Printf("\t%d\t%s\n", i+1, constantToString(k))
	}

	fmt.Printf("locals (%d):\n", len(f.LocVars))
	for i, locVar := range f.LocVars {
		fmt.Printf("\t%d\t%s\t%d\t%d\n", i, locVar.VarName, locVar.StartPC, locVar.EndPC)
	}

	fmt.Printf("upvalues (%d):\n", len(f.Upvalues))
	for i, upval := range f.Upvalues {
		fmt.Printf("\t%d\t%s\t%d\t%d\n", i, upvalName(f, i), upval.Instack, upval.Idx)
	}
}

func constantToString(k interface{}) string {
	switch k.(type) {
	case nil:
		return "nil"
	case bool:
		return fmt.Sprintf("%t", k)
	case float64:
		return fmt.Sprintf("%s", k)
	case int64:
		return fmt.Sprintf("%d", k)
	case string:
		return fmt.Sprintf("%s", k)
	default:
		return "?"
	}
}

func upvalName(f *binchunk.Prototype, idx int) string {
	if len(f.UpvalueNames) > 0 {
		return f.UpvalueNames[idx]
	}
	return "-"
}
