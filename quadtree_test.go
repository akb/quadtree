package quadtree

import (
	"math"
	"math/rand"
	"testing"
)

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
