package ray

import (
	"fmt"
	"testing"

	"github.com/noahssarcastic/tddraytracer/math/tuple"
)

func TestPosition(t *testing.T) {
	r := New(tuple.Point(2, 3, 4), tuple.Vector(1, 0, 0))
	tests := []struct {
		ray  Ray
		t    float64
		want tuple.Tuple
	}{
		{r, 0, tuple.Point(2, 3, 4)},
		{r, 1, tuple.Point(3, 3, 4)},
		{r, -1, tuple.Point(1, 3, 4)},
		{r, 2.5, tuple.Point(4.5, 3, 4)},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%v,%v", tt.ray, tt.t)
		t.Run(name, func(t *testing.T) {
			ans := Position(tt.ray, tt.t)
			if !tuple.Equal(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
