package tree

import "fmt"

type Noe struct {
	Value int
	Left,Right *Noe
}

func (node Noe) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Noe) SetValue(value int) {
	node.Value = value
}

func (node *Noe) Traverse() {
	if node == nil {
		return
	}
	// 中序便利

	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

// 公创函数
func CreateNode(value int) *Noe  {
	return &Noe{Value:value}
}
