package main

//结构体Player定义了一个玩家的基本属性和方法。结构体的currPos表示当前位置，speed表示速度
type Player struct {
	currPos   Vec2
	targetPos Vec2
	speed     float32
}

//定义玩家的移动方法。逻辑层通过这个函数告知玩家要去的目标位置，随后的移动过程由Updata()方法负责
func (p *Player) MoveTo(v Vec2) {

	p.targetPos = v
}

//使用Pos方法实现玩家currPos的属性访问封装
func (p *Player) Pos() Vec2 {
	return p.currPos
}

//判断玩家是否到达目标点。玩家每次移动的半径就是速度，因此，如果目标点的距离小于速度，表示已经非常靠近目标，可以视为到达目标
func (p *Player) IsArrived() boll {

	return p.currPos.DistanceTo(p.targetPos) < p.speed
}

func (p *Player) Update() {

	if !p.IsArrived() { //如果已经到达，则不必再更新

		//数学中，两矢量详见将获得指向被减矢量的新矢量，Sub()方法返回的新矢量使用Normailze()方法单位化。
		//最终返回的dir矢量就是移动方向
		dir := p.targetPos.Sub(p.currPos).Normalize()

		newPos := p.targetPos.Add(dir.Scale(p.speed))

		p.currPos = newPos
	}
}

//玩家的构造函数，创建一个玩家实例需要传入的速度值
func NewPlayer(spdde float32) *Player {

	return &Player{
		speed: speed,
	}
}
