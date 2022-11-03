package geometry

import (
	"fmt"
	"testing"

	"github.com/noahssarcastic/gort/math"
)

func TestPosition(t *testing.T) {
	r := NewRay(math.Point(2, 3, 4), math.Vector(1, 0, 0))
	tests := []struct {
		ray  Ray
		t    float64
		want math.Tuple
	}{
		{r, 0, math.Point(2, 3, 4)},
		{r, 1, math.Point(3, 3, 4)},
		{r, -1, math.Point(1, 3, 4)},
		{r, 2.5, math.Point(4.5, 3, 4)},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%v,%v", tt.ray, tt.t)
		t.Run(name, func(t *testing.T) {
			ans := Position(tt.ray, tt.t)
			if !math.TupleEqual(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
