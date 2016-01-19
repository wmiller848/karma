package game

import (
	"github.com/wmiller848/Karma/game/cognition"
	"github.com/wmiller848/Karma/renderer"
)

type GenericEntity interface {
	Tick()
}

type Game struct {
	renderer renderer.GenericRenderer
	paused   bool
	Players  []cognition.Player
	Entities []GenericEntity
}

func CreateGame(r renderer.GenericRenderer) *Game {
	r.UpdateCameraPos([3]float32{0, 0, -10})
	return &Game{
		renderer: r,
		paused:   false,
		Players:  []cognition.Player{},
		Entities: []GenericEntity{},
	}
}

func (g *Game) LoadLevel() {
	// var vertices = []float32{
	// 	//  X, Y, Z, U, V
	// 	1.0, -1.0, 0.0, 1.0, 0.0,
	// 	-1.0, 1.0, 0.0, 0.0, 1.0,
	// 	-1.0, -1.0, 0.0, 0.0, 0.0,
	// 	1.0, -1.0, 0.0, 1.0, 0.0,
	// 	-1.0, 1.0, 0.0, 0.0, 1.0,
	// 	1.0, 1.0, 0.0, 1.0, 1.0,
	// }
	//
	// g.renderer.AddMesh(vertices, []string{"barb.png"})
	// g.renderer.UpdateMeshPos(0, [3]float32{1, 0, 0})
	//
	// g.renderer.AddMesh(vertices, []string{"barb.png"})
	// g.renderer.UpdateMeshPos(1, [3]float32{-1, 0, 0})
	g.Entities = append(g.Entities, CreateActor("", g.renderer))
}

func (g *Game) SaveLevel() {

}

func (g *Game) Play() {
	for !g.renderer.ShouldDie() && !g.paused {
		//
		// TODO : rate limit rendering to 60htz
		g.renderer.Render()
		for _, entity := range g.Entities {
			entity.Tick()
		}
	}
}

func (g *Game) Pause() {
	g.paused = !g.paused
}
