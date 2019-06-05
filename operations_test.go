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
