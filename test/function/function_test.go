package function

import (
	"io/ioutil"
	"luago/state"
	"os"
	"testing"
)

func TestFunction(t *testing.T) {
	if len(os.Args) > 1 {
		data, err := ioutil.ReadFile("function_test.luac")
		if err != nil {
			panic(err)
		}
		ls := state.New(20, nil)
		ls.Load(data, os.Args[1], "b")
		ls.Call(0, 0)
	}
}
