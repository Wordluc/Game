package assets

import (
	"github.com/Wordluc/TheEngine/core"
	"github.com/Wordluc/TheEngine/core/base"
	"github.com/Wordluc/TheEngine/core/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Character struct {
	node base.Node
	o    base.Object
	rg   *base.RigidBody
	Opt  struct {
		Pos  base.Vec[float32]
		Size base.Vec[float32]
		Mass float32
	}
	leftjump   int
	hangOnWall bool
}

func (c *Character) Start() error {
	c.o = new(core.NewRectangle(c.Opt.Size.X, c.Opt.Size.Y))
	c.o.MoveTo(c.Opt.Pos)
	c.rg = new(base.NewRigidBody(true, false, c.Opt.Mass))
	c.rg.Id = "my"
	c.o.SetModifier(c.rg)
	c.node = base.NewNode(base.Vec[float32]{})
	c.node.AddObject(c.o)
	return nil
}

func (c *Character) GetEntity() base.Node {
	return c.node
}

func (c *Character) Update(dt float32) error {
	isTouchingDown := c.rg.Collision.CheckIf(func(cd base.CollisionDetail) bool {
		return cd.Y < 0
	})
	isTouchingWall := c.rg.Collision.CheckIf(func(cd base.CollisionDetail) bool {
		return cd.X != 0
	})
	v := utils.GetVecForKeyboard(300)
	v.Y = 0
	if isTouchingWall {
		if c.rg.IsStatic && rl.IsKeyPressed(rl.KeySpace) {
			c.rg.IsStatic = false
		} else if !c.rg.IsStatic && rl.IsKeyPressed(rl.KeySpace) {
			c.rg.IsStatic = true
			c.hangOnWall = true
			c.rg.GetVelocity().SetXY(base.NULL_VEC32)
			c.rg.GetForce().SetXY(base.NULL_VEC32)
		}
	}
	if c.hangOnWall || isTouchingDown {
		c.leftjump = 2
	}
	if (isTouchingDown || c.leftjump > 0) || c.hangOnWall {
		if rl.IsKeyPressed(rl.KeyW) {
			c.leftjump--
			if c.hangOnWall {
				c.rg.IsStatic = false
				c.hangOnWall = false
				c.rg.ApplyImpulse(base.Vec[float32]{Y: -400})
			} else {
				c.rg.ApplyImpulse(base.Vec[float32]{Y: -300})
			}
		}
	}
	if !c.rg.IsStatic {
		c.rg.ApplyImpulse(v)
	}
	return nil
}
