基本实例化格式：
var ins T 

type Point struct {
	
	X int
	Y int
	
}

var p Point
p.X = 10
p.Y = 20

使用new的格式如下：
ins := new(T)

type Player struct{
	
	Name string
	HealthPoint int
	MagicPoint int
	
}

tank := new(Player)
tank.Name = "Canon"
tank.HealhPoint = 300

