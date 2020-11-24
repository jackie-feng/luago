package state

// 索引从 1 开始
type luaStack struct {
	slots []luaValue
	// 栈顶的下标 + 1
	// 等于下一个推入元素的位置
	top   int
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
