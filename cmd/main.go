package main

import (
	"os"

	"github.com/ArtLkv/struct_mirea/internal/algos"
	"github.com/ArtLkv/struct_mirea/internal/lists"
	"github.com/ArtLkv/struct_mirea/pkg/utils"
)

func main() {
	args := os.Args[1:]
	taskNumber := utils.ConvertStringToInt(args[0]) // Получаем номер нужной нам задачи
	sliceSize := utils.ConvertStringToInt(args[1])  // Получаем размер массива, который мы хотим создать

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
		lists.RunUnidirectionalLists()
	case 8:
		lists.RunBidirectionalLists()
	}
}
