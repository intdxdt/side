package side

const (
    Left = -1 + iota
    On
    Right
)

//type Side of zero
type Side struct {
    s int
}

//New Side
func NewSide() *Side {
    return &Side{On}
}

//get value
func (self *Side) Value() int {
    return self.s
}

//is left
func (side *Side) IsLeft() bool {
    return side.s == Left
}

//as left
func (side *Side) AsLeft() *Side {
    side.s = Left
    return side
}

//Is on
func (side *Side) IsOn() bool {
    return side.s == On
}

//As on
func (side *Side) AsOn() *Side {
    side.s = On
    return side
}

//Is right
func (side *Side) IsRight() bool {
    return side.s == Right
}

//As right
func (side *Side) AsRight() *Side {
    side.s = Right
    return side
}

//is on or left
func (side *Side) IsOnOrLeft() bool {
    return side.IsOn() || side.IsLeft()
}

//is on or right
func (side *Side) IsOnOrRight() bool {
    return side.IsOn() || side.IsRight()
}

