package tree

import (
	"fmt"
	"testing"
)

func initBinaryTree() *BinaryTree {
	setTreeNodeHandler(func(node *treeNode) {
		fmt.Println(node.Value)
	})

	node1 := &treeNode{
		Value: 99,
		left:  nil,
		right: nil,
	}
	node2 := &treeNode{
		Value: 98,
		left:  nil,
		right: nil,
	}
	node3 := &treeNode{
		Value: 97,
		left:  nil,
		right: nil,
	}

	leftNode := &treeNode{
		Value: 1,
		left:  nil,
		right: node1,
	}
	rightNode := &treeNode{
		Value: 2,
		left:  node2,
		right: node3,
	}
	rootNode := &treeNode{
		Value: 0,
		left:  leftNode,
		right: rightNode,
	}

	return &BinaryTree{root: rootNode}
}

func TestBinaryTreeRecursiveTraversal(T *testing.T) {
	binaryTree := initBinaryTree()
	binaryTree.root.recursiveTraversal()
}

type testStack struct {
	innerArray []interface{}
}

func NewTestStack(size uint64) *testStack {
	return &testStack{innerArray: make([]interface{}, 0, size)}
}

func (s *testStack) Push(item interface{}) {
	s.innerArray = append(s.innerArray, item)
}

func (s *testStack) Pop() interface{} {
	if len(s.innerArray) > 0 {
		item := s.innerArray[len(s.innerArray)-1]
		s.innerArray = s.innerArray[:len(s.innerArray)-1]
		return item
	}

	return nil
}

func TestBinaryTreeNotRecursiveTraversal(t *testing.T) {
	binaryTree := initBinaryTree()
	testStack := NewTestStack(1024)
	binaryTree.root.noRecursivePreOrderTraversal(testStack)
	fmt.Println("-----------------------------")
	binaryTree.root.noRecursiveMiddleOrderTraversal(testStack)
	fmt.Println("-----------------------------")
	binaryTree.root.noRecursivePostOrderTraversal(testStack)
}
