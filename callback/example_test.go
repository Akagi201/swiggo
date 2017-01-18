package callback_test

import (
	"fmt"
	"testing"

	"github.com/Akagi201/swiggo/callback"
)

func TestExample(t *testing.T) {
	fmt.Println("Adding and calling a normal C++ callback")
	fmt.Println("----------------------------------------")

	caller := callback.NewCaller()
	cb := callback.NewCallback()

	caller.SetCallback(cb)
	caller.Call()
	caller.DelCallback()

	go_callback := callback.NewGoCallback()

	fmt.Println()
	fmt.Println("Adding and calling a Go callback")
	fmt.Println("--------------------------------")

	caller.SetCallback(go_callback)
	caller.Call()
	caller.DelCallback()

	callback.DeleteGoCallback(go_callback)

	fmt.Println()
	fmt.Println("Go exit")
}
