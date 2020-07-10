package src

import "fmt"

func solution(n int, array []int) []int {
	sortedArray := MergeSort(array)
	result := make([]int, len(sortedArray))
	start, end := 0, 0
	if n%2 == 0 {
		start = n - 1
	} else {
		start = n - 2
	}
	for i := 0; i < n; i++ {
		if i <= int(n/2) {
			result[i] = sortedArray[end]
			end = end + 2
		} else {
			result[i] = sortedArray[start]
			start = start - 2
		}
		fmt.Print(result)

	}

	return result
}

func merge(arr1, arr2 []int) []int {
	size, i, j := len(arr1)+len(arr2), 0, 0
	result := make([]int, size)
	for k := 0; k < size; k++ {
		switch true {
		case i == len(arr1):
			result[k] = arr2[j]
			j += 1
		case j == len(arr2):
			result[k] = arr1[i]
			i += 1
		case arr1[i] > arr2[j]:
			result[k] = arr2[j]
			j += 1
		case arr1[i] < arr2[j]:
			result[k] = arr1[i]
			i += 1
		case arr1[i] == arr2[j]:
			result[k] = arr2[j]
			j += 1
		}
	}
	return result
}
func MergeSort(numbers []int) []int {
	if len(numbers) < 2 {
		return numbers
	}
	middle := int(len(numbers) / 2)
	return merge(MergeSort(numbers[middle:]), MergeSort(numbers[:middle]))
}
