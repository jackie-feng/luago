package table

import (
	"io/ioutil"
	"luago/state"
	"testing"
)

func TestTable(t *testing.T) {
	data, err := ioutil.ReadFile("table_test.luac")
	if err != nil {
		panic(err)
	}
	luaMain(data)
}

func luaMain(chunk []byte) {
	ls := state.New(20, nil)
	ls.Load(chunk, "main", "b")
	ls.Call(0, 0)
}
