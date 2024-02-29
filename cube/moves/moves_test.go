package moves

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvertMoves(t *testing.T) {
	movesToTest := []Move{
		RIGHT,
		UP,
		RIGHT_INVERTED,
		UP_INVERTED,
	}
	expected := []Move{
		UP,
		RIGHT,
		UP_INVERTED,
		RIGHT_INVERTED,
	}

	actual := InvertMoves(movesToTest)

	assert.Equal(t, len(expected), len(actual))

	for i := range actual {
		assert.Equal(t, expected[i], actual[i])
	}
}
