package main
/*
1、 如果结点 p、q 都存在且为左右结点，那么根结点 root 就是最近公共祖先；

2、 如果结点 p、q 都存在且都为左结点，那么在根结点 root 的左子树查找；

3、 如果结点 p、q 都存在且都为右结点，那么在根结点 root 的右子树查找。

*/


type treeNode struct {
	Val int
	left *treeNode
	right *treeNode
}

func lowestCommonAncestor(root, p, q *treeNode) *treeNode {
	if root == nil {
		return root
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.left, p, q)
	right := lowestCommonAncestor(root.right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

