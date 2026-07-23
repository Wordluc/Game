package main

import (
	"Game/assets"

	"github.com/Wordluc/TheEngine/core"
	_ "github.com/Wordluc/TheEngine/core"
	"github.com/Wordluc/TheEngine/core/base"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(500, 500, "")
	engine := core.NewEntityEngine()
	c := assets.Character{Opt: struct {
		Pos  base.Vec[float32]
		Size base.Vec[float32]
		Mass float32
	}{Size: base.Vec[float32]{X: 10, Y: 10}, Pos: base.NewVec[float32](20, 450), Mass: 50}}
	box := assets.EmptyBox{
		Opt: struct {
			Pos   base.Vec[float32]
			Size  base.Vec[float32]
			Thick float32
		}{
			Size:  base.NewVec[float32](500, 500),
			Thick: 5,
		},
	}
	engine.AppendEntity(&c)
	engine.AppendEntity(&box)
	engine.RunEntitiesStarts()
	engine.AccelerationToApply = append(engine.AccelerationToApply, base.Vec[float32]{Y: 9})
	for {
		if rl.WindowShouldClose() {
			return
		}
		rl.ClearBackground(rl.White)
		rl.BeginDrawing()

		engine.RunEntitiesUpdates()

		rl.EndDrawing()
	}
}
