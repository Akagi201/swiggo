package director_test

import (
	"fmt"
	"testing"

	"os"

	"github.com/Akagi201/swiggo/director"
)

func Compare(name string, got string, exp string) error {
	fmt.Printf("%s; Got: '%s'; Expected: '%s'\n", name, got, exp)
	if got != exp {
		return fmt.Errorf("%s returned unexpected string! Got: '%s'; Expected: '%s'\n", name, got, exp)
	}
	return nil
}

func testFooBarCpp() error {
	fb := director.NewFooBarCpp()
	defer director.DeleteFooBarCpp(fb)
	return Compare("FooBarCpp.FooBar()", fb.FooBar(), "C++ Foo, C++ Bar")
}

func testFooBarGo() error {
	fb := director.NewFooBarGo()
	defer director.DeleteFooBarGo(fb)
	return Compare("FooBarGo.FooBar()", fb.FooBar(), "Go Foo, Go Bar")
}

func TestExample(t *testing.T) {
	fmt.Println("Test output:")
	fmt.Println("------------")
	err := testFooBarCpp()
	err = testFooBarGo()
	fmt.Println("------------")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Tests failed! Last error: %s\n", err.Error())
		os.Exit(1)
	}
}
