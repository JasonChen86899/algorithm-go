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
	Pop() interface{}
	Push(interface{})
}

// 非递归遍历 前序
func (tn *treeNode) noRecursivePreOrderTraversal(s stack) {
	s.Push(tn)
	for {
		popValue := s.Pop()
		if popValue == nil {
			break
		}
		node := popValue.(*treeNode)
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
				node = popValue.(*treeNode)
				goto Right
			}

			// stack is empty, end for
			break
		}
	}
}

// 非递归遍历 后序
func (tn *treeNode) noRecursivePostOrderTraversal(s stack) {
	traversalMap := make(map[*treeNode]uint8)
	node := tn
	s.Push(node)
	for {
		popValue := s.Pop()
		if popValue != nil {
			node = popValue.(*treeNode)
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
