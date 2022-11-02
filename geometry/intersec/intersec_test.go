package intersec

import (
	"testing"

	"github.com/noahssarcastic/tddraytracer/geometry/ray"
)

type ObjectMock struct{}

func (obj *ObjectMock) Intersect(ray ray.Ray) []Intersection {
	return []Intersection{}
}

func TestHit_all_positive(t *testing.T) {
	obj := &ObjectMock{}
	i1 := New(1, obj)
	i2 := New(2, obj)
	xs := Combine(i1, i2)
	want := i1
	got := xs.Hit()
	if want != got {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestHit_some_negative(t *testing.T) {
	obj := &ObjectMock{}
	i1 := New(-1, obj)
	i2 := New(1, obj)
	xs := Combine(i1, i2)
	want := i2
	got := xs.Hit()
	if want != got {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestHit_all_negative(t *testing.T) {
	obj := &ObjectMock{}
	i1 := New(-2, obj)
	i2 := New(-1, obj)
	xs := Combine(i1, i2)
	var want *Intersection = nil
	got := xs.Hit()
	if want != got {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestHit_unsorted(t *testing.T) {
	obj := &ObjectMock{}
	i1 := New(5, obj)
	i2 := New(7, obj)
	i3 := New(-3, obj)
	i4 := New(2, obj)
	xs := Combine(i1, i2, i3, i4)
	want := i4
	got := xs.Hit()
	if want != got {
		t.Errorf("want %v; got %v", want, got)
	}
}
