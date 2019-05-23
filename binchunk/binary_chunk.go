package binchunk

// chunk header 对应的常量
const (
	LuaSignature    = "\x1bLua"
	LuacVersion     = 0x53
	LuacFormat      = 0
	LuacData        = "\x19\x93\r\n\x1a\n"
	CintSize        = 4
	CsizetSize      = 8
	InstructionSize = 4
	LuaIntegerSize  = 8
	LuaNumberSize   = 8
	LuacInt         = 0x5678
	LuacNum         = 370.5
)

// 二进制chunk常量tag值
const (
	// 字面值类型nil
	TagNil = 0x00
	// 字面值类型boolean
	TagBoolean = 0x01
	// 字面值类型number
	TagNumber = 0x03
	// 字面值类型integer
	TagInteger = 0x13
	// 字面值类型string 短字符串
	TagShortStr = 0x04
	// 字面值类型string 长字符串
	TagLongStr = 0x14
)

type header struct {
	signature       [4]byte // Magic Number "\x1bLua"
	version         byte    // eg. 5.3.4, version = 5*16+3, 即 0x53
	format          byte    // 格式号, 官方格式号是0
	luacData        [6]byte // "\x19\x93\r\n\xla\n", 进一步校验
	cintSize        byte    // cint类型占用字节数
	sizetSize       byte    // size_t类型占用字节数
	instructionSize byte    // lua 虚拟机指令占用字节数
	luaIntegerSize  byte    // lua 整数占用字节数
	luaNumberSize   byte    // lua 浮点数占用字节数
	luacInt         int64   // 存放Ox5678, 校验大小端和字节数
	luacNum         float64 // 存放370.5, 校验浮点字节数
}

type binaryChunk struct {
	header                  // 头部
	sizeUpvalues byte       // 主函数upvalue数量
	mainFunc     *Prototype // 主函数原型
}

// Prototype 函数原型
type Prototype struct {
	Source          string        // 源文件名
	LineDefined     uint32        // 起止行号
	LastLineDefined uint32        // 起止行号
	NumParams       byte          // 函数固定参数个数
	IsVararg        byte          // 函数是否为 Vararg 函数
	MaxStackSize    byte          // 寄存器数量
	Code            []uint32      // 指令表, 每条指令四字节
	Constants       []interface{} // 常量表, 代码里出现的字面量, 包括 nil、布尔 值、整数、浮点数和字符串五种
	Upvalues        []Upvalue
	Protos          []*Prototype // 子函数原型表
	LineInfo        []uint32     // 行号表,行号表中的行号和指令表中的指令一一对应,分别记录每条指令在源代码中对应的行号
	LocVars         []LocVar     // 局部变量表
	UpvalueNames    []string     // Upvalue名列表
}

// Upvalue ...
type Upvalue struct {
	Instack byte
	Idx     byte
}

// LocVar 局部变量表
type LocVar struct {
	VarName string // 变量名
	StartPC uint32 // 起止指令索引
	EndPC   uint32 // 起止指令索引
}

// Undump 解析二进制chunk
func Undump(data []byte) *Prototype {
	reader := &reader{data}
	reader.checkHeader()        // 校验头部
	reader.readByte()           // 跳过Upvalue数量
	return reader.readProto("") // 读取函数原型
}
