package entities

type Player struct {
	Position Vec2
	Velocity Vec2
	Char     rune
}

type Vec2 struct {
	X int
	Y int
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
