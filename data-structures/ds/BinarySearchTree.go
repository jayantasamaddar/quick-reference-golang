package ds

type TreeNode[T Number] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

type BinarySearchTree[T Number] struct {
	Root *TreeNode[T]
}

// Add to Binary Search Tree
func (this *BinarySearchTree[T]) Add(items ...T) {
	for _, item := range items {
		node := this.Root
		if node == nil {
			this.Root = &TreeNode[T]{Value: item}
		} else {
			var searchAndInsert func(node *TreeNode[T]) error
			searchAndInsert = func(node *TreeNode[T]) error {
				if item < node.Value {
					if node.Left == nil {
						node.Left = &TreeNode[T]{Value: item}
						return nil
					} else {
						return searchAndInsert(node.Left)
					}
				} else {
					if node.Right == nil {
						node.Right = &TreeNode[T]{Value: item}
						return nil
					} else {
						return searchAndInsert(node.Right)
					}
				}
			}
			searchAndInsert(node)
		}
	}
}
