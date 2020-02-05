package main
import (
	"fmt"
	"os"
)
// 定义 emp
type Emp struct {
	Id int
	Name string
	Next *Emp
}

// 定义 EmpLink
// 第一个的 EmpLink 不带表头， 即第一个节点就存放雇员
type EmpLink struct {
	Head *Emp
}

// 定义Hashtable， 含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

// 1. 添加员工的方法,保证添加时， 编号是从下到大
func (this *EmpLink) Insert(emp *Emp) {
	current := this.Head // 辅助指针
	var pre *Emp = nil // 这是一个辅助指针 pre 在 current 前面
	// 如果当前的 EmpLink 就是一个空链表
	if current == nil {
		this.Head = emp // 完成
		return
	}
	// 如果不是 一个空链表， 给emp找到对应的位置并插入
	// 让 current 和 emp 比较， 然后让 pre 保持在 current 前面
	for {
		if current != nil {
			// 比较
			if current.Id > emp.Id {
				// 找到位置
				break
			}
			pre = current // 保证同步
			current = current.Next
		} else {
			break
		}
	}
	// 退出时，我们看一下是否将emp 添加到链表最后
	pre.Next = emp
	emp.Next = current
}

// 显示链表的信息
func (this *EmpLink) ShowLink(no int) {
	if this.Head == nil {
		fmt.Printf("链表 %d 为空\n", no)
		return 
	}
	// 遍历当前链表，并显示数据
	current := this.Head // 辅助的指针
	for {
		if current != nil {
			fmt.Printf("链表%d - 雇员id=%d - 名字=%s -> ", no, current.Id, current.Name)
			current = current.Next
		} else {
			break
		}
	}
	fmt.Println()
}

// 给HashTable 编写 Insert 雇员的方法
func (this *HashTable) Insert(emp *Emp) {
	// 使用散列函数，确定将该雇员添加到哪个链表
	linkNo := this.HashFun(emp.Id)
	// 使用对应的链表添加
	this.LinkArr[linkNo].Insert(emp)
}

// 编写方法， 显示Hashtable的所有雇员
func (this *HashTable) ShowAll() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
	}
}

// 编写一个散列方法
func (this *HashTable) HashFun(id int) int {
	return id % 7 // 得到的值，就是对应的链表的下标
}

func main() {
	key := ""
	id := 0
	name := ""
	var hashtable HashTable
	for {
		fmt.Println("=================雇员系统菜单====================")
		fmt.Println("input 表示添加雇员")
		fmt.Println("show  表示显示雇员")
		fmt.Println("find  表示查找雇员")
		fmt.Println("exit  表示退出系统员")
		fmt.Println("请输入你的选择")
		fmt.Scanln(&key)
		switch key {
			case "input" :
				fmt.Println("输入雇员 id")
				fmt.Scanln(&id)
				fmt.Println("输入雇员姓名")
				fmt.Scanln(&name)
				emp := &Emp{
					Id : id,
					Name : name,
				}
				hashtable.Insert(emp)
			case "show" :
				hashtable.ShowAll()
			case "find" :
				
			case "exit" :
				os.Exit(0)
			default :
				fmt.Println("输入有误")
			
			}

	}
}
