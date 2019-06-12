package state

import (
	"fmt"
	"luago/api"
)

func (self *luaState) TypeName(tp api.LuaType) string {
	switch tp {
	case api.LuaTNone:
		return "no value"
	case api.LuaTNil:
		return "nil"
	case api.LuaTBoolean:
		return "boolean"
	case api.LuaTNumber:
		return "number"
	case api.LuaTString:
		return "string"
	case api.LuaTTable:
		return "table"
	case api.LuaTFunction:
		return "function"
	case api.LuaTThread:
		return "thread"
	default:
		return "userdata"
	}
}

func (self *luaState) Type(idx int) LuaType {
	if self.stack.isValid(idx) {
		val := self.stack.get(idx)
		return typeOf(val)
	}
	return api.LuaTNone
}

func (self *luaState) IsNone(idx int) bool {
	return self.Type(idx) == api.LuaTNone
}

func (self *luaState) IsNil(idx int) bool {
	return self.Type(idx) == api.LuaTNil
}

func (self *luaState) IsNoneOrNil(idx int) bool {
	return self.Type(idx) <= api.LuaTNil
}

func (self *luaState) IsBoolean(idx int) bool {
	return self.Type(idx) == api.LuaTBoolean
}

func (self *luaState) IsString(idx int) bool {
	t := self.Type(idx)
	return t == api.LuaTString || t == api.LuaTNumber
}

func (self *luaState) IsNumber(idx int) bool {
	_, ok := self.ToNumberX(idx)
	return ok
}

func (self *luaState) IsInteger(idx int) bool {
	val := self.stack.get(idx)
	_, ok := val.(int64)
	return ok
}

func (self *luaState) ToBoolean(idx int) bool {
	val := self.stack.get(idx)
	return convertToBoolean(val)
}

func (self *luaState) ToNumber(idx int) float64 {
	n, _ := self.ToNumberX(idx)
	return n
}

func (self *luaState) ToNumberX(idx int) (float64, bool) {

	val := self.stack.get(idx)
	switch x := val.(type) {
	case float64:
		return x, true
	case int64:
		return float64(x), true
	default:
		return 0, false
	}
}

func (self *luaState) ToInteger(idx int) int64 {
	i, _ := self.ToIntegerX(idx)
	return i
}

func (self *luaState) ToIntegerX(idx int) (int64, bool) {
	val := self.stack.get(idx)
	i, ok := val.(int64)
	return i, ok
}

func (self *luaState) ToString(idx int) string {
	s, _ := self.ToStringX(idx)
	return s
}

func (self *luaState) ToStringX(idx int) (string, bool) {
	val := self.stack.get(idx)
	switch x := val.(type) {
	case string:
		return x, true
	case int64, float64:
		s := fmt.Sprintf("%v", x)
		self.stack.set(idx, s) // TODO 后续修改
		return s, true
	default:
		return "", false
	}
}
