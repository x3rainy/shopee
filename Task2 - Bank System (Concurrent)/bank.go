package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
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

func serviceRequest(m map[string][]string, username string, wg *sync.WaitGroup, mut *sync.Mutex) {
	//var service int

	defer wg.Done()

	// prompt for services
	//fmt.Printf("\nYou have logged in successfully.\n")
	//op:for {
	//	fmt.Println("1 - Check account balance")
	//	fmt.Println("2 - Deposit")
	//	fmt.Println("3 - Withdrawal")
	//	fmt.Printf("\nSelect the service you want: ")
	//	fmt.Scanf("%d", &service)
	//
	//	switch service {
	//	case 1:
	//		queryC := make(chan string)
	//		go query(m, username, queryC)
	//		fmt.Printf("Your account balance: $%s\n", <-queryC)
	//		break op
	//	case 2:
	//		depositC := make(chan string)
	//		go deposit(m, username, depositC, mut)
	//		fmt.Printf("Your account balance: $%s\n", <-depositC)
	//		break op
	//	case 3:
	//		withdrawC := make(chan string)
	//		go withdraw(m, username, withdrawC, mut)
	//		fmt.Printf("Your account balance: $%s\n", <-withdrawC)
	//		break op
	//	default:
	//		fmt.Printf("\nInvalid service, please try again.\n")
	//	}
	//}

	depositC := make(chan string)
	go deposit(m, username, depositC, mut)
	//fmt.Printf("Your account balance: $%s\n", <-depositC)
	<- depositC
}

func query(m map[string][]string, user string, queryC chan string) {
	defer close(queryC)
	queryC <- m[user][1]
}

func deposit(m map[string][]string, user string, depositC chan string, mut *sync.Mutex){
	//var amount int

	//fmt.Print("Enter amount to deposit: $")
	//fmt.Scanf("%d", &amount)

	amount := 2000

	//fmt.Printf("\nDeposit successful!\n")
	go updateAccount(m, user, amount, depositC, mut)
}

func withdraw(m map[string][]string, user string, withdrawC chan string, mut *sync.Mutex){
	var amount int

	fmt.Print("Enter amount to withdraw: $")
	fmt.Scanf("%d", &amount)

	balance, _ := strconv.Atoi(m[user][1])

	if (balance < amount){
		fmt.Printf("\nInsufficient balance!\n")
		withdrawC <- strconv.Itoa(balance)
		return
	}

	fmt.Printf("\nWithdrawal successful!\n")
	go updateAccount(m, user, -amount, withdrawC, mut)
}

func updateAccount(m map[string][]string, user string, amount int, c chan <- string, mut *sync.Mutex) {
	mut.Lock()
	defer close(c)
	balance, _ := strconv.Atoi(m[user][1])
	balance += amount

	m[user][1] = strconv.Itoa(balance)
	fmt.Println("Your account balance: $", m[user][1])
	c <- m[user][1]
	mut.Unlock()
}

func main() {
	m := make(map[string][]string)
	var wg sync.WaitGroup
	var mut sync.Mutex

	fmt.Println("Welcome to bank.")

	// get data from bank system file
	data, err := scanLines("system.txt")
	if err != nil {
		panic(err)
	}

	// put data into map
	for _, format := range data {
		s := strings.Split(format, " ")
		m[s[0]] = append(m[s[0]], s[1])
		m[s[0]] = append(m[s[0]], s[2])
	}

	// prompt for username
	fmt.Print("Please enter your username: ")
	reader := bufio.NewReader(os.Stdin)
	username, _ := reader.ReadString('\n')
	username = strings.TrimSuffix(username, "\n")

	// This section is used to test for multiple service requests
	input := strings.Split(username, " ")

	//This section is used to test multiple service requests
	for _, username := range input {
		wg.Add(1)
		go serviceRequest(m, username, &wg, &mut)
	}
	wg.Wait()

	fmt.Printf("\nThank you. Terminating...\n")

	// overwrite bank system file
	writeFile("system.txt", m, username)
}