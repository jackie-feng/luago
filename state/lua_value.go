package state

import "luago/api"

type luaValue interface{}

func typeOf(val luaValue) LuaType {
	switch val.(type) {
	case nil:
		return api.LuaTNil
	case bool:
		return api.LuaTBoolean
	case int64:
		return api.LuaTNumber
	case float64:
		return api.LuaTNumber
	case string:
		return api.LuaTString
	default:
		panic("todo!")
	}
}
