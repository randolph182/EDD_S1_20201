
package Estructuras

type BTNode struct{
	Item int
	Left *BTNode
	Right *BTNode
	Balance int
}

const (
	LEFT_HEAVY = -1
	BALANCED = 0
	RIGHT_HEAVY = +1
)


//Constructor
func New_BTNode(item int) *BTNode{
	return &BTNode{item,nil,nil,0}
}

//Constructor_2
func New_BTNode_2(item int, left_child *BTNode, right_child *BTNode) *BTNode{
	return &BTNode{item,left_child,right_child,0}
}

