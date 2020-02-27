package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func prefix (arr []string, c chan string) {
	result := ""

	if len(arr) == 0{
		c <- result
	}

	if len(arr) == 1{
		c <- arr[0]
	}

	sort.Strings(arr) // sort the strings

	str1 := arr[0] // get the first string of arr
	str2 := arr[len(arr)-1] // last string of arr

	// loop and compare common characters and append to result
	for i := 0; i < len(str1); i++ {
		if str1[i] == str2[i] {
			result += string(str1[i])
		} else {
			break
		}
	}

	c <- result
}

func main() {
	var numset int // number of sets
	var chs []chan string // a channel for a set

	for {
		fmt.Println("Hello, welcome to finding the longest common prefix!")
		fmt.Printf("Enter how many sets of words you wish to find or non-number to terminate: ")

		fmt.Scanf("%d", &numset)
		if numset == 0 {
			os.Exit(0)
		}
		fmt.Println("")

		ptr := []*[]string{} // array of pointers to string arrays

		for i := 0; i < numset; i++ {
			fmt.Printf("SET %d - Enter strings seperated by space and press enter when done: \n", i+1)
			reader := bufio.NewReader(os.Stdin)
			set, _ := reader.ReadString('\n')
			set = strings.TrimSuffix(set, "\n")
			setarr := strings.Split(set, " ")
			ptr = append(ptr, &setarr)
		}

		fmt.Printf("\nLongest common prefix results\n")
		fmt.Println("-------------------------------")

		for i := 0; i < numset; i++ {
			c := make(chan string)
			chs = append(chs, c)
			go prefix(*ptr[i], c)
		}

		for i, v := range chs {
			fmt.Printf("SET %d - %s\n%s\n", i+1, *ptr[i], <-v)
		}
		fmt.Println("")
		chs, numset = nil, 0
	}
}
