package lists

import (
	"fmt"
)

// Cтруктура, которая хранит в себе состояние DoubleLinkedList
type DoubleLinkedList struct {
	length       int
	direction    int
	firstElement *DoubleLLNode
	lastElement  *DoubleLLNode
}

// Cтруктура, которая хранит в себе состояние элемента DoubleLinkedList
type DoubleLLNode struct {
	value        string
	nextNode     *DoubleLLNode
	previousNode *DoubleLLNode
}

// Метод добавления DoubleLLNode
func (dll *DoubleLinkedList) Push(value string) {
	dll.length += 1
	if dll.firstElement == nil {
		dll.firstElement = dll.newNode(value, nil, nil)
	} else if dll.lastElement == nil {
		dll.lastElement = dll.newNode(value, nil, dll.firstElement)
		dll.firstElement.nextNode = dll.lastElement
	} else {
		dll.lastElement = dll.newNode(value, nil, dll.lastElement)
		dll.lastElement.previousNode.nextNode = dll.lastElement
	}
}

// Метод удаления DoubleLLNode
func (dll *DoubleLinkedList) Remove(node *DoubleLLNode, index int, initialLength int) {
	if index == initialLength-1 {
		dll.lastElement = node.previousNode
		node.previousNode.nextNode = nil
	} else if index > 0 && index < dll.length {
		node.nextNode.previousNode = node.previousNode
		node.previousNode.nextNode = node.nextNode
	} else if index == 0 {
		dll.firstElement = node.nextNode
		node.nextNode.previousNode = nil
	}
	node = nil // Удаление ссылки на DoubleLLNode
}

// Метод смены направления DoubleLinkedList
func (dll *DoubleLinkedList) ChangeDirection() {
	if dll.direction == 1 {
		dll.direction = 0
	} else {
		dll.direction = 1
	}
}

// Метод поиска DoubleLLNode, если DoubleLLNode найден то возвращается ссылка на него
func (dll *DoubleLinkedList) Find(value string) *DoubleLLNode {
	if dll.firstElement != nil {
		if dll.firstElement.value == value {
			return dll.firstElement
		}
		if dll.firstElement.nextNode != nil {
			node := dll.firstElement.nextNode
			for i := 0; i < dll.length; i++ {
				if i+1 != dll.length-1 {
					if node.value == value {
						return node
					}
					node = node.nextNode
				}
			}
		}
	}
	return nil
}

// Метод вывода DoubleLinkedList
func (dll *DoubleLinkedList) PrintSequence() {
	seq := dll.generateSequence()
	fmt.Print(seq)
}

// Метод слияния всех узлов в один DoubleLinkedList
func (dll *DoubleLinkedList) generateSequence() string {
	seq := ""
	if dll.direction == 0 {
		if dll.firstElement != nil {
			seq = seq + dll.firstElement.value
			if dll.firstElement.nextNode != nil {
				seq = seq + "->"
				node := dll.firstElement.nextNode
				for i := 0; i < dll.length; i++ {
					if i+1 != dll.length-1 {
						seq = seq + node.value
						if i+1 < dll.length {
							seq = seq + "->"
						}
						if node.nextNode != nil {
							node = node.nextNode
						}
					}
				}
			}
		}
	} else {
		if dll.lastElement != nil {
			seq = seq + dll.lastElement.value
			if dll.lastElement.previousNode != nil {
				seq = seq + "<-"
				node := dll.lastElement.previousNode
				for i := 0; i < dll.length; i++ {
					if i+1 != dll.length-1 {
						seq = seq + node.value
						if i+1 < dll.length {
							seq = seq + "<-"
						}
						if node.previousNode != nil {
							node = node.previousNode
						}
					}
				}
			}
		}
	}

	return seq
}

// Метод, которые олицетворяет собой конструктор DoubleLinkedList
func (dll *DoubleLinkedList) newNode(value string, next *DoubleLLNode, previous *DoubleLLNode) *DoubleLLNode {
	return &DoubleLLNode{
		previousNode: previous,
		nextNode:     next,
		value:        value,
	}
}

// Метод, которые олицетворяет собой конструктор DoubleLinkedList
func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{}
}
