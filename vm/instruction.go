package vm

import "luago/api"

// MaxArgBX BX 无符号数最大值
const MaxArgBX = 1<<18 - 1

// MaxArgsBx BX 无符号数最大值一半
const MaxArgsBx = MaxArgBX >> 1

// Instruction 指令
type Instruction uint32

// Opcode 返回指令码
func (i Instruction) Opcode() int {
	return int(i & 0x3F)
}

// ABC return a,b,c for IABC
func (i Instruction) ABC() (a, b, c int) {
	a = int(i >> 6 & 0xFF)
	c = int(i >> 14 & 0x1FF)
	b = int(i >> 23 & 0x1FF)
	return
}

// ABx return a, bx for IAbx
func (i Instruction) ABx() (a, bx int) {
	a = int(i >> 6 & 0xFF)
	bx = int(i >> 14)
	return
}

// AsBx return a, sbx for IAsBx
func (i Instruction) AsBx() (a, sbx int) {
	a, bx := i.ABx()
	return a, bx - MaxArgsBx
}

// Ax return ax for IAx
func (i Instruction) Ax() int {
	return int(i >> 6)
}

// OpName 返回指令名
func (i Instruction) OpName() string {
	return opcodes[i.Opcode()].name
}

// OpMode 返回指令模式
func (i Instruction) OpMode() byte {
	return opcodes[i.Opcode()].opMode
}

// BMode 返回B操作数类型
func (i Instruction) BMode() byte {
	return opcodes[i.Opcode()].argBMode
}

// CMode 返回C操作数类型
func (i Instruction) CMode() byte {
	return opcodes[i.Opcode()].argCMode

}

func (i Instruction) Execute(vm api.LuaVM) {
	action := opcodes[i.Opcode()].action
	if action != nil {
		action(i, vm)
	} else {
		panic(i.OpName())
	}
}
