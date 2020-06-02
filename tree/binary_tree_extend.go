package tree

// 二叉树的高度 根据递归方法 稍作修改而来
// 递归函数返回
func (tn *treeNode) treeHeight() uint64 {
	if tn.left == nil && tn.right == nil {
		return 1
	} else {
		if tn.left != nil && tn.right != nil {
			leftHeight := tn.left.treeHeight()
			rightHeight := tn.right.treeHeight()
			if leftHeight >= rightHeight {
				return leftHeight + 1
			} else {
				return rightHeight + 1
			}
		}

		if tn.left == nil {
			return tn.right.treeHeight() + 1
		}

		return tn.left.treeHeight() + 1 //tn.right == nil
	}
}


