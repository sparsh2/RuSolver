package colors

type Color byte

const (
	BLACK Color = iota
	WHITE
	YELLOW
	GREEN
	BLUE
	RED
	ORANGE
)

func (c Color) ToString() string {
	switch c {
	case WHITE:
		return "W"
	case YELLOW:
		return "Y"
	case GREEN:
		return "G"
	case BLUE:
		return "B"
	case RED:
		return "R"
	case ORANGE:
		return "O"
	default:
		return "-"
	}
}

func Parse(c string) Color {
	switch c {
	case "W":
		return WHITE
	case "Y":
		return YELLOW
	case "G":
		return GREEN
	case "B":
		return BLUE
	case "R":
		return RED
	case "O":
		return ORANGE
	default:
		return BLACK
	}
}
