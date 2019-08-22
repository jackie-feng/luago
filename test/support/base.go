package support

import (
	"fmt"
	"luago/api"
)

func PrintStack(ls api.LuaState) {
	top := ls.GetTop()
	for i := 1; i <= top; i++ {
		t := ls.Type(i)
		switch t {
		case api.LuaTBoolean:
			fmt.Printf("[%t]", ls.ToBoolean(i))
		case api.LuaTNumber:
			fmt.Printf("[%g]", ls.ToNumber(i))
		case api.LuaTString:
			fmt.Printf("[%q]", ls.ToString(i))
		default:
			fmt.Printf("[%s]", ls.TypeName(t))
		}
	}
	fmt.Println()
}
