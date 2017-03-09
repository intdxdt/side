package side

import (
	"testing"
	"github.com/franela/goblin"
)

func TestSideOf(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Side", func() {
		g.It("sidedness", func() {
			var s = NewSide()
			g.Assert(s.AsOn().IsOn()).IsTrue()
			g.Assert(s.Value()).Equal(0)

			g.Assert(s.AsLeft().IsLeft()).IsTrue()
			g.Assert(s.Value()).Equal(-1)
			g.Assert(s.AsLeft().IsOnOrLeft()).IsTrue()
			g.Assert(s.IsOn()).IsFalse()

			g.Assert(s.AsOn().IsOn()).IsTrue()
			g.Assert(s.IsOnOrRight()).IsTrue()
			g.Assert(s.IsOnOrLeft()).IsTrue()

			g.Assert(s.AsRight().IsRight()).IsTrue()
			g.Assert(s.Value()).Equal(1)
			g.Assert(s.AsRight().IsOnOrRight()).IsTrue()
			g.Assert(s.IsOn()).IsFalse()
		})
	})

}
