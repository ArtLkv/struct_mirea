package lists

import (
	"strings"
)

// Cтруктура, которая хранит в себе состояние LinkedList
type LinkedList struct {
	length       int
	firstElement *LLNode
}

// Cтруктура, которая хранит в себе состояние элемента LinkedList
type LLNode struct {
	value    string
	nextNode *LLNode
}

// Метод добавления LLNode
func (ll *LinkedList) Push(value string) {
	ll.length += 1
	temp := ll.newNode(value, ll.firstElement)
	ll.firstElement = temp
}

// Метод поиска вхождений
func (ll *LinkedList) IsFound(value string) bool {
	seq := ll.GenerateSequence()
	for _, el := range strings.Split(seq, "->") {
		if el == value {
			return true
		}
	}
	return false
}

// Метод слияния всех LLNode в один LinkedList
func (ll *LinkedList) GenerateSequence() string {
	seq := ""
	if ll.firstElement != nil {
		seq = seq + ll.firstElement.value
		if ll.firstElement.nextNode != nil {
			seq = seq + "->"
			node := ll.firstElement.nextNode
			for i := 0; i < ll.length; i++ {
				if i+1 != ll.length-1 {
					seq = seq + node.value
					if i+1 < ll.length {
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

// Метод, которые олицетворяет собой конструктор узла
func (ll *LinkedList) newNode(value string, next *LLNode) *LLNode {
	return &LLNode{
		nextNode: next,
		value:    value,
	}
}

// Метод, которые олицетворяет собой конструктор LinkedList
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}
