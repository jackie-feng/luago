package state

// 获取栈顶
func (self *luaState) GetTop() int {
	return self.stack.top
}

// 转化为绝对索引
func (self *luaState) AbsIndex(idx int) int {
	return self.stack.absIndex(idx)
}

func (self *luaState) CheckStack(n int) bool {
	// TODO: 扩容失败？
	self.stack.check(n)
	return true
}

func (self *luaState) Pop(n int) {
	for i := 0; i < n; i++ {
		self.stack.pop()
	}
}

// 把 fromIdx 处复制到 toIdx 处
func (self *luaState) Copy(fromIdx, toIdx int) {
	val := self.stack.get(fromIdx)
	self.stack.set(toIdx, val)
}

// 把 idx 处元素推入栈顶
func (self *luaState) PushValue(idx int) {
	val := self.stack.get(idx)
	self.stack.push(val)
}

// PushValue 反操作, 栈顶弹出写入指定位置
func (self *luaState) Replace(idx int) {
	val := self.stack.pop()
	self.stack.set(idx, val)
}

// idx ～ top 之间的元素朝栈顶方向移动一个位置
// 原 top 处元素弹出, 并写入到 idx 处
func (self *luaState) Insert(idx int) {
	self.Rotate(idx, 1)
}

// 删除 idx 元素, idx 上元素向下填补
func (self *luaState) Remove(idx int) {
	self.Rotate(idx, -1)
	self.Pop(1)
}

// 旋转, 找到中位点 m, p-m-t
// 两边逆序, p<->m, m+1<->t, 再 p<->t
func (self *luaState) Rotate(idx, n int) {
	t := self.stack.top - 1
	p := self.stack.absIndex(idx) - 1
	var m int
	if n >= 0 {
		m = t - n
	} else {
		m = p - n - 1
	}
	self.stack.reverse(p, m)
	self.stack.reverse(m+1, t)
	self.stack.reverse(p, t)
}

// 设置栈顶位置
// 弹出多余元素, 或者以 nil 值填补元素
func (self *luaState) SetTop(idx int) {
	newTop := self.stack.absIndex(idx)
	if newTop < 0 {
		panic("stack underflow!")
	}

	n := self.stack.top - newTop
	if n > 0 {
		for i := 0; i < n; i++ {
			self.stack.pop()
		}
	} else if n < 0 {
		for i := 0; i > n; i-- {
			self.stack.push(nil)
		}
	}
}
