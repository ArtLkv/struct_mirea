package lists

import "fmt"

// Будет реализован вариант 1
func RunBidirectionalLists() {
	var task int
	fmt.Print("Хотите проверить корректность работы списка?(1 - да, 2 - нет): ")
	fmt.Scan(&task)
	switch task {
	case 1:
		testDoubleLinkedList()
	case 2:
		runTask()
	default:
		fmt.Println("Вы ввели неверное число")
	}
}

func runTask() {
	// A := NewDoubleLinkedList()
}

// Код ниже - проверка работоспособности структуры DoubleLinkedList(не относится к реализации задачи)
func testDoubleLinkedList() {
	first_dll := NewDoubleLinkedList()
	// Проверяем, что список корректно генерируется
	first_dll.Push(1)
	first_dll.Push(2)
	first_dll.Push(3)
	first_dll.Push(4)
	first_dll.Push(5)
	outputDoubleLinkedList(first_dll, 1)
	// Проверяем, что поиск работает и корректно выводит указатель на элемент
	findValue(first_dll, 4, 1)
	findValue(first_dll, 3, 1)
}

func outputDoubleLinkedList(dll *DoubleLinkedList, listNumber int) {
	fmt.Printf("\n(%v)Список с прямым направлением: ", listNumber)
	dll.PrintSequence()
	dll.ChangeDirection()
	fmt.Printf("\n(%v)Список с обратным направлением: ", listNumber)
	dll.PrintSequence()
}

func findValue(dll *DoubleLinkedList, value int, listNumber int) {
	fmt.Printf("\n\n(%v)Указатель на искомый элемент[со значением %v]: %p", listNumber, value, dll.Find(value))
	fmt.Printf("\n(%v)Развернутый указатель на искомый элемент[со значением %v]: %+v", listNumber, value, dll.Find(value))
}
