package state

import "luago/api"

type luaValue interface{}
type LuaType = int

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

func convertToBoolean(val luaValue) bool {
	switch x := val.(type) {
	case nil:
		return false
	case bool:
		return x
	default:
		return true
	}
}
