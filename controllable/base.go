package controllable

import (
	"errors"
	"image/png"
	"os"

	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/inventory"
	"github.com/df-mc/dragonfly/server/player/form"
	"github.com/df-mc/dragonfly/server/player/skin"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/google/uuid"
	"golang.org/x/text/language"
)

type ControllableBase struct {
	Transform

	uUID uuid.UUID
}

func (s *ControllableBase) Message(...any) {

}

func (s *ControllableBase) HeldItems() (mainHand, offHand item.Stack) {
	return item.NewStack(item.Diamond{}, 0), item.NewStack(item.Diamond{}, 0)
}
func (s *ControllableBase) UsingItem() bool                                                  { return false }
func (s *ControllableBase) ReleaseItem()                                                     {}
func (s *ControllableBase) UseItem()                                                         {}
func (s *ControllableBase) SendForm(form.Form)                                               {}
func (s *ControllableBase) SendCommandOutput(o *cmd.Output)                                  {}
func (s *ControllableBase) Locale() language.Tag                                             { return language.BritishEnglish }
func (s *ControllableBase) SetHeldItems(right, left item.Stack)                              {}
func (s *ControllableBase) Move(deltaPos mgl64.Vec3, deltaYaw, deltaPitch float64)           {}
func (s *ControllableBase) Speed() float64                                                   { return 1 }
func (s *ControllableBase) Chat(msg ...any)                                                  {}
func (s *ControllableBase) ExecuteCommand(commandLine string)                                {}
func (s *ControllableBase) GameMode() world.GameMode                                         { return world.GameModeCreative }
func (s *ControllableBase) SetGameMode(mode world.GameMode)                                  {}
func (s *ControllableBase) Effects() []effect.Effect                                         { return []effect.Effect{} }
func (s *ControllableBase) UseItemOnBlock(pos cube.Pos, face cube.Face, clickPos mgl64.Vec3) {}
func (s *ControllableBase) UseItemOnEntity(e world.Entity) bool                              { return false }
func (s *ControllableBase) BreakBlock(pos cube.Pos)                                          {}
func (s *ControllableBase) PickBlock(pos cube.Pos)                                           {}
func (s *ControllableBase) AttackEntity(e world.Entity) bool                                 { return false }
func (s *ControllableBase) Drop(st item.Stack) (n int)                                       { return 1 }
func (s *ControllableBase) SwingArm()                                                        {}
func (s *ControllableBase) PunchAir()                                                        {}
func (s *ControllableBase) ExperienceLevel() int                                             { return 1 }
func (s *ControllableBase) SetExperienceLevel(level int)                                     {}
func (s *ControllableBase) EnchantmentSeed() int64                                           { return 1 }
func (s *ControllableBase) ResetEnchantmentSeed()                                            {}
func (s *ControllableBase) Respawn()                                                         {}
func (s *ControllableBase) Dead() bool                                                       { return false }
func (s *ControllableBase) StartSneaking()                                                   {}
func (s *ControllableBase) Sneaking() bool                                                   { return false }
func (s *ControllableBase) StopSneaking()                                                    {}
func (s *ControllableBase) StartSprinting()                                                  {}
func (s *ControllableBase) Sprinting() bool                                                  { return false }
func (s *ControllableBase) StopSprinting()                                                   {}
func (s *ControllableBase) StartSwimming()                                                   {}
func (s *ControllableBase) Swimming() bool                                                   { return false }
func (s *ControllableBase) StopSwimming()                                                    {}
func (s *ControllableBase) StartFlying()                                                     {}
func (s *ControllableBase) Flying() bool                                                     { return false }
func (s *ControllableBase) StopFlying()                                                      {}
func (s *ControllableBase) StartGliding()                                                    {}
func (s *ControllableBase) Gliding() bool                                                    { return false }
func (s *ControllableBase) StopGliding()                                                     {}
func (s *ControllableBase) Jump()                                                            {}
func (s *ControllableBase) StartBreaking(pos cube.Pos, face cube.Face)                       {}
func (s *ControllableBase) ContinueBreaking(face cube.Face)                                  {}
func (s *ControllableBase) FinishBreaking()                                                  {}
func (s *ControllableBase) AbortBreaking()                                                   {}
func (s *ControllableBase) Exhaust(points float64)                                           {}
func (s *ControllableBase) EditSign(pos cube.Pos, text string) error                         { return nil }
func (s *ControllableBase) EnderChestInventory() *inventory.Inventory                        { return inventory.New(1, nil) }
func (s *ControllableBase) XUID() string                                                     { return "" }
func (s *ControllableBase) SetSkin(skin.Skin)                                                {}

func (s *ControllableBase) UUID() uuid.UUID {
	return s.uUID
}

func (c *ControllableBase) Skin() skin.Skin {
	panic(errors.New("'Skin() skin.Skin' must be implemented on custom controllable"))
}

func MapVec3(x map[string]any, k string) mgl64.Vec3 {
	if i, ok := x[k].([]any); ok {
		if len(i) != 3 {
			return mgl64.Vec3{}
		}
		var v mgl64.Vec3
		for index, f := range i {
			f32, _ := f.(float32)
			v[index] = float64(f32)
		}
		return v
	} else if i, ok := x[k].([]float32); ok {
		if len(i) != 3 {
			return mgl64.Vec3{}
		}
		return mgl64.Vec3{float64(i[0]), float64(i[1]), float64(i[2])}
	}
	return mgl64.Vec3{}
}

func Vec3ToFloat32Slice(x mgl64.Vec3) []float32 {
	return []float32{float32(x[0]), float32(x[1]), float32(x[2])}
}

func LoadSkin(width int, height int, pathToImage string, pathToModel string, modelName string) skin.Skin {
	s := skin.New(width, height)
	s.ModelConfig = skin.ModelConfig{Default: modelName}
	if file, err := os.ReadFile(pathToModel); err == nil {
		s.Model = file
	} else {
		panic(err)
	}

	if file, err := os.Open(pathToImage); err == nil {
		if i, err := png.Decode(file); err == nil {
			data := make([]byte, 0, i.Bounds().Dx()*i.Bounds().Dy()*4)
			for currentX := i.Bounds().Min.X; currentX < i.Bounds().Max.X; currentX++ {
				for currentY := i.Bounds().Min.Y; currentY < i.Bounds().Max.Y; currentY++ {
					r, g, b, a := i.At(currentX, currentY).RGBA()
					data = append(data, byte(r>>8), byte(g>>8), byte(b>>8), byte(a>>8))
				}
			}
			s.Pix = data
			return s

		}
	} else {
		panic(err)
	}
	return skin.Skin{}
}
