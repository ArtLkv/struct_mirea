package lists

import "fmt"

// Cтруктура, которая хранит в себе состояние DoubleLinkedList
type DoubleLinkedList struct {
	length    int
	direction int
	head      *DoubleLLNode
	tail      *DoubleLLNode
}

// Cтруктура, которая хранит в себе состояние элемента DoubleLinkedList
type DoubleLLNode struct {
	value        int
	nextNode     *DoubleLLNode
	previousNode *DoubleLLNode
}

// Метод добавления DoubleLLNode
func (dll *DoubleLinkedList) Push(value int) {
	dll.length += 1
	if dll.head == nil {
		dll.head = dll.newNode(value, nil, nil)
	} else if dll.tail == nil {
		dll.tail = dll.newNode(value, nil, dll.head)
		dll.head.nextNode = dll.tail
	} else {
		dll.tail = dll.newNode(value, nil, dll.tail)
		dll.tail.previousNode.nextNode = dll.tail
	}
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
func (dll *DoubleLinkedList) Find(value int) *DoubleLLNode {
	if dll.head != nil {
		if dll.head.value == value {
			return dll.head
		}
		if dll.head.nextNode != nil {
			node := dll.head.nextNode
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
		if dll.head != nil {
			seq = seq + fmt.Sprintf("%v->", dll.head.value)
			if dll.head.nextNode != nil {
				node := dll.head.nextNode
				for i := 0; i < dll.length; i++ {
					if i+1 != dll.length-1 {
						seq = seq + fmt.Sprintf("%v", node.value)
						if i+1 < dll.length {
							seq = seq + "->"
						}
						node = node.nextNode
					}
				}
			}
		}
	} else {
		if dll.tail != nil {
			seq = seq + fmt.Sprintf("%v<-", dll.tail.value)
			if dll.tail.previousNode != nil {
				node := dll.tail.previousNode
				for i := 0; i < dll.length; i++ {
					if i+1 != dll.length-1 {
						seq = seq + fmt.Sprintf("%v", node.value)
						if i+1 < dll.length {
							seq = seq + "<-"
						}
						node = node.previousNode
					}
				}
			}
		}
	}

	return seq
}

// Метод, которые олицетворяет собой конструктор DoubleLinkedList
func (dll *DoubleLinkedList) newNode(value int, next *DoubleLLNode, previous *DoubleLLNode) *DoubleLLNode {
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
