package easy

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	t.Parallel()

	type st1 struct {
		a int64
		b float64
	}

	type st2 struct {
		a int64
		s string
	}

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		var src []st1
		converted := Map(src, func(e st1) st2 {
			return st2{
				a: e.a,
				s: strconv.FormatInt(e.a, 10) + " " + strconv.FormatFloat(e.b, 'f', 2, 64),
			}
		})

		require.Equal(t, []st2{}, converted)
	})

	t.Run("convert structs", func(t *testing.T) {
		t.Parallel()

		src := []st1{{a: 15, b: 25.0}, {a: 25, b: 35.0}}
		converted := Map(src, func(e st1) st2 {
			return st2{
				a: e.a,
				s: strconv.FormatInt(e.a, 10) + " " + strconv.FormatFloat(e.b, 'f', 2, 64),
			}
		})
		require.Equal(t, []st2{
			{
				a: 15,
				s: "15 25.00",
			},
			{
				a: 25,
				s: "25 35.00",
			},
		}, converted)
	})
}

func TestFilter(t *testing.T) {
	t.Parallel()

	numbers := []int{1, 208, 54, 100, 20932, 533, 8, 211}

	numbersOverOneHundred := Filter(numbers, func(n int) bool { return n > 100 })

	require.Equal(t, []int{208, 20932, 533, 211}, numbersOverOneHundred)
}

func TestMapFilter(t *testing.T) {
	t.Parallel()

	numbers := []int{1, 208, 54, 100, 20932, 533, 8, 211}

	numbersOverOneHundredStrings := MapFilter(numbers, func(n int) (string, bool) {
		if n <= 100 {
			return "", false
		}

		return strconv.FormatInt(int64(n), 10), true
	})

	require.Equal(t, []string{"208", "20932", "533", "211"}, numbersOverOneHundredStrings)
}

func TestContains(t *testing.T) {
	t.Parallel()

	numbers := []int{1, 2, 3}

	require.True(t, Contains(numbers, 2))
	require.False(t, Contains(numbers, 5))
}

func TestContainsFunc(t *testing.T) {
	t.Parallel()

	type st struct {
		a int
	}

	s := []st{
		{a: 15},
		{a: 25},
	}

	require.True(t, ContainsFunc(s, func(e st) bool { return e.a == 25 }))
	require.False(t, ContainsFunc(s, func(e st) bool { return e.a == 55 }))
}

func TestBatch(t *testing.T) {
	t.Parallel()

	t.Run("empty slice", func(t *testing.T) {
		t.Parallel()

		require.Equal(t, [][]string{}, Batch([]string{}, 10))
	})

	t.Run("batch size is zero", func(t *testing.T) {
		t.Parallel()

		s := []string{"s1", "s2"}

		require.Equal(t, [][]string{s}, Batch(s, 0))
	})

	t.Run("one batch", func(t *testing.T) {
		t.Parallel()

		s := make([]int, 0, 20)

		for i := range 20 {
			s = append(s, i+1)
		}

		batches := Batch(s, 20)
		require.Equal(t,
			[][]int{s},
			batches,
		)
	})

	t.Run("even batch", func(t *testing.T) {
		t.Parallel()

		s := make([]int, 0, 20)

		for i := range 20 {
			s = append(s, i+1)
		}

		batches := Batch(s, 5)
		require.Equal(t,
			[][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
			},
			batches,
		)
	})

	t.Run("uneven batch", func(t *testing.T) {
		t.Parallel()

		s := make([]int, 0, 23)

		for i := range 23 {
			s = append(s, i+1)
		}

		batches := Batch(s, 5)
		require.Equal(t,
			[][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
				{21, 22, 23},
			},
			batches,
		)
	})
}
