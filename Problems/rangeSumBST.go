package problems

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	var sum int

	if root.Val <= high && root.Val >= low {
		sum += root.Val
	}

	sum += rangeSumBST(root.Left, low, high)
	sum += rangeSumBST(root.Right, low, high)

	return sum
}
