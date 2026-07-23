package assets

import (
	"github.com/Wordluc/TheEngine/core"
	"github.com/Wordluc/TheEngine/core/base"
)

type EmptyBox struct {
	node base.Node
	Opt  struct {
		Pos   base.Vec[float32]
		Size  base.Vec[float32]
		Thick float32
	}
}

func (c *EmptyBox) Start() error {
	c.node = base.NewNode(base.Vec[float32]{})
	//UPPER
	up := new(core.NewRectangle(c.Opt.Size.X, c.Opt.Thick))
	up.SetModifier(new(base.NewRigidBody(true, true, 0)))
	//BOTTOM
	down := new(core.NewRectangle(c.Opt.Size.X, c.Opt.Thick))
	down.MoveTo(base.Vec[float32]{Y: c.Opt.Size.Y - c.Opt.Thick})
	down.SetModifier(new(base.NewRigidBody(true, true, 0)))
	//RIGHT
	right := new(core.NewRectangle(c.Opt.Thick, c.Opt.Size.Y))
	right.MoveTo(base.Vec[float32]{X: c.Opt.Size.X - c.Opt.Thick})
	right.SetModifier(new(base.NewRigidBody(true, true, 0)))
	//LEFT
	left := new(core.NewRectangle(c.Opt.Thick+10, c.Opt.Size.Y))
	left.SetModifier(new(base.NewRigidBody(true, true, 0)))
	//creating node
	c.node.AddObject(up)
	c.node.AddObject(down)
	c.node.AddObject(right)
	c.node.AddObject(left)
	return nil
}
func (c *EmptyBox) GetEntity() base.Node {
	return c.node
}

func (c *EmptyBox) Update(dt float32) error {
	return nil
}
