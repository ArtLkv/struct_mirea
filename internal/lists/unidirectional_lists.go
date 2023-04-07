package lists

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/ArtLkv/struct_mirea/pkg/utils"
)

func RunUnidirectionalLists() {
	first_ll := NewLinkedList()
	second_ll := NewLinkedList()
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < 4; i++ {
		first_ll.Push(rnd.Intn(100))
		second_ll.Push(rnd.Intn(100))
	}

	fmt.Print("Первый список: ")
	first_ll.PrintSequence()
	fmt.Print("\nВторой список: ")
	second_ll.PrintSequence()
	// Проверка того добавляется ли корретно элемент в самое начало
	first_ll.Push(101)
	first_ll.Push(102)
	fmt.Print("\nПервый список: ")
	first_ll.PrintSequence()

	// Создание итогового списка (По варианту 1)
	result_ll := NewLinkedList()
	for _, v := range strings.Split(first_ll.generateSequence(), "->") {
		value := utils.ConvertStringToInt(v)
		if !result_ll.IsFound(value) {
			result_ll.Push(value)
		}
	}

	for _, v := range strings.Split(second_ll.generateSequence(), "->") {
		value := utils.ConvertStringToInt(v)
		if !result_ll.IsFound(value) {
			result_ll.Push(value)
		}
	}

	// Вывод итогового списка
	fmt.Print("\nИтоговый список: ")
	result_ll.PrintSequence()
}
