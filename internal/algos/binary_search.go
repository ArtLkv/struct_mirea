package algos

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ArtLkv/struct_mirea/pkg/utils"
)

func RunBinarySearch(slice []int) {
	quickSort(slice)
	utils.OutputSlice(true, slice)
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	target := slice[rnd.Intn(len(slice))]
	defer utils.TaskTimer(time.Now(), "binarySearch")
	index := binarySearch(slice, target)
	utils.CheckFoundIndex(index)
	fmt.Printf("\nСф = %v", sum)
}

func binarySearch(slice []int, target int) int {
	size := len(slice)
	left := 0
	right := size - 1

	for left <= right {
		sum++
		middle := int((left + right) / 2)
		pivot := slice[middle]
		if pivot == target {
			return middle
		}
		if pivot < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
		sum++
	}
	sum++
	return -1
}
