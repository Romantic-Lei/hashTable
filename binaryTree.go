package main
import (
	"fmt"
)

type Hero struct {
	No int
	Name string
	Left *Hero
	Right *Hero
}

// 前序遍历[先输出 root 结点， 然后在输出左子树，然后在输出右子树]
func PreOrder(node *Hero) {
	if node != nil {
		fmt.Printf("no = %d name = %s \n", node.No, node.Name)
		PreOrder(node.Left) // 遍历左子树
		PreOrder(node.Right) // 遍历右子树
	}
}

// 中序遍历[先输出 root 的左子树， 然后在输出 root 结点，然后在输出 root 的右子树]
func InfixOrder(node *Hero) {
	if node != nil {
		InfixOrder(node.Left) // 遍历左子树
		fmt.Printf("no = %d name = %s \n", node.No, node.Name)
		InfixOrder(node.Right) // 遍历右子树
	}
}

// 后序遍历[先输出 root 的左子树， 然后在输出 root 的右子树，然后在输出 root 结点]
func PostOrder(node *Hero) {
	if node != nil {
		PostOrder(node.Left) // 遍历左子树
		PostOrder(node.Right) // 遍历右子树
		fmt.Printf("no = %d name = %s \n", node.No, node.Name)
	}
}

func main() {
	// 构建一个二叉树
	root := &Hero{
		No : 1,
		Name : "宋江",
	}

	left1 := &Hero{
		No : 2,
		Name : "吴用",
	}

	right1 := &Hero{
		No : 3,
		Name : "卢俊义",
	}

	heroLC := &Hero{
		No : 5,
		Name : "林冲",
	}

	heroWS := &Hero{
		No : 6,
		Name : "武松",
	}

	root.Left = left1
	root.Right = right1
	left1.Left = heroLC
	left1.Right = heroWS

	right2 := &Hero{
		No : 4,
		Name : "林冲",
	}
	right1.Right = right2

	fmt.Println("前序遍历：")
	PreOrder(root)
	// no = 1 name = 宋江
	// no = 2 name = 吴用
	// no = 5 name = 林冲
	// no = 6 name = 武松
	// no = 3 name = 卢俊义
	// no = 4 name = 林冲
	fmt.Println("中序遍历：")
	InfixOrder(root)
	// no = 5 name = 林冲
	// no = 2 name = 吴用
	// no = 6 name = 武松
	// no = 1 name = 宋江
	// no = 3 name = 卢俊义
	// no = 4 name = 林冲
	fmt.Println("后序遍历：")
	PostOrder(root)
	// no = 5 name = 林冲
	// no = 6 name = 武松
	// no = 2 name = 吴用
	// no = 4 name = 林冲
	// no = 3 name = 卢俊义
	// no = 1 name = 宋江
}