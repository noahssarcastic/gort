package ray

import (
	"testing"
)

type ObjectMock struct{}

func (obj *ObjectMock) Intersect(ray Ray) []Intersection {
	return []Intersection{}
}

func intersectEqual(i1, i2 *Intersection) bool {
	if i1.t != i2.t {
		return false
	}
	if i1.object != i2.object {
		return false
	}
	return true
}

func TestHit_all_positive(t *testing.T) {
	obj := &ObjectMock{}
	i1 := NewIntersection(1, obj)
	i2 := NewIntersection(2, obj)
	xs := make([]Intersection, 0, 2)
	xs = InsertIntersection(xs, i1)
	xs = InsertIntersection(xs, i2)
	want := i1
	got := Hit(xs)
	if !intersectEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestHit_some_negative(t *testing.T) {
	obj := &ObjectMock{}
	i1 := NewIntersection(-1, obj)
	i2 := NewIntersection(1, obj)
	xs := make([]Intersection, 0, 2)
	xs = InsertIntersection(xs, i1)
	xs = InsertIntersection(xs, i2)
	want := i2
	got := Hit(xs)
	if !intersectEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestHit_all_negative(t *testing.T) {
	obj := &ObjectMock{}
	i1 := NewIntersection(-2, obj)
	i2 := NewIntersection(-1, obj)
	xs := make([]Intersection, 0, 2)
	xs = InsertIntersection(xs, i1)
	xs = InsertIntersection(xs, i2)
	var want *Intersection = nil
	got := Hit(xs)
	if want != got {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestHit_unsorted(t *testing.T) {
	obj := &ObjectMock{}
	i1 := NewIntersection(5, obj)
	i2 := NewIntersection(7, obj)
	i3 := NewIntersection(-3, obj)
	i4 := NewIntersection(2, obj)
	xs := make([]Intersection, 0, 2)
	xs = InsertIntersection(xs, i1)
	xs = InsertIntersection(xs, i2)
	xs = InsertIntersection(xs, i3)
	xs = InsertIntersection(xs, i4)
	want := i4
	got := Hit(xs)
	if !intersectEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}
