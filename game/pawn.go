package game

import "github.com/wmiller848/Karma/renderer"

type Pawn struct {
	actor        *Actor
	acceleration float32
	velocity     [3]float32
}

func CreatePawn(config string, r renderer.GenericRenderer) *Pawn {
	return &Pawn{
		actor:        CreateActor(config, r),
		acceleration: 0,
		velocity:     [3]float32{0, 0, 0},
	}
}

func (a *Pawn) MoveLeft(force float32) {
	a.acceleration += force
	velocity := [3]float32{0, 0, 0}
	velocity[0] = -1 * a.acceleration
	velocity[1] = 0
	velocity[2] = 0
	a.actor.Move(velocity)
}

func (a *Pawn) MoveRight(force float32) {
	a.acceleration += force
	velocity := [3]float32{0, 0, 0}
	velocity[0] += a.acceleration
	velocity[1] += 0
	velocity[2] += 0
}

func (a *Pawn) MoveUp(force float32) {
	a.acceleration += force
	velocity := [3]float32{0, 0, 0}
	velocity[0] += 0
	velocity[1] += a.acceleration
	velocity[2] += 0
}

func (a *Pawn) MoveDown(force float32) {
	a.acceleration += force
	velocity := [3]float32{0, 0, 0}
	velocity[0] += 0
	velocity[1] += -1 * a.acceleration
	velocity[2] += 0
}
