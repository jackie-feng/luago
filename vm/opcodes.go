package vm

import "luago/api"

// 操作码四种模式
const (
	IABC = iota
	IABx
	IAsBx // A, sBx: sBx 为有符号数
	IAx
)

// 操作码
const (
	OpMove = iota
	OpLoadK
	OpLoadKX
	OpLoadBool
	OpLoadNil
	OpGetUpVal
	OpGetTabUp
	OpGetTable
	OpSetTabUp
	OpSetUpVal
	OpSetTable
	OpNewTable
	OpSelf
	OpAdd
	OpSub
	OpMul
	OpMod
	OpPow
	OpDiv
	OpIDiv
	OpBand
	OpBor
	OpBxor
	OpShl
	OpShr
	OpUnm
	OpBnot
	OpNot
	OpLen
	OpConcat
	OpJmp
	OpEq
	OpLt
	OpLe
	OpTest
	OpTestSet
	OpCall
	OpTallCall
	OpReturn
	OpForLoop
	OpForPrep
	OpTForCall
	OpTForLoop
	OpSetList
	OpClosure
	OpVararg
	OpExtraArg
)

// 操作数类型
const (
	OpArgN = iota // argument is not used
	OpArgU        // argument is used
	OpArgR        // argument is a register or a jump offest
	OpArgK        // argument is a constant or register/constant
)

type opcode struct {
	testFlag byte // operator is a test (next instruction must be a jump)
	setAFlag byte // instruction set register A
	argBMode byte // B arg mode
	argCMode byte // C arg mode
	opMode   byte // op mode
	name     string
	action   func(i Instruction, vm api.LuaVM)
}

// 指令表
var opcodes = []opcode{
	/*     T  A  B       C       mode  name    action*/
	opcode{0, 1, OpArgR, OpArgN, IABC, "MOVE", move},
	opcode{0, 1, OpArgK, OpArgN, IABx, "LOADK", loadK},
	opcode{0, 1, OpArgN, OpArgN, IABx, "LOADKX", loadKx},
	opcode{0, 1, OpArgU, OpArgU, IABC, "LOADBOOL", loadBool},
	opcode{0, 1, OpArgU, OpArgN, IABC, "LOADNIL", loadNil},
	opcode{0, 1, OpArgU, OpArgN, IABC, "GETUPVAL", nil},
	opcode{0, 1, OpArgU, OpArgK, IABC, "GETTABUP", nil},
	opcode{0, 1, OpArgR, OpArgK, IABC, "GETTABLE", getTable},
	opcode{0, 0, OpArgK, OpArgK, IABC, "SETTABUP", nil},
	opcode{0, 0, OpArgU, OpArgN, IABC, "SETUPVAL", nil},
	opcode{0, 0, OpArgK, OpArgK, IABC, "SETTABLE", setTable},
	opcode{0, 1, OpArgU, OpArgU, IABC, "NEWTABLE", newTable},
	opcode{0, 1, OpArgR, OpArgK, IABC, "SELF", self},
	opcode{0, 1, OpArgK, OpArgK, IABC, "ADD", add},
	opcode{0, 1, OpArgK, OpArgK, IABC, "SUB", sub},
	opcode{0, 1, OpArgK, OpArgK, IABC, "MUL", mul},
	opcode{0, 1, OpArgK, OpArgK, IABC, "MOD", mod},
	opcode{0, 1, OpArgK, OpArgK, IABC, "POW", pow},
	opcode{0, 1, OpArgK, OpArgK, IABC, "DIV", div},
	opcode{0, 1, OpArgK, OpArgK, IABC, "IDIV", idiv},
	opcode{0, 1, OpArgK, OpArgK, IABC, "BAND", band},
	opcode{0, 1, OpArgK, OpArgK, IABC, "BOR", bor},
	opcode{0, 1, OpArgK, OpArgK, IABC, "BXOR", bxor},
	opcode{0, 1, OpArgK, OpArgK, IABC, "SHL", shl},
	opcode{0, 1, OpArgK, OpArgK, IABC, "SHR", shr},
	opcode{0, 1, OpArgR, OpArgN, IABC, "UNM", unm},
	opcode{0, 1, OpArgR, OpArgN, IABC, "BNOT", bnot},
	opcode{0, 1, OpArgR, OpArgN, IABC, "NOT", not},
	opcode{0, 1, OpArgR, OpArgN, IABC, "LEN", _len},
	opcode{0, 1, OpArgR, OpArgR, IABC, "CONCAT", concat},
	opcode{0, 0, OpArgR, OpArgN, IAsBx, "JMP", jmp},
	opcode{1, 0, OpArgK, OpArgK, IABC, "EQ", eq},
	opcode{1, 0, OpArgK, OpArgK, IABC, "LT", lt},
	opcode{1, 0, OpArgK, OpArgK, IABC, "LE", le},
	opcode{1, 0, OpArgN, OpArgU, IABC, "TEST", test},
	opcode{1, 1, OpArgR, OpArgU, IABC, "TESTSET", testset},
	opcode{0, 1, OpArgU, OpArgU, IABC, "CALL", call},
	opcode{0, 1, OpArgU, OpArgU, IABC, "TALLCALL", tailCall},
	opcode{0, 0, OpArgU, OpArgN, IABC, "RETURN", _return},
	opcode{0, 1, OpArgR, OpArgN, IAsBx, "FORLOOP", forLoop},
	opcode{0, 1, OpArgR, OpArgN, IAsBx, "FORPREP", forPrep},
	opcode{0, 0, OpArgN, OpArgU, IABC, "TFORCALL", nil},
	opcode{0, 1, OpArgR, OpArgN, IAsBx, "TFORLOOP", nil},
	opcode{0, 0, OpArgU, OpArgU, IABC, "SETLIST", setList},
	opcode{0, 1, OpArgU, OpArgN, IABx, "CLOSURE", closure},
	opcode{0, 1, OpArgU, OpArgN, IABC, "VARARG", vararg},
	opcode{0, 0, OpArgU, OpArgU, IAx, "EXTRAARG", nil},
}
