package lists

import (
	"fmt"
	"strings"

	"github.com/ArtLkv/struct_mirea/pkg/utils"
)

// Cтруктура, которая хранит в себе состояние LinkedList
type LinkedList struct {
	length int
	head   *LLNode
}

// Cтруктура, которая хранит в себе состояние элемента LinkedList
type LLNode struct {
	value    int
	nextNode *LLNode
}

// Метод добавления LLNode
func (ll *LinkedList) Push(value int) {
	ll.length += 1
	temp := ll.newNode(value, ll.head)
	ll.head = temp
}

// Метод поиска вхождений
func (ll *LinkedList) IsFound(value int) bool {
	seq := ll.generateSequence()
	for _, el := range strings.Split(seq, "->") {
		element := utils.ConvertStringToInt(el)
		if element == value {
			return true
		}
	}
	return false
}

// Метод вывода LinkedList
func (ll *LinkedList) PrintSequence() {
	seq := ll.generateSequence()
	fmt.Print(seq)
}

// Метод слияния всех узлов в один LinkedList
func (ll *LinkedList) generateSequence() string {
	seq := ""
	if ll.head != nil {
		seq = seq + fmt.Sprintf("%v->", ll.head.value)
		if ll.head.nextNode != nil {
			node := ll.head.nextNode
			for i := 0; i < ll.length; i++ {
				if i+1 != ll.length-1 {
					seq = seq + fmt.Sprintf("%v", node.value)
					if i+1 < ll.length {
						seq = seq + "->"
					}
					node = node.nextNode
				}
			}
		}
	}

	return seq
}

// Метод, которые олицетворяет собой конструктор узла
func (ll *LinkedList) newNode(value int, next *LLNode) *LLNode {
	return &LLNode{
		nextNode: next,
		value:    value,
	}
}

// Метод, которые олицетворяет собой конструктор LinkedList
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}
