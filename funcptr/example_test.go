package funcptr_test

import (
	"fmt"
	"testing"

	"github.com/Akagi201/swiggo/funcptr"
)

func TestExample(t *testing.T) {
	a := 37
	b := 42

	// Now call our C function with a bunch of callbacks

	fmt.Println("Trying some C callback functions")
	fmt.Println("    a        = ", a)
	fmt.Println("    b        = ", b)
	fmt.Println("    ADD(a,b) = ", funcptr.Do_op(a, b, funcptr.ADD))
	fmt.Println("    SUB(a,b) = ", funcptr.Do_op(a, b, funcptr.SUB))
	fmt.Println("    MUL(a,b) = ", funcptr.Do_op(a, b, funcptr.MUL))

	fmt.Println("Here is what the C callback function classes are called in Go")
	fmt.Println("    ADD      = ", funcptr.ADD)
	fmt.Println("    SUB      = ", funcptr.SUB)
	fmt.Println("    MUL      = ", funcptr.MUL)
}
