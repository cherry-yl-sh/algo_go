package dylan_algo

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	nodeQue := []*TreeNode{root}
	nodeValQue := []int{root.Val}

	for len(nodeQue) != 0 {
		popCurNode := nodeQue[0]
		popCurNodeVal := nodeValQue[0]
		nodeQue = nodeQue[1:]
		if popCurNode.Left == nil && popCurNode.Right == nil {
			if popCurNodeVal == targetSum {
				return true
			}
		}
		if popCurNode.Left != nil {
			leftNode := popCurNode.Left
			leftNodeValue := leftNode.Val

			pushVal := popCurNodeVal + leftNodeValue
			nodeQue = append(nodeQue, leftNode)
			nodeValQue = append(nodeValQue, pushVal)
		}
		if popCurNode.Right != nil {
			rightNode := popCurNode.Left
			rightNodeValue := rightNode.Val

			pushVal := popCurNodeVal + rightNodeValue
			nodeQue = append(nodeQue, rightNode)
			nodeValQue = append(nodeValQue, pushVal)
		}

	}
	return false
}
