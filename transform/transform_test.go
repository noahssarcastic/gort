package transform

import (
	"testing"

	"github.com/noahssarcastic/tddraytracer/tuple"
)

func TestTranslate_point(t *testing.T) {
	tform := Translate(5, -3, 2)
	pt := tuple.Point(-3, 4, 5)
	want := tuple.Point(2, 1, 7)
	got := tform.MultTuple(pt)
	if !tuple.Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestTranslate_vector(t *testing.T) {
	tform := Translate(5, -3, 2)
	vec := tuple.Vector(-3, 4, 5)
	want := vec
	got := tform.MultTuple(vec)
	if !tuple.Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestScale(t *testing.T) {

}

func TestRotate(t *testing.T) {

}

func TestShear(t *testing.T) {

}

func TestChainTransforms(t *testing.T) {

}
