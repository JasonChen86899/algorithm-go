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
	if tn.left == nil && tn.right == nil {
		return
	}

	doForEachNode(tn)
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
func (tn *treeNode) noRecursiveTraversal(s stack) {
	s.Push(tn)
	for {
		node := s.Pop().(*treeNode)
		doForEachNode(node)
		s.Push(node.right)
		s.Push(node.left)
	}
}
