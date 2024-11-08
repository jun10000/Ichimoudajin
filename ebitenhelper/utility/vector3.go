package utility

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

func NewVector3(x float64, y float64, z float64) Vector3 {
	return Vector3{x, y, z}
}

func ZeroVector3() Vector3 {
	return NewVector3(0, 0, 0)
}

func (v Vector3) Dot(value Vector3) float64 {
	return v.X*value.X + v.Y*value.Y + v.Z*value.Z
}

func (v Vector3) Cross(value Vector3) Vector3 {
	return NewVector3(
		v.Y*value.Z-v.Z*value.Y,
		v.Z*value.X-v.X*value.Z,
		v.X*value.Y-v.Y*value.X,
	)
}
