package tree

type BinaryTree struct {
	root *treeNode
}

type treeNode struct {
	left  *treeNode
	right *treeNode
	Value interface{}
}

type EachNodeHandler func(node *treeNode)

var (
	doForEachNode EachNodeHandler
)

func setTreeNodeHandler(f EachNodeHandler) {
	doForEachNode = f
}

// 递归遍历 前序遍历
func (tn *treeNode) recursiveTraversal() {
	doForEachNode(tn)

	if tn.left == nil && tn.right == nil {
		return
	}

	if tn.left != nil {
		tn.left.recursiveTraversal()
	}
	if tn.right != nil {
		tn.right.recursiveTraversal()
	}
}

// 层次遍历 采用队列进行
func (tn *treeNode) levelTraversal(queue chan *treeNode) {
	queue <- tn
	for {
		select {
		case popNode := <-queue:
			doForEachNode(popNode)
			if popNode.left != nil {
				queue <- popNode.left
			}
			if popNode.right != nil {
				queue <- popNode.right
			}

		// no data in queue
		default:
			return
		}
	}
}

// 非递归遍历 采用堆栈进行
type stack interface {
	Pop() *treeNode
	Push(node *treeNode)
	Fetch() *treeNode
}

// 非递归遍历 前序
func (tn *treeNode) noRecursivePreOrderTraversal(s stack) {
	if tn == nil {
		return
	}

	node := tn
	s.Push(node)
	for {
		popValue := s.Pop()
		if popValue == nil {
			break
		}
		node := popValue
		doForEachNode(node)
		if node.right != nil {
			s.Push(node.right)
		}
		if node.left != nil {
			s.Push(node.left)
		}
	}
}

// 非递归遍历 中序
func (tn *treeNode) noRecursiveMiddleOrderTraversal(s stack) {
	if tn == nil {
		return
	}

	node := tn
	for {
		for node.left != nil {
			s.Push(node)
			node = node.left
		}

	Right:
		doForEachNode(node)
		if node.right != nil {
			node = node.right
		} else {
			popValue := s.Pop()
			if popValue != nil {
				node = popValue
				goto Right
			}

			// stack is empty, end loop
			break
		}
	}
}

// 非递归遍历 后序
func (tn *treeNode) noRecursivePostOrderTraversal(s stack) {
	if tn == nil {
		return
	}

	traversalMap := make(map[*treeNode]uint8)
	node := tn
	s.Push(node)
	for {
		popValue := s.Pop()
		if popValue != nil {
			node = popValue
			if !checkNodeTraversal(traversalMap, node) {
				setNodeHasTraversal(traversalMap, node)
				s.Push(node)
				if node.right != nil || node.left != nil {
					if node.right != nil {
						s.Push(node.right)
					}
					if node.left != nil {
						s.Push(node.left)
					}
					continue
				}
			} else {
				doForEachNode(node)
			}
		} else {
			// stack is empty, then end
			break
		}
	}
}

func checkNodeTraversal(m map[*treeNode]uint8, node *treeNode) bool {
	if v, ok := m[node]; ok {
		if v == 1 {
			return true
		}
	}

	return false
}

func setNodeHasTraversal(m map[*treeNode]uint8, node *treeNode) {
	m[node] = 1
}

// 后序遍历 非递归 改进版本
func (tn *treeNode) noRecursivePostOrderTraversalV2(s stack) {
	if tn == nil {
		return
	}

	node := tn
	s.Push(node)

	traversalMap := make(map[*treeNode]uint8)
	for v := s.Fetch(); v != nil; v = s.Fetch() {
		node := v

		leftNode := node.left
		rightNode := node.right
		if leftNode != nil && !checkNodeTraversal(traversalMap, leftNode) {
			s.Push(leftNode)
			continue
		}

		if rightNode != nil && !checkNodeTraversal(traversalMap, rightNode) {
			s.Push(rightNode)
			continue
		}

		popNode := s.Pop()
		doForEachNode(popNode)
		setNodeHasTraversal(traversalMap, popNode)
	}
}

// 非递归遍历 中序 改进版本
func (tn *treeNode) noRecursiveMiddleOrderTraversalV2(s stack) {
	if tn == nil {
		return
	}

	node := tn

	for {
		if node != nil {
			s.Push(node)
			node = node.left
		} else {
			node = s.Pop()
			if node == nil {
				// empty
				break
			}
			doForEachNode(node)

			if node.right != nil {
				node = node.right
			} else {
				// no right, for next pop
				node = nil
			}
		}
	}
}
