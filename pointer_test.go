package easy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPinter(t *testing.T) {
	t.Parallel()

	var myIntPtr *int

	require.Zero(t, GetPointer(myIntPtr))

	someInt := 15
	myIntPtr = &someInt
	require.Equal(t, 15, GetPointer(myIntPtr))
}

func TestPointerTo(t *testing.T) {
	t.Parallel()

	myIntPtr := PointerTo(25)
	require.Equal(t, 25, *myIntPtr)
}
