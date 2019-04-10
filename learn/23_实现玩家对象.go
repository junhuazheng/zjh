package main

type Player struct {
	cuP   Vec2    //当前位置
	taP   Vec2    //目标位置
	speed float32 //移动速度
}

//移动到某个点就是设置目标位置
func (p *Player) MoveTo(v Vec2) {
	p.taP = v
}

//获取当前位置
func (p *Player) Pos() Vec2 {
	return p.cuP
}

//是否到达
func (p *Player) IsAr() bool {
	return p.cuP.Dis(p.taP) < p.speed
}

func (p *Player) Update() {
	if !p.IsAr() {
		dir := p.taP.Sub(p.cuP).Nor()

		newP := p.cuP.Add(dir.Scale(p.speed))

		p.cuP = newP
	}
}

func NewPlayer(speed float32) *Player {
	return &Player{
		speed: speed,
	}
}
