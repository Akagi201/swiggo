package enums_test

import (
	"testing"

	"github.com/Akagi201/swiggo/enums"
	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	// Print out the value of some enums
	// *** color ***
	color := enums.RED
	assert.Equal(t, enums.Color(0), color)
	color = enums.BLUE
	assert.Equal(t, enums.Color(1), color)
	color = enums.GREEN
	assert.Equal(t, enums.Color(2), color)

	// *** Foo::speed ***
	speed := enums.FooIMPULSE
	assert.Equal(t, enums.FooSpeed(10), speed)
	speed = enums.FooWARP
	assert.Equal(t, enums.FooSpeed(20), speed)
	speed = enums.FooLUDICROUS
	assert.Equal(t, enums.FooSpeed(30), speed)

	// Testing use of enums with functions
	enums.Enum_test(enums.RED, enums.FooIMPULSE)
	enums.Enum_test(enums.BLUE, enums.FooWARP)
	enums.Enum_test(enums.GREEN, enums.FooLUDICROUS)

	// Testing use of enum with class method
	f := enums.NewFoo()

	f.Enum_test(enums.FooIMPULSE)
	f.Enum_test(enums.FooWARP)
	f.Enum_test(enums.FooLUDICROUS)
}
