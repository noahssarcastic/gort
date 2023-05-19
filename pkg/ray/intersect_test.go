package ray

import (
	"errors"
	"testing"

	"github.com/noahssarcastic/gort/pkg/material"
	"github.com/noahssarcastic/gort/pkg/tuple"
)

type ObjectMock struct{}

func (objPtr *ObjectMock) Intersect(ray Ray) []Intersect {
	return []Intersect{}
}

func (objPtr *ObjectMock) NormalAt(pt tuple.Tuple) tuple.Tuple {
	return tuple.Vector(0, 0, 0)
}

func (objPtr *ObjectMock) Material() material.Material {
	return material.Default()
}

func intersectEqual(i1, i2 Intersect) bool {
	if i1.t != i2.t {
		return false
	}
	if i1.objPtr != i2.objPtr {
		return false
	}
	return true
}

func TestHit_all_positive(t *testing.T) {
	objPtr := &ObjectMock{}
	i1 := NewIntersect(1, objPtr)
	i2 := NewIntersect(2, objPtr)
	xs := make([]Intersect, 0, 2)
	xs = InsertIntersect(xs, i1)
	xs = InsertIntersect(xs, i2)
	want := i1
	got, err := Hit(xs)
	if errors.Is(err, ErrNoHits) {
		t.Error(err)
	} else if !intersectEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestHit_some_negative(t *testing.T) {
	objPtr := &ObjectMock{}
	i1 := NewIntersect(-1, objPtr)
	i2 := NewIntersect(1, objPtr)
	xs := make([]Intersect, 0, 2)
	xs = InsertIntersect(xs, i1)
	xs = InsertIntersect(xs, i2)
	want := i2
	got, err := Hit(xs)
	if errors.Is(err, ErrNoHits) {
		t.Error(err)
	} else if !intersectEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestHit_all_negative(t *testing.T) {
	objPtr := &ObjectMock{}
	i1 := NewIntersect(-2, objPtr)
	i2 := NewIntersect(-1, objPtr)
	xs := make([]Intersect, 0, 2)
	xs = InsertIntersect(xs, i1)
	xs = InsertIntersect(xs, i2)
	_, err := Hit(xs)
	if !errors.Is(err, ErrNoHits) {
		t.Errorf("want %v; got %v", ErrNoHits, err)
	}
}

func TestHit_unsorted(t *testing.T) {
	objPtr := &ObjectMock{}
	i1 := NewIntersect(5, objPtr)
	i2 := NewIntersect(7, objPtr)
	i3 := NewIntersect(-3, objPtr)
	i4 := NewIntersect(2, objPtr)
	xs := make([]Intersect, 0, 2)
	xs = InsertIntersect(xs, i1)
	xs = InsertIntersect(xs, i2)
	xs = InsertIntersect(xs, i3)
	xs = InsertIntersect(xs, i4)
	want := i4
	got, err := Hit(xs)
	if errors.Is(err, ErrNoHits) {
		t.Error(err)
	} else if !intersectEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}
