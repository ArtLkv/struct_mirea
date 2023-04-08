package stacks

import (
	"fmt"

	"github.com/ArtLkv/struct_mirea/pkg/utils"
)

func makeTowersWithStack() int {
	A := NewStack()
	B := NewStack()
	var n int
	fmt.Print("Введите количество сооружений: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		var h string
		fmt.Printf("Введите высоту %v сооружения: ", i)
		fmt.Scan(&h)
		A.Push(h)
	}
	for i := 1; i <= n; i++ {
		B.Push(A.Pop())
	}
	floor := -1
	count := 0
	for i := 1; i <= n; i++ {
		temp := utils.ConvertStringToInt(B.Pop())
		if temp > floor {
			count += 1
			floor = temp
		}
	}
	return count
}

func makeTowersWithoutStack() int {
	floor := -1
	count := 0
	var n int
	fmt.Print("Введите количество сооружений: ")
	fmt.Scan(&n)
	H := []int{}
	for i := 1; i <= n; i++ {
		var h int
		fmt.Printf("Введите высоту %v сооружения: ", i)
		fmt.Scan(&h)
		H = append(H, h)
	}
	for _, v := range H {
		if v > floor {
			count += 1
			floor = v
		}
	}
	// fmt.Printf("Мы можем увидеть %v строений", count)
	return count
}
