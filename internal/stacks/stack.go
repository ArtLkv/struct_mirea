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

// Метод, которые олицетворяет собой конструктор StackNode
func (s *Stack) newNode(value string, nextNode *StackNode, previousNode *StackNode) *StackNode {
	return &StackNode{
		value:        value,
		nextNode:     nextNode,
		previousNode: previousNode,
	}
}

// Метод, который добавляет в стек StackNode
func (s *Stack) Push() {}

// Метод, который удаляет из стека StackNode
func (s *Stack) Pop() {}

// Метод слияния всех StackNode в один Stack
func (s *Stack) GenerateSequence() string {
	seq := ""
	if s.firstElement != nil {
		seq = seq + s.firstElement.value
		if s.firstElement.nextNode != nil {
			seq = seq + "->"
			node := s.firstElement.nextNode
			for i := 0; i < s.length; i++ {
				if i+1 != s.length-1 {
					seq = seq + node.value
					if i+1 < s.length {
						seq = seq + "->"
					}
					if node.nextNode != nil {
						node = node.nextNode
					}
				}
			}
		}
	}

	return seq
}

// Метод, который проверяет не пуст ли список
func (s *Stack) IsEmpty() {}

// Метод, который проверяет не полон ли список
func (s *Stack) IsFull() {}

// Метод, которые олицетворяет собой конструктор Stack
func NewStack() *Stack {
	return &Stack{}
}
