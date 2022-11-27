package controllable

import (
	"sync"

	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

// Transform holds the base position and velocity of an entity. It holds several methods which can be used when
// embedding the struct.
type Transform struct {
	e        world.Entity
	mu       sync.Mutex
	vel, pos mgl64.Vec3
}

// newTransform creates a new transform to embed for the world.Entity passed.
func NewTransform(e world.Entity, pos mgl64.Vec3) Transform {
	return Transform{e: e, pos: pos}
}

// Position returns the current position of the entity.
func (t *Transform) Position() mgl64.Vec3 {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.pos
}

// Velocity returns the current velocity of the entity. The values in the Vec3 returned represent the speed on
// that axis in blocks/tick.
func (t *Transform) Velocity() mgl64.Vec3 {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.vel
}

// SetVelocity sets the velocity of the entity. The values in the Vec3 passed represent the speed on
// that axis in blocks/tick.
func (t *Transform) SetVelocity(v mgl64.Vec3) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.vel = v
}

// Rotation always returns an empty cube.Rotation.
func (t *Transform) Rotation() cube.Rotation { return cube.Rotation{} }

// World returns the world of the entity.
func (t *Transform) World() *world.World {
	w, _ := world.OfEntity(t.e)
	return w
}

// Close closes the transform and removes the associated entity from the world.
func (t *Transform) Close() error {
	w, _ := world.OfEntity(t.e)
	w.RemoveEntity(t.e)
	return nil
}
