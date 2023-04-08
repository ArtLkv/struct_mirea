package stacks

type Stack struct {
	length       int
	firstElement *StackNode
	lastElement  *StackNode
}

type StackNode struct {
	value        string
	nextNode     *StackNode
	previousNode *StackNode
}

func (s *Stack) newNode(value string, nextNode *StackNode, previousNode *StackNode) *StackNode {
	return &StackNode{
		value:        value,
		nextNode:     nextNode,
		previousNode: previousNode,
	}
}

func NewStack() *Stack {
	return &Stack{}
}
