package side

const (
	Left = -1 + iota
	On
	Right
)

// Side type
type Side struct {
	s int
}

// NewSide - Create New Side
func NewSide() *Side {
	return &Side{On}
}

// Value - get value
func (side *Side) Value() int {
	return side.s
}

// IsLeft - is left
func (side *Side) IsLeft() bool {
	return side.s == Left
}

// AsLeft - as left
func (side *Side) AsLeft() *Side {
	side.s = Left
	return side
}

// IsOn - Is on
func (side *Side) IsOn() bool {
	return side.s == On
}

// AsOn - As on
func (side *Side) AsOn() *Side {
	side.s = On
	return side
}

// IsRight - Is right
func (side *Side) IsRight() bool {
	return side.s == Right
}

// AsRight - As right
func (side *Side) AsRight() *Side {
	side.s = Right
	return side
}

// IsOnOrLeft - is on or left
func (side *Side) IsOnOrLeft() bool {
	return side.IsOn() || side.IsLeft()
}

// IsOnOrRight - is on or right
func (side *Side) IsOnOrRight() bool {
	return side.IsOn() || side.IsRight()
}

// IsSameSide - is on the same side
func (side *Side) IsSameSide(other *Side) bool {
	return side.s == other.s
}
