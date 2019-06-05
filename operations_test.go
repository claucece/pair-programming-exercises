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
