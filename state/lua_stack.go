package state

type luaStack struct {
	slots []luaValue
	// top 栈顶指针(从1开始), 下标 + 1
	top int
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

// 索引转化为绝对索引: 下标 + 1
func (s *luaStack) absIndex(idx int) int {
	if idx >= 0 {
		return idx
	}

	return idx + s.top + 1
}

// 对于 slots len = n, 有效索引 [1, n]
func (s *luaStack) isValid(idx int) bool {
	absIdx := s.absIndex(idx)
	return absIdx > 0 && absIdx <= s.top
}

// 有效索引 = 下标 + 1
func (s *luaStack) get(idx int) luaValue {
	absIndex := s.absIndex(idx)
	if absIndex > 0 && absIndex <= s.top {
		return s.slots[absIndex-1]
	}
	return nil
}

// @idx: 有效索引
func (s *luaStack) set(idx int, val luaValue) {
	absIndex := s.absIndex(idx)
	if absIndex > 0 && absIndex <= s.top {
		s.slots[absIndex-1] = val
		return
	}
	panic("invalid index!")
}

// [from ... to]
func (s *luaStack) reverse(from, to int) {
	slots := s.slots
	for from < to {
		slots[from], slots[to] = slots[to], slots[from]
		from++
		to--
	}
}
