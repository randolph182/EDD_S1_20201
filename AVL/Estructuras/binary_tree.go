package Estructuras

import (
	"log"
	"fmt"
)

var increase = false

type Binary_Tree struct {
	Root *BTNode
	
}

//constructor
func New_Binary_Tree() *Binary_Tree{
	return &Binary_Tree{nil}
}

//constructor2 
func New_Binary_Tree_2(left_child *Binary_Tree, right_child *Binary_Tree, item int)  *Binary_Tree{
	return &Binary_Tree{New_BTNode_2(item,left_child.Root,right_child.Root)}
}
 

func get_left_subtree(tree *Binary_Tree) *BTNode{
	if tree.Root == nil{
		log.Fatal("get left subtee on empty tree")
		return nil
	}
	return tree.Root.Left 
}

func get_right_subtree(tree *Binary_Tree) *BTNode{
	if tree.Root == nil{
		log.Fatal("get right subtee on empty tree")
		return nil
	}
	return tree.Root.Right
}

func get_data(tree *Binary_Tree) int{
	return tree.Root.Item
}



func Insert(tree * Binary_Tree ,item int) bool{
	increase = false
	return Insert_2(&tree.Root,item,&increase)
}

func Insert_2(local_root ** BTNode, item int,increase *bool) bool{
	if *local_root == nil{
		*local_root = New_BTNode(item)
		*increase = true
		return true
	} else {
		if(item < (*local_root).Item){
			return_value :=  Insert_2(&(*local_root).Left,item,increase) //RETORNA UN BOOL

			if *increase { //si hubo una insercion 
				switch (*local_root).Balance {
				case BALANCED:
					// LEFT-HEAVY
					(*local_root).Balance = LEFT_HEAVY
				case RIGHT_HEAVY:
					//RIGH-HEAVY
					(*local_root).Balance = BALANCED
					*increase = false
					break
				case LEFT_HEAVY:
					//ES CRITICO PORQUE TIENE UN LEFT-LEFT POR LO TANTO LA RAIZ ES -2
					Rebalance_left(&*local_root)
					*increase = false
					break 
				}
			}
			return return_value
		} else if(item > (*local_root).Item){
			return_value :=  Insert_2(&(*local_root).Right,item,increase) //RETORNA UN BOOL

			if *increase { //si hubo una insercion 
				switch (*local_root).Balance {
				case BALANCED:
					// RIGHT-HEAVY
					(*local_root).Balance = RIGHT_HEAVY
				case LEFT_HEAVY:
					//LEFT-HEAVY
					(*local_root).Balance = BALANCED
					*increase = false
					break
				case RIGHT_HEAVY:
					//ES CRITICO PORQUE TIENE UN RIGHT-RIGHT POR LO TANTO LA RAIZ ES -2
					Rebalance_right(&*local_root)
					*increase = false
					break 
				}
			}
			return return_value
		} else {
			return false;
		}
	}
}

func Find( tree *Binary_Tree, item int) int{
	return Find_2(tree.Root, item)
}

func Find_2(local_root *BTNode, item int) int{
	if local_root == nil {
		return -1
	}
	if item < local_root.Item {
		return Find_2(local_root.Left,item)
	} else if item > local_root.Item {
		return Find_2(local_root.Right,item)
	} else {
		return local_root.Item
	}
}

func PreOrden(tree *Binary_Tree){
	PreOrden_2(tree.Root)
}

func PreOrden_2(root *BTNode){
	if root != nil {
		fmt.Println(root.Item)
		PreOrden_2(root.Left)
		PreOrden_2(root.Right)
	}
}


func InOrden(tree *Binary_Tree){
	InOrden_2(tree.Root)
}

func InOrden_2(root *BTNode){
	if root != nil {
		
		InOrden_2(root.Left)
		fmt.Println(root.Item)
		InOrden_2(root.Right)

	}
}

func Rebalance_left(local_root ** BTNode){
	left_child := (*local_root).Left

	if left_child.Balance == RIGHT_HEAVY{ // CASO LEFT-RIGHT
		//OBTENGO UNA REFERENCIA DE LEFT-RIGHT CHILD
		left_right_child := left_child.Right
		//se ajustan los nuevos valances despues de haber realizado la rotacion
		if left_right_child.Balance == LEFT_HEAVY{
			left_child.Balance = BALANCED
			left_right_child.Balance = BALANCED
			(*local_root).Balance = RIGHT_HEAVY
		} else if left_right_child.Balance == BALANCED{
			left_child.Balance = BALANCED
			left_right_child.Balance = BALANCED
			(*local_root).Balance = BALANCED
		} else {
			left_child.Balance = LEFT_HEAVY
			left_right_child.Balance = BALANCED
			(*local_root).Balance = BALANCED
		}
		// REALIZO LEFT ROTATION
		Rotate_Left(&(*local_root).Left)
	} else { // CASO LEFT-LEFT
		left_child.Balance = BALANCED
		(*local_root).Balance = BALANCED
	}
	//REALIZO RIGHT ROTATION
	Rotate_Right(&*local_root)
}


func Rebalance_right(local_root ** BTNode){
	right_child := (*local_root).Right

	if right_child.Balance == LEFT_HEAVY{ // CASO RIGH-LEFT
		//OBTENGO UNA REFERENCIA DE LEFT-RIGHT CHILD
		right_left_child := right_child.Left
		//se ajustan los nuevos valances despues de haber realizado la rotacion
		if right_left_child.Balance == RIGHT_HEAVY{
			right_child.Balance = BALANCED
			right_left_child.Balance = BALANCED
			(*local_root).Balance = LEFT_HEAVY
		} else if right_left_child.Balance == BALANCED{
			right_child.Balance = BALANCED
			right_left_child.Balance = BALANCED
			(*local_root).Balance = BALANCED
		} else {
			right_child.Balance = RIGHT_HEAVY
			right_left_child.Balance = BALANCED
			(*local_root).Balance = BALANCED
		}
		// REALIZO RIGH ROTATION
		Rotate_Right(&(*local_root).Right)
	} else { // CASO RIGHT-RIGHT
		right_child.Balance = BALANCED
		(*local_root).Balance = BALANCED
	}
	//REALIZO RIGHT ROTATION
	Rotate_Left(&*local_root)
}

func Rotate_Right(local_root ** BTNode){
	tmp := (*local_root).Left
	(*local_root).Left = tmp.Right
	tmp.Right = *local_root
	*local_root = tmp
}

func Rotate_Left(local_root ** BTNode){
	tmp := (*local_root).Right
	(*local_root).Right = tmp.Left
	tmp.Left = *local_root
	*local_root = tmp
}