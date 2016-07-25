package quadtree

// Quadrant is a type that represents one of the 4 locations for a quadrant
// within a QuadTreeNode
type Quadrant int

const (
	NorthWest Quadrant = iota
	NorthEast
	SouthWest
	SouthEast
)

// Point2D represents a point in 2D space
type Point2D struct{ X, Y float64 }

// AABB is an axis-aligned bounding box. It represents a rectangular region in
// 2D cartesian space. The X and Y values represent the cartesian coordinates
// for the center of the bounding box. HalfWidth and HalfHeight represent
// one-half of the bounding box's width and height, respectively
type AABB struct{ X, Y, HalfWidth, HalfHeight float64 }

// Width returns the full width of the bounding box
func (self AABB) Width() float64 { return self.HalfWidth * 2 }

// Height returns the full height of the bounding box
func (self AABB) Height() float64 { return self.HalfHeight * 2 }

// Contains returns whether the point passed to it is contained within the
// boundaries of the bounding box
func (self AABB) Contains(point *Point2D) bool {
	return point.X >= self.X-self.HalfWidth && point.X < self.X+self.HalfWidth &&
		point.Y >= self.Y-self.HalfHeight && point.Y < self.Y+self.HalfHeight
}

// Intersects returns whether the passed bounding box intersects this one
func (self AABB) Intersects(boundary *AABB) bool {
	return boundary.X < self.X+self.HalfWidth &&
		boundary.X+(boundary.Width()) > self.X-self.HalfWidth &&
		boundary.Y < self.Y+self.HalfHeight &&
		boundary.Y+(boundary.Height()) > self.Y-self.HalfHeight
}

// QuadTreeNode represents a node within the quadtree. It has a rectangular
// boundary, and can either contain a Location or 4 QuadTreeNodes representing
// the children of this node.
type QuadTreeNode struct {
	*AABB
	Location  *Point2D
	Quadrants [4]*QuadTreeNode
}

// NewQuadTreeNode returns an empty QuadTreeNode for the passed-in boundary
func NewQuadTreeNode(boundary *AABB) *QuadTreeNode {
	return &QuadTreeNode{boundary, nil, [4]*QuadTreeNode{}}
}

// Fetch returns the points contained within the passed in boundary
func (self *QuadTreeNode) Fetch(boundary *AABB) []*Point2D {
	if self.Location != nil {
		if self.Location.X >= boundary.X &&
			self.Location.X < boundary.X+boundary.Width() &&
			self.Location.Y >= boundary.Y &&
			self.Location.Y < boundary.Y+boundary.Height() {
			return []*Point2D{self.Location}
		} else {
			return []*Point2D{}
		}
	}

	if self.Quadrants[0] == nil {
		return []*Point2D{}
	}

	points := []*Point2D{}
	for _, node := range self.Quadrants {
		if node.Intersects(boundary) {
			nodes := node.Fetch(boundary)
			if len(nodes) > 0 {
				points = append(points, nodes...)
			}
		}
	}
	return points
}

// Insert adds a point into the quadtree, subdividing the node if necessary
func (self *QuadTreeNode) Insert(point *Point2D) {
	if point == nil || !self.Contains(point) {
		return
	} else if self.Location == nil && self.Quadrants[0] == nil {
		self.Location = point
	} else if self.Location == nil && self.Quadrants[0] != nil {
		self.update(point)
	} else if self.Location != nil && self.Quadrants[0] == nil {
		self.subdivide()
		self.update(self.Location)
		self.Location = nil
		self.update(point)
	} else {
		panic("This should never run. Node has been subdivided and is not null")
	}
}

func (self *QuadTreeNode) subdivide() {
	width := self.HalfHeight / 2
	height := self.HalfHeight / 2
	self.Quadrants[NorthWest] = NewQuadTreeNode(
		&AABB{self.X - width, self.Y + height, width, height})
	self.Quadrants[NorthEast] = NewQuadTreeNode(
		&AABB{self.X + width, self.Y + height, width, height})
	self.Quadrants[SouthWest] = NewQuadTreeNode(
		&AABB{self.X - width, self.Y - height, width, height})
	self.Quadrants[SouthEast] = NewQuadTreeNode(
		&AABB{self.X + width, self.Y - height, width, height})
}

func (self *QuadTreeNode) update(point *Point2D) {
	for i := 0; i < 4; i++ {
		self.Quadrants[i].Insert(point)
	}
}
