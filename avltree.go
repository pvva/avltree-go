package avltree

type Comparable interface {
	CompareTo(i interface{}) int
}

type AvlTree struct {
	root *avlNode

	avgDistance float64
	size        int
}

type avlNode struct {
	key Comparable

	left, right *avlNode
	height      int
}

// node private methods

func height(node *avlNode) int {
	if node == nil {
		return -1
	}
	return node.height
}

func (node *avlNode) balanceFactor() int {
	return height(node.right) - height(node.left)
}

func (node *avlNode) fixHeight() {
	hl := height(node.left)
	hr := height(node.right)

	if hl > hr {
		node.height = hl + 1
	} else {
		node.height = hr + 1
	}
}

// tree

func NewAvlTree() *AvlTree {
	return &AvlTree{}
}

func (tree *AvlTree) Insert(key Comparable) {
	tree.root = tree.insert(tree.root, key)
}

func (tree *AvlTree) Remove(key Comparable) {
	tree.root = tree.remove(tree.root, key)
}

func (tree *AvlTree) Traverse(f func(c Comparable, level int) bool) {
	tree.traverse(tree.root, f)
}

func (tree *AvlTree) Height() int {
	return height(tree.root)
}

// tree private methods

func (tree *AvlTree) insert(after *avlNode, key Comparable) *avlNode {
	if after == nil {
		return &avlNode{key: key}
	}

	if key.CompareTo(after.key) < 0 {
		after.left = tree.insert(after.left, key)
	} else {
		after.right = tree.insert(after.right, key)
	}

	return tree.balance(after)
}

func (tree *AvlTree) rotateLeft(node *avlNode) *avlNode {
	if node == nil || node.right == nil {
		return node
	}

	p := node.right
	node.right = p.left
	p.left = node

	node.fixHeight()
	p.fixHeight()

	return p
}

func (tree *AvlTree) rotateRight(node *avlNode) *avlNode {
	if node == nil || node.left == nil {
		return node
	}

	q := node.left
	node.left = q.right
	q.right = node

	node.fixHeight()
	q.fixHeight()

	return q
}

func (tree *AvlTree) balance(node *avlNode) *avlNode {
	node.fixHeight()

	bFactor := node.balanceFactor()

	if bFactor == 2 {
		if node.right != nil && node.right.balanceFactor() < 0 {
			node.right = tree.rotateRight(node.right)
		}
		return tree.rotateLeft(node)
	}
	if bFactor == -2 {
		if node.left != nil && node.left.balanceFactor() > 0 {
			node.left = tree.rotateLeft(node.left)
		}
		return tree.rotateRight(node)
	}

	return node
}

func (tree *AvlTree) findMin(from *avlNode) *avlNode {
	if from.left != nil {
		return tree.findMin(from.left)
	}
	return from
}

func (tree *AvlTree) removeMin(from *avlNode) *avlNode {
	if from.left == nil {
		return from.right
	}
	from.left = tree.removeMin(from.left)

	return tree.balance(from)
}

func (tree *AvlTree) remove(from *avlNode, key Comparable) *avlNode {
	if from == nil {
		return nil
	}

	compare := key.CompareTo(from.key)
	if compare < 0 {
		from.left = tree.remove(from.left, key)
	} else if compare > 0 {
		from.right = tree.remove(from.right, key)
	} else {
		q := from.left
		r := from.right

		if r == nil {
			return tree.remove(q, key)
		}
		for {
			min := tree.findMin(r)
			right := tree.removeMin(r)

			if min.key.CompareTo(key) == 0 {
				if right == nil {
					return nil
				}

				r = right
			} else {
				min.right = right
				min.left = tree.remove(q, key)

				return tree.balance(min)
			}
		}
	}
	return tree.balance(from)
}

func (tree *AvlTree) traverse(from *avlNode, f func(c Comparable, level int) bool) bool {
	if from == nil {
		return true
	}
	if !tree.traverse(from.left, f) || !f(from.key, height(from)) || !tree.traverse(from.right, f) {
		return false
	}

	return true
}
