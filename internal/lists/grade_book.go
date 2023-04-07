package lists

import (
	"fmt"
	"strings"
)

type GradeBook struct {
	dll *DoubleLinkedList
}

// Метод, который связывает ключ с GradeBook
func (gb *GradeBook) Assign(key string) {
	isDone := false
	node := gb.dll.head
	gb.dll.length += 1
	gb.add(node, key, &isDone)
	if node != nil {
		if node.nextNode != nil {
			for i := 0; i < gb.dll.length; i++ {
				if isDone {
					break
				}
				gb.add(node.nextNode, key, &isDone)
				if node.nextNode != nil {
					node = node.nextNode
				}
			}
		}
	}

}
func (gb *GradeBook) DeleteByGroupName()             {}
func (gb *GradeBook) CreateNewDLLWithNegativeMarks() {}

func NewGradeBook(dll *DoubleLinkedList) *GradeBook {
	return &GradeBook{
		dll: dll,
	}
}

func (gb *GradeBook) add(node *DoubleLLNode, key string, isDone *bool) {
	if node != nil {
		val := strings.ReplaceAll(strings.ReplaceAll(node.value, "[", ""), "]", "")
		el := strings.Split(val, ", ")
		isContains := strings.Contains(gb.dll.generateSequence(), fmt.Sprintf("[%v]", key))

		// Если такой ключ есть, то вставить перед первым узлом с таким же ключем
		if el[0] == key && isContains {
			if node.previousNode == nil {
				newObj := gb.dll.newNode(gb.createValueByKey(key), node, nil)
				node.previousNode = newObj
				gb.dll.head = newObj
			} else {
				newObj := gb.dll.newNode(gb.createValueByKey(key), node, node.previousNode)
				node.previousNode.nextNode = newObj
				node.previousNode = newObj
			}
			*isDone = true
			// Если такого ключа нет то, вставьте перед первым узлом у которого ключ больше
		} else if el[0] > key && !isContains {
			if node.previousNode == nil {
				newObj := gb.dll.newNode(gb.createValueByKey(key), node, nil)
				node.previousNode = newObj
				gb.dll.head = newObj
			} else {
				newObj := gb.dll.newNode(gb.createValueByKey(key), node, node.previousNode)
				node.previousNode.nextNode = newObj
				node.previousNode = newObj
			}
			*isDone = true
		}
	}
}

func (gb *GradeBook) createValueByKey(key string) string {
	var (
		group int
		mark  int
	)
	fmt.Print("\nВведите группу: ")
	fmt.Scan(&group)
	fmt.Print("\nВведите оценку: ")
	fmt.Scan(&mark)
	return fmt.Sprintf("[[%s], [%v, %v]]", key, group, mark)
}
