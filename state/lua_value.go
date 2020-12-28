package state

import (
	"fmt"
	"luago/api"
	"luago/number"
)

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
	case *luaTable:
		return api.LuaTTable
	case *closure:
		return api.LuaTFunction
	default:
		panic("todo!")
	}
}

func stringOf(val luaValue) string {
	switch val.(type) {
	case nil:
		return "[nil]"
	case bool:
		return fmt.Sprintf("[%t]", val.(bool))
	case int64:
		return fmt.Sprintf("[%g]", float64(val.(int64)))
	case string:
		return fmt.Sprintf("[%q]", val.(string))
	case float64:
		return fmt.Sprintf("[%g]", val.(float64))
	case *luaTable:
		return "[luaTable]"
	case *closure:
		return "[closure]"
	default:
		return fmt.Sprintf("[%s]", val)
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

func convertToFloat(val luaValue) (float64, bool) {
	switch x := val.(type) {
	case float64:
		return x, true
	case int64:
		return float64(x), true
	case string:
		return number.ParseFloat(x)
	default:
		return 0, false
	}
}
func convertToInteger(val luaValue) (int64, bool) {
	switch x := val.(type) {
	case int64:
		return x, true
	case float64:
		return number.FloatToInteger(x)
	case string:
		return _stringToInteger(x)
	default:
		return 0, false
	}
}

func _stringToInteger(s string) (int64, bool) {
	if i, ok := number.ParseInteger(s); ok {
		return i, true
	}

	if f, ok := number.ParseFloat(s); ok {
		return number.FloatToInteger(f)
	}
	return 0, false
}
