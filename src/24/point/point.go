package point

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point { // Конструктор
	return &Point{x: x, y: y}
}

func (p *Point) GetX() float64 { // Getter X
	return p.x
}

func (p *Point) GetY() float64 { // Getter X
	return p.y
}

func (p *Point) SetX(val float64) { // Setter x
	p.x = val
}

func (p *Point) SetY(val float64) { // Setter y
	p.y = val
}
