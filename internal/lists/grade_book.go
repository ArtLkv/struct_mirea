package lists

import (
	"fmt"
	"strings"
)

type GradeBook struct {
	dll *DoubleLinkedList
}

// Метод, который связывает ключ с GradeBook
func (gb *GradeBook) PushById(key string) {
	isDone := false
	node := gb.dll.firstElement
	gb.dll.length += 1
	if node != nil {
		gb.addElementById(node, key, &isDone)
		if node.nextNode != nil {
			for i := 0; i < gb.dll.length; i++ {
				if isDone {
					break
				}
				gb.addElementById(node.nextNode, key, &isDone)
				if node.nextNode != nil {
					node = node.nextNode
				}
			}
		}
	}

}

// Метод, который удаляет все вхождения той или иной группы
func (gb *GradeBook) DeleteByGroup(group string) {
	node := gb.dll.firstElement
	count := 0
	initialLength := gb.dll.length
	if node != nil {
		for i := 0; i < initialLength; i++ {
			val := strings.ReplaceAll(strings.ReplaceAll(node.value, "[", ""), "]", "")
			el := strings.Split(val, ", ")
			if el[1] == group {
				gb.dll.Remove(node, count, initialLength)
				gb.dll.length -= 1
			}
			count += 1
			if node.nextNode != nil {
				node = node.nextNode
			}
		}
	}
}
func (gb *GradeBook) CreateNewDLLWithNegativeMarks() {}

func NewGradeBook(dll *DoubleLinkedList) *GradeBook {
	return &GradeBook{
		dll: dll,
	}
}

// Вспомогательная функция
func (gb *GradeBook) addElementById(node *DoubleLLNode, key string, isDone *bool) {
	if node != nil {
		val := strings.ReplaceAll(strings.ReplaceAll(node.value, "[", ""), "]", "")
		el := strings.Split(val, ", ")
		isContains := strings.Contains(gb.dll.generateSequence(), fmt.Sprintf("[%v]", key))

		// Если такой ключ есть, то вставить перед первым узлом с таким же ключем
		if el[0] == key && isContains {
			if node.previousNode == nil {
				newObj := gb.dll.newNode(gb.createValuesByKey(key), node, nil)
				node.previousNode = newObj
				gb.dll.firstElement = newObj
			} else {
				newObj := gb.dll.newNode(gb.createValuesByKey(key), node, node.previousNode)
				node.previousNode.nextNode = newObj
				node.previousNode = newObj
			}
			*isDone = true
			// Если такого ключа нет то, вставьте перед первым узлом у которого ключ больше
		} else if el[0] > key && !isContains {
			if node.previousNode == nil {
				newObj := gb.dll.newNode(gb.createValuesByKey(key), node, nil)
				node.previousNode = newObj
				gb.dll.firstElement = newObj
			} else {
				newObj := gb.dll.newNode(gb.createValuesByKey(key), node, node.previousNode)
				node.previousNode.nextNode = newObj
				node.previousNode = newObj
			}
			*isDone = true
		}
	}
}

// Вспомогательная функция
func (gb *GradeBook) createValuesByKey(key string) string {
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
