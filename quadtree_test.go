package quadtree

import (
	"math"
	"math/rand"
	"testing"
)

type FlatStore struct {
	Points []*Point2D
}

func (self *FlatStore) Insert(point *Point2D) {
	self.Points = append(self.Points, point)
}

func (self *FlatStore) Fetch(boundary *AABB) []*Point2D {
	points := []*Point2D{}
	for _, point := range self.Points {
		if boundary.Contains(point) {
			points = append(points, point)
		}
	}
	return points
}

func TestBasicFunctionality(t *testing.T) {
	// create a quadtree covering a 2000x2000 area
	tree := NewQuadTreeNode(&AABB{0, 0, 1000, 1000}, 2)

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

func BenchmarkInsertFlat100(b *testing.B)     { benchmarkInsertFlat(100, b) }
func BenchmarkInsertFlat1000(b *testing.B)    { benchmarkInsertFlat(1000, b) }
func BenchmarkInsertFlat10000(b *testing.B)   { benchmarkInsertFlat(10000, b) }
func BenchmarkInsertFlat100000(b *testing.B)  { benchmarkInsertFlat(100000, b) }
func BenchmarkInsertFlat1000000(b *testing.B) { benchmarkInsertFlat(1000000, b) }

func benchmarkInsert(size int, b *testing.B) {
	boundary := &AABB{0, 0, 1000, 1000}
	tree := buildTree(boundary, size)
	for n := 0; n < b.N; n++ {
		tree.Insert(&Point2D{rand.Float64()*2000 - 1000, rand.Float64()*2000 - 1000})
	}
}

func benchmarkInsertFlat(size int, b *testing.B) {
	boundary := &AABB{0, 0, 1000, 1000}
	store := buildFlatStore(boundary, size)
	for n := 0; n < b.N; n++ {
		store.Insert(&Point2D{rand.Float64()*2000 - 1000, rand.Float64()*2000 - 1000})
	}
}

func BenchmarkFetch100(b *testing.B)     { benchmarkFetch(100, b) }
func BenchmarkFetch1000(b *testing.B)    { benchmarkFetch(1000, b) }
func BenchmarkFetch10000(b *testing.B)   { benchmarkFetch(10000, b) }
func BenchmarkFetch100000(b *testing.B)  { benchmarkFetch(100000, b) }
func BenchmarkFetch1000000(b *testing.B) { benchmarkFetch(1000000, b) }

func BenchmarkFetchFlat100(b *testing.B)     { benchmarkFetchFlat(100, b) }
func BenchmarkFetchFlat1000(b *testing.B)    { benchmarkFetchFlat(1000, b) }
func BenchmarkFetchFlat10000(b *testing.B)   { benchmarkFetchFlat(10000, b) }
func BenchmarkFetchFlat100000(b *testing.B)  { benchmarkFetchFlat(100000, b) }
func BenchmarkFetchFlat1000000(b *testing.B) { benchmarkFetchFlat(1000000, b) }

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

func benchmarkFetchFlat(size int, b *testing.B) {
	boundary := &AABB{0, 0, 1000, 1000}
	store := buildFlatStore(boundary, size)

	for n := 0; n < b.N; n++ {
		store.Fetch(&AABB{
			rand.Float64()*boundary.Width() - boundary.HalfWidth,
			rand.Float64()*boundary.Height() - boundary.HalfHeight,
			math.Min(rand.Float64()*boundary.Width()-boundary.HalfWidth, boundary.HalfWidth),
			math.Min(rand.Float64()*boundary.Height()-boundary.HalfHeight, boundary.HalfHeight),
		})
	}
}

func buildTree(boundary *AABB, nodeCount int) *QuadTreeNode {
	tree := NewQuadTreeNode(boundary, 4)

	for n := 0; n < nodeCount; n++ {
		tree.Insert(&Point2D{
			rand.Float64()*boundary.Width() - boundary.HalfWidth,
			rand.Float64()*boundary.Height() - boundary.HalfHeight,
		})
	}

	return tree
}

func buildFlatStore(boundary *AABB, nodeCount int) *FlatStore {
	store := &FlatStore{}

	for n := 0; n < nodeCount; n++ {
		store.Insert(&Point2D{
			rand.Float64()*boundary.Width() - boundary.HalfWidth,
			rand.Float64()*boundary.Height() - boundary.HalfHeight,
		})
	}

	return store
}
