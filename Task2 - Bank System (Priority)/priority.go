package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func scanLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func writeFile(path string, m map[string][]string, user string) error {
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	read, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	defer file.Close()

	lines := strings.Split(string(read), "\n")

	w := bufio.NewWriter(file)

	for i, line := range lines {
		update := strings.Split(line, " ")
		s := strings.Split(user, " ")
		if update[0] == s[0] {
			lines[i] = s[0] + " " + m[s[0]][0] + " " + m[s[0]][1]
		}
	}
	override := strings.Join(lines, "\n")
	err = ioutil.WriteFile(path, []byte(override), 0644)
	if err != nil {
		return err
	}

	return w.Flush()
}

func main(){
	m := make(map[string][]string)

	data, err := scanLines("system.txt")
	if err != nil {
		panic(err)
	}

	for _, format := range data {
		s := strings.Split(format, " ")
		m[s[0]] = append(m[s[0]], s[1])
		m[s[0]] = append(m[s[0]], s[2])
	}

	input := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	hi := make(chan []string)
	lo := make(chan []string)

	for i := 0; i < len(input)/2; i++ {
		go checkBalance(m, input[i], hi)
	}

	for i := 4; i < len(input); i++ {
		go checkBalance(m, input[i], lo)
	}

	for i := 0; i < 500; i++ {
		select {
		case v := <-hi:
			output(v)
			default:
				select{
					case v := <-hi:
						output(v)
					case v := <-lo:
						output(v)
						default:
				}
		}
	}
}

func output(v []string) {
	fmt.Printf("%s has $%s\n", v[0], v[1])
}

func checkBalance(m map[string][]string, user string, c chan <- []string) {
	c <- m[user]
	//fmt.Printf("%s has $%s\n", user, m[user][1])
}
