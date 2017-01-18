package simple_test

import (
	"fmt"

	"github.com/Akagi201/swiggo/simple"
)

func Example() {
	// Call our gcd() function
	x := 42
	y := 105
	g := simple.Gcd(x, y)
	fmt.Println("The gcd of", x, "and", y, "is", g)

	// Manipulate the Foo global variable

	// Print its current value
	fmt.Println("Foo =", simple.GetFoo())

	// Change its value
	simple.SetFoo(3.1415926)

	// See if the change took effect
	fmt.Println("Foo =", simple.GetFoo())

	// Output:
	// The gcd of 42 and 105 is 21
	// Foo = 3
	// Foo = 3.1415926
}
