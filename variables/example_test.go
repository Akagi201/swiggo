package variables_test

import (
	"fmt"
	"testing"

	"github.com/Akagi201/swiggo/variables"
	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	// Try to set the values of some global variables

	variables.SetIvar(42)
	variables.SetSvar(-31000)
	variables.SetLvar(65537)
	variables.SetUivar(123456)
	variables.SetUsvar(61000)
	variables.SetUlvar(654321)
	variables.SetScvar(-13)
	variables.SetUcvar(251)
	variables.SetCvar('S')
	variables.SetFvar(3.14159)
	variables.SetDvar(2.1828)
	variables.SetStrvar("Hello World")
	variables.SetIptrvar(variables.New_int(37))
	variables.SetPtptr(variables.New_Point(37, 42))
	variables.SetName("Bill")

	// Now print out the values of the variables
	// Variables (values printed from Go)

	assert.Equal(t, 42, variables.GetIvar())
	assert.Equal(t, int16(-31000), variables.GetSvar())
	assert.Equal(t, int64(65537), variables.GetLvar())
	assert.Equal(t, uint(123456), variables.GetUivar())
	assert.Equal(t, uint16(61000), variables.GetUsvar())
	assert.Equal(t, uint64(654321), variables.GetUlvar())
	assert.Equal(t, int8(-13), variables.GetScvar())
	assert.Equal(t, uint8(251), variables.GetUcvar())
	assert.Equal(t, string("S"), fmt.Sprintf("%c", variables.GetCvar()))
	assert.Equal(t, float32(3.14159), variables.GetFvar())
	//assert.Equal(t, 2.1828, variables.GetDvar())

	assert.Equal(t, "Hello World", variables.GetStrvar())
	assert.Equal(t, variables.New_int(37), variables.GetIptrvar())
	assert.Equal(t, "Bill", variables.GetName())
	assert.Equal(t, "Goodbye", variables.GetCstrvar())
	//assert.Equal(t, variables.Point_print(variables.GetPtptr()), variables.GetPtptr())
	//assert.Equal(t, variables.Point_print(variables.GetPt()), variables.GetPt())

	// Variables (values printed from C)

	variables.Print_vars()

	// This line would not compile: since status is marked with
	// %immutable, there is no SetStatus function.
	// fmt.Println("\nNow I'm going to try and modify some read only variables")
	// variables.SetStatus(0)

	// I'm going to try and update a structure variable.

	variables.SetPt(variables.GetPtptr())

	//assert.Equal(t, variables.Point_print(variables.GetPtptr()), variables.Pt_print())
}
