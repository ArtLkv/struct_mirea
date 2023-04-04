package algos

import (
	"fmt"
	"time"

	"github.com/ArtLkv/struct_mirea/pkg/utils"
)

func RunMergeSort(slice []int) {
	defer utils.TaskTimer(time.Now(), "mergeSort")
	utils.OutputSlice(true, mergeSort(slice))
	fmt.Printf("\nСф + Мф = %v", sum)
}

func mergeSort(slice []int) []int {
	size := len(slice)

	if size == 1 {
		return slice
	}

	middle := int(size / 2)
	left := make([]int, middle)
	right := make([]int, size-middle)

	for i := 0; i < size; i++ {
		if i < middle {
			left[i] = slice[i]
		} else {
			right[i-middle] = slice[i]
		}
		sum++ // Операция сравнения
	}
	return mergeParts(mergeSort(left), mergeSort(right))
}

func mergeParts(left []int, right []int) []int {
	result := make([]int, len(left)+len(right))
	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		sum++ // Операция перемещения
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		sum++ // Операция перемещения
		i++
	}
	return result
}
