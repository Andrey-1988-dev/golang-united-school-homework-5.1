package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	endPoint := s.start
	endPoint.x += int(s.a)
	endPoint.y += int(s.a)
	return endPoint
}

func (s Square) Area() uint {
	return s.a * s.a
}

func (s Square) Perimeter() uint {
	return s.a * 4
}
