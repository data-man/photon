package photon

type Point struct {
	X int
	Y int
}

func (p *Point) SetPoint(x, y int) {
	p.X = x
	p.Y = y
}
