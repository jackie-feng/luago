package state

import "luago/api"

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
