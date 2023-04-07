package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func TaskTimer(start time.Time, funcName string) {
	duration := time.Since(start)
	fmt.Printf("\nАлгоритм %s выполняется %v\n", funcName, duration.Round(8))
}

func GenerateRandomSlice(size int) []int {
	slice := make([]int, size)
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	for index := range slice {
		slice[index] = rnd.Intn(1000)
	}
	return slice
}
func OutputSlice(isSorted bool, slice []int) {
	if isSorted {
		fmt.Println("--- Sorted ---")
		fmt.Println(slice)
	} else {
		fmt.Println("--- Unsorted ---")
		fmt.Println(slice)
	}
}

func CheckFoundIndex(index int) {
	if index == -1 {
		fmt.Println("Искомого элемента нету в массиве")
	} else {
		fmt.Printf("Искомый элемент был найден в массиве под индексом %v ", index)
	}
}

func Reverse(slice []int) []int {
	size := len(slice)
	reverseSlice := make([]int, size)

	for index, value := range slice {
		j := size - index - 1
		reverseSlice[j] = value
	}

	return reverseSlice
}

func ReadTextFromFile(fileName string) string {
	fullFileName := fmt.Sprintf("text/%s", fileName)
	fmt.Println(fullFileName)
	file, _ := os.Open(fullFileName)
	defer file.Close()
	var b bytes.Buffer
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		b.WriteString(scanner.Text())
	}
	return b.String()
}

func ConvertStringToInt(str string) int {
	value, _ := strconv.Atoi(str)
	return value
}
