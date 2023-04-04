package algos

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ArtLkv/struct_mirea/pkg/utils"
)

func RunQuickSort(slice []int) {
	defer utils.TaskTimer(time.Now(), "quickSort")
	quickSort(slice)
	utils.OutputSlice(true, slice)
	fmt.Printf("\nСф + Мф = %v", sum)
}

func quickSort(slice []int) []int {
	size := len(slice)
	if size < 2 {
		return slice
	}
	left := 0
	right := size - 1

	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	pivot := rnd.Int() % size

	slice[pivot], slice[right] = slice[right], slice[pivot]
	sum++ // Операция перемещения

	for index := range slice {
		if slice[index] < slice[right] {
			sum++ // Операция сравнения
			slice[left], slice[index] = slice[index], slice[left]
			sum++ // Операция перемещения
			left++
		}
	}

	slice[left], slice[right] = slice[right], slice[left]
	sum++ // Операция перемещения

	quickSort(slice[:left])
	quickSort(slice[left+1:])

	return slice
}
