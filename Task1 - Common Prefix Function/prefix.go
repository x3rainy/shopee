package main

import (
	"fmt"
	"sort"
)

func prefix (arr []string) string {
	result := ""

	if len(arr) == 0{
		return result
	}
	if len(arr) == 1{
		return arr[0]
	}

	sort.Strings(arr) // sort the strings

	str1 := arr[0] // get the first string of arr
	str2 := arr[len(arr)-1] // last string of arr

	// loop and compare common characters and append to result
	for i := 0; i < len(str1); i++ {
		if str1[i] == str2[i]{
			result += string(str1[i])
		} else {
			break
		}
	}

	return result
}

func main() {
	arr1 := []string{"flower","flow","flight"}
    arr2 := []string{"dog","racecar","car"}
    arr3 := []string{"thanos", "titanic", "thor", "trim"}
	arr4 := []string{"thanatos"}
	arr5 := []string{}
	arr6 := []string{"torch","torches","torchic"}

	fmt.Println(prefix(arr1))
	fmt.Println(prefix(arr2))
	fmt.Println(prefix(arr3))
	fmt.Println(prefix(arr4))
	fmt.Println(prefix(arr5))
	fmt.Println(prefix(arr6))
}
