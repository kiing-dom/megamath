package entities

import "github.com/kiing-dom/megamath/common"

type Player struct {
	Position common.Vec2
	Velocity common.Vec2
	Char     rune
}

func (p *Player) MoveUp() {
	p.Position.Y -= p.Velocity.Y
}

func (p *Player) MoveDown() {
	p.Position.Y += p.Velocity.Y
}

func (p *Player) MoveLeft() {
	p.Position.X -= p.Velocity.X
}

func (p *Player) MoveRight() {
	p.Position.X += p.Velocity.X
}
