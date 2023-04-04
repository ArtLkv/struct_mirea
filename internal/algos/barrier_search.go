package algos

import (
	"fmt"
	"time"

	"github.com/ArtLkv/struct_mirea/pkg/utils"
)

func RunBarrierSearch(slice []int) {
	fmt.Printf("Выберете критерии поиска(рандомный - 1, строго возрастающий - 2, строго убывающий - 3): ")
	var n int
	fmt.Scan(&n)
	switch n {
	case 1:
		utils.OutputSlice(false, slice)
		searchValue = slice[int(len(slice)/2)]
		// getSearchValue()
		sum = 0
		runBarrierSearchWithUnsortedSlice(slice)
	case 2:
		quickSort(slice)
		utils.OutputSlice(true, slice)
		searchValue = slice[int(len(slice)/2)]
		// getSearchValue()
		sum = 0
		runBarrierSearchWithStrongPositiveSortedSlice(slice)
	case 3:
		quickSort(slice)
		reverseSlice := utils.Reverse(slice)
		utils.OutputSlice(true, reverseSlice)
		searchValue = slice[int(len(slice)/2)]
		// getSearchValue()
		sum = 0
		runBarrierSearchWithStrongNegativeSortedSlice(reverseSlice)
	}
	fmt.Printf("\nСф = %v", sum)
}

func runBarrierSearchWithUnsortedSlice(slice []int) {
	defer utils.TaskTimer(time.Now(), "barrierSearchWithUnsortedSlice")
	index := barrierSearch(slice, searchValue)
	utils.CheckFoundIndex(index)
}

func runBarrierSearchWithStrongPositiveSortedSlice(slice []int) {
	defer utils.TaskTimer(time.Now(), "barrierSearchWithStrongPositiveSortedSlice")
	index := barrierSearch(slice, searchValue)
	utils.CheckFoundIndex(index)
}

func runBarrierSearchWithStrongNegativeSortedSlice(slice []int) {
	defer utils.TaskTimer(time.Now(), "barrierSearchWithStrongNegativeSortedSlice")
	index := barrierSearch(slice, searchValue)
	utils.CheckFoundIndex(index)
}

func barrierSearch(slice []int, searchValue int) int {
	size := len(slice)
	last := slice[size-1] // Сохраняем последний элемент
	slice[size-1] = searchValue
	i := 0
	for slice[i] != searchValue {
		sum++ // Операция сравнения
		i++
	}
	slice[size-1] = last // Возвращаем последний элемент на свое место
	if i != (size-1) || searchValue == last {
		sum++
		return i
	}
	return -1
}

// func getSearchValue() {
// 	fmt.Printf("Какой элемент вы хотите найти?: ")
// 	fmt.Scan(&searchValue)
// }
