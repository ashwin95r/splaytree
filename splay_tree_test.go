package splaytree

import (
	"sort"
	"testing"
)

func TestSplayTree(t *testing.T) {
	tree := &splayTree{}
	n := 40000
	gap := 307
	for i := gap; i != 0; i = (i + gap) % n {
		tree.insert(Int(i))
	}

	for i := gap; i != 0; i = (i + gap) % n {
		g := tree.Get(Int(i))
		if g != Int(i) {
			t.Errorf("g = %v, want %v", g, i)
		}
	}
}

func TestSplayDeleteMin(t *testing.T) {
	tree := &splayTree{}
	n := 40000
	gap := 307
	var insert []Int
	for i := gap; i != 0; i = (i + gap) % n {
		tree.insert(Int(i))
		insert = append(insert, Int(i))
	}

	sort.Slice(insert, func(i, j int) bool {
		return insert[i] < insert[j]
	})
	for i := gap; i != 0; i = (i + gap) % n {
		g := tree.DeleteMin()
		if g != insert[0] {
			t.Errorf("g = %v, want %v", g, i)
		}
		insert = insert[1:]
	}
}

func TestSplayDeleteMax(t *testing.T) {
	tree := &splayTree{}
	n := 40000
	gap := 307
	var insert []Int
	for i := gap; i != 0; i = (i + gap) % n {
		tree.insert(Int(i))
		insert = append(insert, Int(i))
	}

	sort.Slice(insert, func(i, j int) bool {
		return insert[i] > insert[j]
	})
	for i := gap; i != 0; i = (i + gap) % n {
		g := tree.DeleteMax()
		if g != insert[0] {
			t.Errorf("g = %v, want %v", g, i)
		}
		insert = insert[1:]
	}
}

// Int implements the Item interface for integers.
type Int int

// Less returns true if int(a) < int(b).
func (a Int) Less(b Item) bool {
	return a < b.(Int)
}
