package state

/* get functions (lua -> stack) */
func (self *luaState) NewTable() {
	self.CreateTable(0, 0)
}

func (self *luaState) CreateTable(nArr, nRec int) {
	t := newLuaTable(nArr, nRec)
	self.stack.push(t)
}

func (self *luaState) GetTable(idx int) LuaType {
	t := self.stack.get(idx)
	k := self.stack.pop()
	return self.getTable(t, k)
}

func (self *luaState) GetField(idx int, k string) LuaType {
	//也可以如下方法, 但是直接取值更加高效
	//self.PushString(k)
	//return self.GetTable(idx)
	t := self.stack.get(idx)
	return self.getTable(t, k)
}

func (self *luaState) GetI(idx int, i int64) LuaType {
	t := self.stack.get(idx)
	return self.getTable(t, i)
}

// 取出 table[key], 将其推入栈顶, 并返回值类型
func (self *luaState) getTable(table, key luaValue) LuaType {
	if tbl, ok := table.(*luaTable); ok {
		v := tbl.get(key)
		self.stack.push(v)
		return typeOf(v)
	}
	panic("not a table!")
}
