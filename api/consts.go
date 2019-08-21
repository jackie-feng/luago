package api

const (
	LuaTNone = iota - 1
	LuaTNil
	LuaTBoolean
	LuaTLightUserData
	LuaTNumber
	LuaTString
	LuaTTable
	LuaTFunction
	LuaTUserData
	LuaTThread
)
const (
	LuaOpADD  = iota // +
	LuaOpSUB         // -
	LuaOpMUL         // *
	LuaOpMOD         // %
	LuaOpPOW         // ^
	LuaOpDIV         // /
	LuaOpIDIV        // //
	LuaOpBAND        // &
	LuaOpBOR         // |
	LuaOpBXOR        // ~
	LuaOpSHL         // <<
	LuaOpSHR         // >>
	LuaOpUNM         // - (unary minus)
	LuaOpBNOT        // ~
)

const (
	LuaOpEQ = iota // ==
	LuaOpLT        // <
	LuaOpLE        // <=
)
