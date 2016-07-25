quadtree
========

Quadtree implementation in Go.

Quadtrees are used to store 2D points in a way which allows regions to be
queried for their contents in O(log n) time.

Usage
-----

An axis-aligned bounding box can be described using `quadtree.AABB` this struct
has a center point and a half-width and half-height and represents a
rectangular region of 2D cartesian space. A node within the quadtree has it's
containing area described by one of these bounding boxes. Regions within the
quadtree can also be queried by padding a bounding box to `Fetch`.

A new quadtree can be created by using `NewQuadTreeNode` to produce the root
node. Additional nodes will be automatically created as necessary by inserting
points into the tree using `Insert`.

    // create a quadtree covering a 2000x2000 area
    tree := NewQuadTreeNode(&AABB{0, 0, 1000, 1000})

    // insert some points into the tree
    tree.Insert(&Point2D{23, 42})
    tree.Insert(&Point2D{6, 29})
    tree.Insert(&Point2D{86, 14})
    tree.Insert(&Point2D{35, 46})

    // fetch a region of the tree
    points := tree.Fetch(&AABB{30, 44, 10, 10})

    // points == []*Point2D{&Point2D{23, 42}, &Point2D{35, 46}}
