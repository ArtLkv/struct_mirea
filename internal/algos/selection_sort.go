package algos

import (
	"fmt"
	"time"

	"github.com/ArtLkv/struct_mirea/pkg/utils"
)

func RunSelectionSort(slice []int) {
	defer utils.TaskTimer(time.Now(), "selectionSort")
	selectionSort(slice)
	utils.OutputSlice(true, slice)
	fmt.Printf("\nСф + Мф = %v", sum)
}

func selectionSort(slice []int) {
	size := len(slice)
	for i := 0; i < size; i++ {
		middle := i
		for j := i; j < size; j++ {
			if slice[j] < slice[middle] {
				middle = j
				sum++ // Операция сравнения
			}
		}
		slice[i], slice[middle] = slice[middle], slice[i]
		sum++ // Операция перемещения
	}
}
