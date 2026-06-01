package structure

import "fmt"

type ListNode struct {
	Value int
	Next *ListNode
}

func prepend(head **ListNode, value int){
	newNode := ListNode{Value: value, Next: *head}
	*head = &newNode
}

func main(){
	var head *ListNode

	prepand(&head, 10)
	prepand(&head, 20)

	current := head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}