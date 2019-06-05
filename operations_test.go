package operations

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type OperationsSuite struct{}

var _ = Suite(&OperationsSuite{})

func (s *OperationsSuite) Test_Add(c *C) {
	var x, y, z number
	x = 5
	y = 3
	z = x.Add(y)
	c.Assert(z, Equals, number(8))
}

func (s *OperationsSuite) Test_Substract(c *C) {
	a := 8
	b := 9
	d := substraction(a, b)
	c.Assert(a-b, Equals, d)
}

func (s *OperationsSuite) Test_Substr(c *C) {
	var a, b, d number
	a = 8
	b = 9
	d = a.substr(b)
	c.Assert(a-b, Equals, d)
}

func (s *OperationsSuite) Test_Multiplication(c *C) {
	var x, y, z number
	x = 2
	y = 8
	z = x.Multiplication(y)
	c.Assert(z, Equals, number(16))
}

func (s *OperationsSuite) Test_Double_Multiplication(c *C) {
	var x, y, z number
	x = 2
	y = 8
	z = x.Multiplication(y).Multiplication(10)
	c.Assert(z, Equals, number(160))
}

func (s *OperationsSuite) Test_Divide_Int(c *C) {
	var divisor number
	var divident number
	var expectedResult float32
	var actualResult float32

	divisor = 4
	divident = 2
	expectedResult = 2

	actualResult, s = divisor.Divide(divident)

	c.Assert(s, Equals, expectedResult)
}

func (s *OperationsSuite) Test_Divide_Float(c *C) {
	var divisor number
	var divident number
	var expectedResult float32
	var actualResult float32

	divisor = 5
	divident = 2
	expectedResult = 2.5

	actualResult = divisor.Divide(divident)

	c.Assert(actualResult, Equals, expectedResult)
}

func (s *OperationsSuite) Test_Divide_Zero(c *C) {
	var divisor number
	var divident number
	var expectedResult string
	var actual_result string

	divisor = 5
	divident = 0
	expectedResult = "NaN"

	actualResult = divisor.Divide(divident)

	c.Assert(actualResult, Equals, expectedResult)
}
