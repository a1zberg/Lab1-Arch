package implementation

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestPrefToPost(c *C) {
	c.Assert(PrefToPost("+5*-423"), HasLen, 7)
	c.Assert(PrefToPost("-53"), Equals, "53-")
	c.Assert(PrefToPost("*-A/BC-/AKL"), Equals, "ABC/-AK/L-*")
	c.Assert(PrefToPost(""), Equals, "The string is empty")
	c.Assert(PrefToPost("Hello!"), Equals, "Unacceptable symbols!")
}
func ExamplePrefToPost() {
	fmt.Println(PrefToPost("*12"))
	//Output:
	//12*
}
