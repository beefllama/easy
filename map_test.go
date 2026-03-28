package easy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKeys(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		m := make(map[string]int)
		require.Equal(t, []string{}, Keys(m))
	})

	t.Run("filled map", func(t *testing.T) {
		t.Parallel()

		m := map[string]int{
			"k1": 15,
			"k2": 25,
		}
		require.ElementsMatch(t, []string{"k1", "k2"}, Keys(m))
	})
}

func TestValues(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		m := make(map[string]int)
		require.Equal(t, []int{}, Values(m))
	})

	t.Run("filled map", func(t *testing.T) {
		t.Parallel()

		m := map[string]int{
			"k1": 15,
			"k2": 25,
		}
		require.ElementsMatch(t, []int{15, 25}, Values(m))
	})
}

func TestSliceToMap(t *testing.T) {
	t.Parallel()

	type st struct {
		id   int
		name string
	}

	toMapFunc := func(e st) (int, string) {
		return e.id, e.name
	}

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		s := []st{}
		require.Equal(t, map[int]string{}, SliceToMap(s, toMapFunc))
	})

	t.Run("filled slice", func(t *testing.T) {
		t.Parallel()

		s := []st{
			{id: 15, name: "name1"},
			{id: 25, name: "name2"},
		}
		require.Equal(t,
			map[int]string{
				15: "name1",
				25: "name2",
			},
			SliceToMap(s, toMapFunc),
		)
	})
}
