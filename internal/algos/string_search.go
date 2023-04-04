package algos

import (
	"fmt"
	"time"

	"github.com/ArtLkv/struct_mirea/pkg/utils"
)

func RunStringSearch() {
	var fileSize int
	var text string
	var word string
	fmt.Print("Файл какой длины вы хотите использовать(100, 1000, 10000): ")
	fmt.Scan(&fileSize)
	switch fileSize {
	case 100:
		text = utils.ReadTextFromFile("textOne.txt")
		word = "Donec"
	case 1000:
		text = utils.ReadTextFromFile("textTwo.txt")
		word = "nisl"
	case 10000:
		text = utils.ReadTextFromFile("textThree.txt")
		word = "Morbi"
	}
	defer utils.TaskTimer(time.Now(), "stringSearchSmallFile")
	index := stringSearch(text, word)
	utils.CheckFoundIndex(index)
	fmt.Printf("\nСф + Мф = %v", sum)

}

func stringSearch(text string, word string) int {
	size := len(text)
	for i := 0; i < size; i++ {
		if text[i] == word[0] {
			sum++
			j := 0
			for j < len(word) && text[i+j] == word[j] {
				sum++
				j++
			}
			if len(word) == j {
				sum++
				return i
			}
			sum++
		}
		sum++
	}
	return -1
}
