package piscine

func BTreeIsBinary(root *TreeNode) bool {
	condLeft := true
	condRight := true

	if root == nil {
		return true
	}

	if root.Left != nil {
		condLeft = BTreeIsBinary(root.Left) && root.Data >= root.Left.Data
	}

	if root.Right != nil {
		condRight = BTreeIsBinary(root.Right) && root.Data <= root.Right.Data
	}

	return condLeft && condRight
}
