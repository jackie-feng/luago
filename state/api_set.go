package state

func (self *luaState) setTable(table luaValue, key luaValue, value luaValue) {
	if tbl, ok := table.(*luaTable); ok {
		tbl.put(key, value)
		return
	}
	panic("not a table!")
}

func (self *luaState) SetTable(idx int) {
	t := self.stack.get(idx)
	v := self.stack.pop()
	k := self.stack.pop()
	self.setTable(t, k, v)
}

func (self *luaState) SetField(idx int, k string) {
	t := self.stack.get(idx)
	v := self.stack.pop()
	self.setTable(t, k, v)
}

func (self *luaState) SetI(idx int, i int64) {
	t := self.stack.get(idx)
	v := self.stack.pop()
	self.setTable(t, i, v)
}
