package lists

import (
	"fmt"
)

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
	// [Номер зачетной книжки], [Группа, оценка]
	A := NewGradeBook(NewDoubleLinkedList())

	// Добавляем стартовые значения
	A.dll.Push("[[1], [31, 5]]")
	A.dll.Push("[[2], [27, 2]]")
	A.dll.Push("[[5], [1, 3]]")
	A.dll.Push("[[4], [15, 5]]")
	fmt.Printf("\n'A' изначальный: %v", A.dll.GenerateSequence())

	A.PushById("0")
	fmt.Printf("\n'A' с добавлением по ключу 0: %v", A.dll.GenerateSequence())
	A.PushById("4")
	fmt.Printf("\n'A' с добавлением по ключу 4: %v", A.dll.GenerateSequence())
	A.PushById("8")
	fmt.Printf("\n'A' с добавлением по ключу 8: %v", A.dll.GenerateSequence())
	fmt.Println()
	A.DeleteByGroup("15")
	fmt.Printf("\n'A' после удаления по 15 номеру группы: %v", A.dll.GenerateSequence())

	B := NewGradeBook(NewDoubleLinkedList())
	B.CreateNewWithNegativeMarks(A, B)
	A.DeleteByMark("2")
	fmt.Println()
	fmt.Printf("\n'A' после удаления оценок 'неуд': %v", A.dll.GenerateSequence())
	fmt.Printf("\n'B' список с оценками 'неуд': %v", B.dll.GenerateSequence())
}

// Код ниже - проверка работоспособности структуры DoubleLinkedList(не относится к реализации задачи)
func testDoubleLinkedList() {
	first_dll := NewDoubleLinkedList()
	// Проверяем, что список корректно генерируется
	first_dll.Push("1")
	first_dll.Push("2")
	first_dll.Push("3")
	first_dll.Push("4")
	first_dll.Push("5")
	outputDoubleLinkedList(first_dll, 1)
	// Проверяем, что поиск работает и корректно выводит указатель на элемент
	findValue(first_dll, "4", 1)
	findValue(first_dll, "3", 1)
}

func outputDoubleLinkedList(dll *DoubleLinkedList, listNumber int) {
	fmt.Printf("\n(%v)Список с прямым направлением: %v", listNumber, dll.GenerateSequence())
	dll.ChangeDirection()
	fmt.Printf("\n(%v)Список с обратным направлением: %v", listNumber, dll.GenerateSequence())
}

func findValue(dll *DoubleLinkedList, value string, listNumber int) {
	fmt.Printf("\n\n(%v)Указатель на искомый элемент[со значением %v]: %p", listNumber, value, dll.Find(value))
	fmt.Printf("\n(%v)Развернутый указатель на искомый элемент[со значением %v]: %+v", listNumber, value, dll.Find(value))
}
