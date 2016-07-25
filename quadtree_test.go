package quadtree

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func TestBasicFunctionality(t *testing.T) {
	// create a quadtree covering a 2000x2000 area
	tree := NewQuadTreeNode(&AABB{0, 0, 1000, 1000})

	// insert some points into the tree
	tree.Insert(&Point2D{23, 42})
	tree.Insert(&Point2D{6, 29})
	tree.Insert(&Point2D{86, 14})
	tree.Insert(&Point2D{35, 46})

	// fetch a region of the tree
	points := tree.Fetch(&AABB{30, 44, 10, 10})

	if !containsPoint(&Point2D{23, 42}, points) {
		t.Error("Result does not contain point (23, 42)")
	}

	if !containsPoint(&Point2D{35, 46}, points) {
		t.Error("Result does not contain point(35, 46)")
	}
}

func containsPoint(needle *Point2D, haystack []*Point2D) bool {
	for _, point := range haystack {
		if point.X == needle.X && point.Y == needle.Y {
			return true
		}
	}
	return false
}

func BenchmarkInsert100(b *testing.B)     { benchmarkInsert(100, b) }
func BenchmarkInsert1000(b *testing.B)    { benchmarkInsert(1000, b) }
func BenchmarkInsert10000(b *testing.B)   { benchmarkInsert(10000, b) }
func BenchmarkInsert100000(b *testing.B)  { benchmarkInsert(100000, b) }
func BenchmarkInsert1000000(b *testing.B) { benchmarkInsert(1000000, b) }

func benchmarkInsert(size int, b *testing.B) {
	boundary := &AABB{0, 0, 1000, 1000}
	tree := buildTree(boundary, size)
	for n := 0; n < b.N; n++ {
		tree.Insert(&Point2D{rand.Float64()*2000 - 1000, rand.Float64()*2000 - 1000})
	}
}

func BenchmarkFetch100(b *testing.B)     { benchmarkFetch(100, b) }
func BenchmarkFetch1000(b *testing.B)    { benchmarkFetch(1000, b) }
func BenchmarkFetch10000(b *testing.B)   { benchmarkFetch(10000, b) }
func BenchmarkFetch100000(b *testing.B)  { benchmarkFetch(100000, b) }
func BenchmarkFetch1000000(b *testing.B) { benchmarkFetch(1000000, b) }

func benchmarkFetch(size int, b *testing.B) {
	boundary := &AABB{0, 0, 1000, 1000}
	tree := buildTree(boundary, size)

	for n := 0; n < b.N; n++ {
		tree.Fetch(&AABB{
			rand.Float64()*boundary.Width() - boundary.HalfWidth,
			rand.Float64()*boundary.Height() - boundary.HalfHeight,
			math.Min(rand.Float64()*boundary.Width()-boundary.HalfWidth, boundary.HalfWidth),
			math.Min(rand.Float64()*boundary.Height()-boundary.HalfHeight, boundary.HalfHeight),
		})
	}
}

func buildTree(boundary *AABB, nodeCount int) *QuadTreeNode {
	tree := NewQuadTreeNode(boundary)

	for n := 0; n < nodeCount; n++ {
		tree.Insert(&Point2D{
			rand.Float64()*boundary.Width() - boundary.HalfWidth,
			rand.Float64()*boundary.Height() - boundary.HalfHeight,
		})
	}

	return tree
}
