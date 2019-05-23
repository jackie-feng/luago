package vm

// 操作码四种模式
const (
	IABC = iota
	IABx
	IAsBx
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
}

// 指令表
var opcodes = []opcode{
	/*     T  A  B       C       mode  name */
	opcode{0, 1, OpArgR, OpArgN, IABC, "MOVE"},
	opcode{0, 1, OpArgK, OpArgN, IABx, "LOADK"},
	opcode{0, 1, OpArgN, OpArgN, IABx, "LOADKX"},
	opcode{0, 1, OpArgU, OpArgU, IABC, "LOADBOOL"},
	opcode{0, 1, OpArgU, OpArgN, IABC, "LOADNIL"},
	opcode{0, 1, OpArgU, OpArgN, IABC, "GETUPVAL"},
	opcode{0, 1, OpArgU, OpArgK, IABC, "GETTABUP"},
	opcode{0, 1, OpArgR, OpArgK, IABC, "GETTABLE"},
	opcode{0, 0, OpArgK, OpArgK, IABC, "SETTABUP"},
	opcode{0, 0, OpArgU, OpArgN, IABC, "SETUPVAL"},
	opcode{0, 0, OpArgK, OpArgK, IABC, "SETTABLE"},
	opcode{0, 1, OpArgU, OpArgU, IABC, "NEWTABLE"},
	opcode{0, 1, OpArgR, OpArgK, IABC, "SELF"},
	opcode{0, 1, OpArgK, OpArgK, IABC, "ADD"},
	opcode{0, 1, OpArgK, OpArgK, IABC, "SUB"},
	opcode{0, 1, OpArgK, OpArgK, IABC, "MUL"},
	opcode{0, 1, OpArgK, OpArgK, IABC, "MOD"},
	opcode{0, 1, OpArgK, OpArgK, IABC, "POW"},
	opcode{0, 1, OpArgK, OpArgK, IABC, "DIV"},
	opcode{0, 1, OpArgK, OpArgK, IABC, "IDIV"},
	opcode{0, 1, OpArgK, OpArgK, IABC, "BAND"},
	opcode{0, 1, OpArgK, OpArgK, IABC, "BOR"},
	opcode{0, 1, OpArgK, OpArgK, IABC, "BXOR"},
	opcode{0, 1, OpArgK, OpArgK, IABC, "SHL"},
	opcode{0, 1, OpArgK, OpArgK, IABC, "SHR"},
	opcode{0, 1, OpArgR, OpArgN, IABC, "UNM"},
	opcode{0, 1, OpArgR, OpArgN, IABC, "BNOT"},
	opcode{0, 1, OpArgR, OpArgN, IABC, "NOT"},
	opcode{0, 1, OpArgR, OpArgN, IABC, "LEN"},
	opcode{0, 1, OpArgR, OpArgR, IABC, "CONCAT"},
	opcode{0, 0, OpArgR, OpArgN, IAsBx, "JMP"},
	opcode{1, 0, OpArgK, OpArgK, IABC, "EQ"},
	opcode{1, 0, OpArgK, OpArgK, IABC, "LT"},
	opcode{1, 0, OpArgK, OpArgK, IABC, "LE"},
	opcode{1, 0, OpArgN, OpArgU, IABC, "TEST"},
	opcode{1, 1, OpArgR, OpArgU, IABC, "TESTSET"},
	opcode{0, 1, OpArgU, OpArgU, IABC, "CALL"},
	opcode{0, 1, OpArgU, OpArgU, IABC, "TALLCALL"},
	opcode{0, 0, OpArgU, OpArgN, IABC, "RETURN"},
	opcode{0, 1, OpArgR, OpArgN, IAsBx, "FORLOOP"},
	opcode{0, 1, OpArgR, OpArgN, IAsBx, "FORPREP"},
	opcode{0, 0, OpArgN, OpArgU, IABC, "TFORCALL"},
	opcode{0, 1, OpArgR, OpArgN, IAsBx, "TFORLOOP"},
	opcode{0, 0, OpArgU, OpArgU, IABC, "SETLIST"},
	opcode{0, 1, OpArgU, OpArgN, IABx, "CLOSURE"},
	opcode{0, 1, OpArgU, OpArgN, IABC, "VARARG"},
	opcode{0, 0, OpArgU, OpArgU, IAx, "EXTRAARG"},
}
