package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ArtLkv/struct_mirea/internal/algos"
	"github.com/ArtLkv/struct_mirea/pkg/utils"
)

func main() {
	args := os.Args[1:]
	taskNumber, _ := strconv.Atoi(args[0]) // Получаем номер нужной нам задачи
	sliceSize, _ := strconv.Atoi(args[1])  // Получаем размер массива, который мы хотим создать

	slice := utils.GenerateRandomSlice(sliceSize)
	switch taskNumber {
	case 1:
		utils.OutputSlice(false, slice)
		algos.RunSelectionSort(slice)
	case 2:
		utils.OutputSlice(false, slice)
		algos.RunQuickSort(slice)
	case 3:
		// TODO: Check sum
		utils.OutputSlice(false, slice)
		algos.RunMergeSort(slice)
	case 4:
		algos.RunBarrierSearch(slice)
	case 5:
		algos.RunBinarySearch(slice)
	case 6:
		algos.RunStringSearch()
	case 7:
		fmt.Println(7)
	case 8:
		fmt.Println(8)
	}
}
