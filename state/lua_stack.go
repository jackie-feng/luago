package state

import (
	"fmt"
	"strings"
	"unsafe"
)

// 索引从 1 开始
type luaStack struct {
	slots []luaValue
	// 栈顶的下标 + 1
	// 等于下一个推入元素的位置
	top     int
	prev    *luaStack
	closure *closure // 闭包
	varargs []luaValue
	pc      int
}

func (ls *luaStack) String() string {
	top := ls.top
	tmp := make([]string, 0, top)
	for idx, val := range ls.slots {
		if idx+1 == ls.top {
			tmp = append(tmp, "🌟"+stringOf(val))
		} else {
			tmp = append(tmp, stringOf(val))
		}
	}
	varargsStrings := make([]string, 0, len(ls.varargs))
	for _, val := range ls.varargs {
		varargsStrings = append(varargsStrings, stringOf(val))
	}
	return strings.Join([]string{
		"--------------stack begin---------------",
		fmt.Sprintf("stack: %v, registerCount: %d", unsafe.Pointer(ls), ls.registerCount()),
		fmt.Sprintf("pc: %d, top: %d", ls.pc, ls.top),
		fmt.Sprintf("slots len: %d", len(ls.slots)),
		fmt.Sprintf("slots: %s", strings.Join(tmp, ",")),
		fmt.Sprintf("varargs len: %d", len(ls.varargs)),
		fmt.Sprintf("varargs: %s", strings.Join(varargsStrings, ",")),
		"--------------stack   end---------------",
	}, "\n")
}
func (ls *luaStack) registerCount() int {
	if ls.closure != nil {
		return int(ls.closure.proto.MaxStackSize)
	}
	return 0
}

func newLuaStack(size int) *luaStack {
	return &luaStack{
		slots: make([]luaValue, size),
		top:   0,
	}
}

// 扩容使剩余空间到达n
func (s *luaStack) check(n int) {
	free := len(s.slots) - s.top
	for i := free; i < n; i++ {
		s.slots = append(s.slots, nil)
	}
}

func (s *luaStack) push(val luaValue) {
	if s.top == len(s.slots) {
		panic("stack overflow!")
	}
	// TODO 这边插值和自增的顺序?
	s.slots[s.top] = val
	s.top++
}

func (s *luaStack) pop() luaValue {
	if s.top < 1 {
		panic("stack underflow!")
	}
	s.top--
	val := s.slots[s.top]
	s.slots[s.top] = nil
	return val
}

func (s *luaStack) popN(n int) []luaValue {
	v := make([]luaValue, n)
	for i := n - 1; i >= 0; i-- {
		v[i] = s.pop()
	}
	return v
}

func (s *luaStack) pushN(vals []luaValue, n int) {
	nVals := len(vals)
	if n < 0 {
		n = nVals
	}
	for i := 0; i < n; i++ {
		if i < nVals {
			s.push(vals[i])
		} else {
			s.push(nil)
		}
	}
}

// 索引转化为绝对索引
// -1 代表第一个
func (s *luaStack) absIndex(idx int) int {
	if idx >= 0 {
		return idx
	}

	return idx + s.top + 1
}

// 索引是否有效: > 0 && <= top
func (s *luaStack) isValid(idx int) bool {
	absIdx := s.absIndex(idx)
	return absIdx > 0 && absIdx <= s.top
}

func (s *luaStack) get(idx int) luaValue {
	absIndex := s.absIndex(idx)
	if absIndex > 0 && absIndex <= s.top {
		return s.slots[absIndex-1]
	}
	return nil
}

func (s *luaStack) set(idx int, val luaValue) {
	absIndex := s.absIndex(idx)
	if absIndex > 0 && absIndex <= s.top {
		s.slots[absIndex-1] = val
		return
	}
	panic("invalid index!")
}

func (s *luaStack) reverse(from, to int) {
	slots := s.slots
	for from < to {
		slots[from], slots[to] = slots[to], slots[from]
		from++
		to--
	}
}
