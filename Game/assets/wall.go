package assets

import (
	"github.com/Wordluc/TheEngine/core"
	"github.com/Wordluc/TheEngine/core/base"
)

type Wall struct {
	o    base.Object
	node base.Node
	Opt  struct {
		Pos  base.Vec[float32]
		Size base.Vec[float32]
	}
}

func (c *Wall) Start() error {
	c.o = new(core.NewRectangle(c.Opt.Size.X, c.Opt.Size.Y))
	c.o.MoveTo(c.Opt.Pos)
	c.o.SetModifier(new(base.NewRigidBody(true, true, 0)))
	c.node = base.NewNode(base.Vec[float32]{})
	c.node.AddObject(c.o)
	return nil
}

func (c *Wall) GetEntity() base.Node {
	return c.node
}

func (c *Wall) Update(dt float32) error {
	return nil
}
