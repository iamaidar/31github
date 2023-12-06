package problems

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var nums []int
	var prev *ListNode

	for elem := head; elem != nil; {
		nums = append(nums, elem.Val)
		next := elem.Next
		elem.Next = prev
		prev = elem
		elem = next
	}
	i := 0
	for elem := prev; elem != nil; elem = elem.Next {
		if elem.Val != nums[i] {
			return false
		}
		i++
	}

	return true
}
