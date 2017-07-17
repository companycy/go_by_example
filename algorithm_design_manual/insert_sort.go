package main

import "fmt"

// abc
func swap(i, j *int) {
	k := *i
	*i = *j
	*j = k
}


func insert_sort(s []int) {
	for i := 0; i < len(s); i++ {
		for j := i; j > 0 && s[j] < s[j-1]; j-- {
			swap(&s[j], &s[j-1])
		}
	}
}

func main() {
	s := []int{1, 3, 2}
	insert_sort(s)
	fmt.Println(s)
}
