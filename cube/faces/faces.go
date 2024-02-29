package faces

type Face byte

const (
	FRONT Face = iota
	BACK
	LEFT
	RIGHT
	TOP
	BOTTOM
)

func (f Face) ToString() string {
	switch f {
	case FRONT:
		return "Front"
	case BACK:
		return "Back"
	case LEFT:
		return "Left"
	case RIGHT:
		return "Right"
	case TOP:
		return "Top"
	case BOTTOM:
		return "Bottom"
	default:
		return ""
	}
}
