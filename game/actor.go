package game

import "github.com/wmiller848/Karma/renderer"

type Actor struct {
	renderer renderer.GenericRenderer
	position [4]float32
}

func CreateActor(config string, r renderer.GenericRenderer) *Actor {

	var vertices = []float32{
		//  X, Y, Z, U, V
		1.0, -1.0, 0.0, 1.0, 0.0,
		-1.0, 1.0, 0.0, 0.0, 1.0,
		-1.0, -1.0, 0.0, 0.0, 0.0,
		1.0, -1.0, 0.0, 1.0, 0.0,
		-1.0, 1.0, 0.0, 0.0, 1.0,
		1.0, 1.0, 0.0, 1.0, 1.0,
	}

	r.AddMesh(vertices, []string{"barb.png"})
	r.UpdateMeshPos(0, [3]float32{0, 0, 0})

	return &Actor{
		renderer: r,
		position: [4]float32{0, 0, 0},
	}
}

func (a *Actor) Move(change [3]float32) {
	a.position[0] += change[0]
	a.position[1] += change[1]
	a.position[2] += change[2]
}

func (a *Actor) DetectCollision() {

}

func (a *Actor) Tick() {
	a.renderer.UpdateMeshPos(0, [3]float32{a.position[0], a.position[1], a.position[2]})
	a.Move([3]float32{0.1, 0, 0})
}
