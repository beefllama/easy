package easy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConvertNullable(t *testing.T) {
	t.Parallel()

	type customIntegerType struct {
		a int
	}

	var myIntPtr *int

	require.Equal(t, (*customIntegerType)(nil), ConvertNullable(myIntPtr, func(myInt int) customIntegerType { return customIntegerType{a: myInt} }))

	someInt := 5
	myIntPtr = &someInt
	require.Equal(t, &customIntegerType{
		a: someInt,
	}, ConvertNullable(myIntPtr, func(myInt int) customIntegerType { return customIntegerType{a: myInt} }))
}

func TestTernary(t *testing.T) {
	t.Parallel()

	a, b := 11, 23

	require.Equal(t, a, Ternary(true, a, b))
	require.Equal(t, b, Ternary(false, a, b))
}
