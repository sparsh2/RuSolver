package moves

type Move byte

const (
	LEFT Move = iota
	RIGHT
	FRONT
	BACK
	UP
	DOWN
	X
	Y
	Z
	M
	E
	S

	LEFT_INVERTED
	RIGHT_INVERTED
	FRONT_INVERTED
	BACK_INVERTED
	UP_INVERTED
	DOWN_INVERTED
	X_INVERTED
	Y_INVERTED
	Z_INVERTED
	M_INVERTED
	E_INVERTED
	S_INVERTED

	LEFT2
	RIGHT2
	FRONT2
	BACK2
	UP2
	DOWN2
	X2
	Y2
	Z2
	M2
	E2
	S2

	INVALID
)
