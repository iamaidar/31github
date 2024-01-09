package problems

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	first := []int{}
	second := []int{}
	obhod(root1, &first)
	obhod(root2, &second)

	if len(first) != len(second) {
		return false
	}

	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			return false
		}
	}

	return true
}

func obhod(root *TreeNode, nums *[]int) {

	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*nums = append(*nums, root.Val)
	}

	if root.Left != nil {
		obhod(root.Left, nums)
	}
	if root.Right != nil {
		obhod(root.Right, nums)
	}

}
