package pointer_test

import (
	"testing"

	"fmt"

	"github.com/Akagi201/swiggo/pointer"
)

func TestExample(t *testing.T) {
	// First create some objects using the pointer library.
	fmt.Println("Testing the pointer library")
	a := pointer.New_intp()
	b := pointer.New_intp()
	c := pointer.New_intp()
	pointer.Intp_assign(a, 37)
	pointer.Intp_assign(b, 42)

	fmt.Println("     a =", a)
	fmt.Println("     b =", b)
	fmt.Println("     c =", c)

	// Call the add() function with some pointers
	pointer.Add(a, b, c)

	// Now get the result
	res := pointer.Intp_value(c)
	fmt.Println("     37 + 42 =", res)

	// Clean up the pointers
	pointer.Delete_intp(a)
	pointer.Delete_intp(b)
	pointer.Delete_intp(c)

	// Now try the typemap library
	// Now it is no longer necessary to manufacture pointers.
	// Instead we use a single element slice which in Go is modifiable.

	fmt.Println("Trying the typemap library")
	r := []int{0}
	pointer.Sub(37, 42, r)
	fmt.Println("     37 - 42 = ", r[0])

	// Now try the version with return value

	fmt.Println("Testing return value")
	q := pointer.Divide(42, 37, r)
	fmt.Println("     42/37 = ", q, " remainder ", r[0])
}
