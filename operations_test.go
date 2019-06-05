package operations

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type OperationsSuite struct{}

var _ = Suite(&OperationsSuite{})

func (s *OperationsSuite) Test_Example(c *C) {
	ok := true
	c.Assert(ok, Equals, true)
}
