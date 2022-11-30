package algorithm

import "fmt"

func Sorting(input []int) {
	for current := 1; current < len(input); current++ {
		i := current
		for i > 0 && input[i-1] > input[i] {
			swap(i-1, i, input)
			i = i - 1
		}
	}

	fmt.Printf("%v", input)

}

func swap(i1 int, i2 int, arr []int) {
	aux := arr[i2]
	arr[i2] = arr[i1]
	arr[i1] = aux
}
