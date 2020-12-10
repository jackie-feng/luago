package api

type LuaType = int
type ArithOp = int
type CompareOp = int

type LuaState interface {
	/* Basic stack manipulation */
	GetTop() int
	AbsIndex(idx int) int
	CheckStack(n int) bool
	Pop(n int)
	Copy(fromIdx, toIdx int)
	PushValue(idx int)
	Replace(idx int)
	Insert(idx int)
	Remove(idx int)
	Rotate(idx int, n int)
	SetTop(idx int)
	/* access functions stack->go */
	TypeName(tp LuaType) string
	Type(idx int) LuaType
	IsNone(idx int) bool
	IsNil(idx int) bool
	IsNoneOrNil(idx int) bool
	IsBoolean(idx int) bool
	IsInteger(idx int) bool
	IsNumber(idx int) bool
	IsString(idx int) bool
	ToBoolean(idx int) bool
	ToInteger(idx int) int64
	ToIntegerX(idx int) (int64, bool)
	ToNumber(idx int) float64
	ToNumberX(idx int) (float64, bool)
	ToString(idx int) string
	ToStringX(idx int) (string, bool)
	/* push functions (GO -> Stack) */
	PushNil()
	PushBoolean(b bool)
	PushInteger(n int64)
	PushNumber(n float64)
	PushString(s string)

	/* Lua运算符相关 */
	Arith(op ArithOp)
	Compare(idx1, idx2 int, op CompareOp) bool
	Len(idx int)
	Concat(n int)

	/* get functions (lua -> stack) */
	// 无法预估容量, -> CreateTable(0, 0)
	NewTable()
	// 创建一个空的lua表， 推入栈顶
	CreateTable(nArr, nRec int)
	// idx 为表在栈中的位置
	// 根据键(栈顶弹出)从表(idx表示的位置)取值, 并将值写入到栈顶
	GetTable(idx int) LuaType
	// GetField, GetI 为 GetTable 的包装
	GetField(idx int, k string) LuaType
	GetI(idx int, i int64) LuaType
	/* set functions (stack -> lua) */
	// idx 为表在栈中的位置
	// value, key 从栈顶先后弹出, 并将值写入表(idx表示的位置)中
	SetTable(idx int)
	// 包装 SetTable, key 由外部传入
	SetField(idx int, k string)
	SetI(idx int, i int64)

	// 加载二进制 chunk, 把主函数原型实例化为闭包并推入栈顶
	// mode "b", "t", "bt": 二进制, 文本, 二进制/文本
	// return 0:成功
	Load(chunk []byte, chunkName, mode string) int
	// nArgs: 需要从栈顶弹出的入参数量
	// nResults: 返回值数量, -1 表示所有返回值留在栈顶
	Call(nArgs, nResults int)
}
