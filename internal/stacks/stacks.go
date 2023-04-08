package stacks

import "fmt"

func RunStack() {
	var task int
	fmt.Print("Хотите проверить корректность работы стека?(1 - да, 2 - нет): ")
	fmt.Scan(&task)
	switch task {
	case 1:
		testStruct()
	case 2:
		runTask()
	default:
		fmt.Println("Вы ввели неверное число")
	}
}

func runTask() {
	fmt.Printf("Ответ(со стеком): %v\n", makeTowersWithStack())
	fmt.Printf("Ответ(без стека): %v\n", makeTowersWithoutStack())
}

func testStruct() {
	// Стек на  однонаправленном динамическом списке
	fmt.Println("--- Стек на однонаправленном динамическом списке ---")
	A := NewStack()
	A.Push("1")
	A.Push("3")
	A.Push("-3")
	A.Push("4")
	fmt.Printf("Стек 'А': %v", A.GenerateSequence())
	fmt.Printf("\nЭлемент, который был вытолкнут из стека 'A': %v", A.Pop())
	fmt.Printf("\nСтек 'А': %v", A.GenerateSequence())
	fmt.Println("\n--- Стек на однонаправленном динамическом списке ---")
	fmt.Println()

	fmt.Print("--- Стек на динамическом массиве ---")
	// Стек на динамическом массиве
	B := []int{}
	B = append(B, 2, 9, 10)
	fmt.Print("\nСтек 'B': ")
	outputStructAsSlice(B)
	fmt.Printf("\nЭлемент, который был вытолкнут из стека 'B': %v", B[len(B)-1])
	B = B[:len(B)-1]
	fmt.Print("\nСтек 'B': ")
	outputStructAsSlice(B)
	fmt.Println("\n--- Стек на динамическом массиве ---")
}

func outputStructAsSlice(slice []int) {
	for _, v := range slice {
		if slice[len(slice)-1] == v {
			fmt.Print(v)
		} else {
			fmt.Printf("%v->", v)
		}
	}
}
