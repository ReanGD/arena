package la

type Rect struct {
	X int32
	Y int32
	W int32
	H int32
}

func NewRect(x, y, w, h int32) Rect {
	return Rect{x, y, w, h}
}

func (r Rect) CenterVec2() Vec2 {
	return NewVec2(float64(r.X)+float64(r.W)/2.0, float64(r.Y)+float64(r.H)/2.0)
}

func (r Rect) Center() (int32, int32) {
	return r.X + r.W/2, r.Y + r.H/2
}

func (r Rect) MinX() int32 {
	return r.X
}

func (r Rect) MinY() int32 {
	return r.Y
}

func (r Rect) MaxX() int32 {
	return r.X + r.W
}

func (r Rect) MaxY() int32 {
	return r.Y + r.H
}
