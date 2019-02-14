package main

import (
	"bufio"
	"fmt"
	"github.com/sidyakina/books_client/sender"
	"os"
	"strconv"
	"strings"
	"time"
)

func main (){
	s, err := sender.Init()
	if err != nil {
		fmt.Printf("Error while connect %v\n", err)
		return
	}
	defer s.Close()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter your command: ")
		scanner.Scan()
		text := scanner.Text()
		cmd := strings.Split(text, " ")
		switch cmd[0] {
		case "help":
			fmt.Println("cmds:")
			fmt.Println("close - for exit (no parameters)")
			fmt.Println("list - for view all books (no parameters)")
			fmt.Println("add - for add book (parameters name, author, year)")
			fmt.Println(" Example: add new book;;author;;1834")
			fmt.Println("add - for remove book (parameter id)")
			fmt.Println(" Example: remove 188809934")
		case "close":
			fmt.Println("cmd close")
			return
		case "list":
			fmt.Println("cmd list")
			s.GetAllBook()
		case "add":
			fmt.Println("cmd add")
			params := strings.Split(strings.Trim(text[3:], " "), ";;")
			if len(params) != 3{
				fmt.Printf("Invalid number parameters!\n")
				break
			}
			year, err := strconv.ParseInt(strings.Trim(params[2], " "), 10, 0)
			if err != nil {
				fmt.Printf("Error while parse: %v\n", err)
				break
			}
			if int(year) > time.Now().Year() || year <= 0{
				fmt.Printf("Invalid parameter for year:%v \n", year)
				break
			}
			if params[0] == "" || params[1] == "" {
				fmt.Printf("Invalid empty parameter \n")
				break
			}
			s.AddBook(params[0], params[1], int16(year))
		case "remove":
			fmt.Println("cmd remove")
			param := strings.Trim(text[6:], " ")
			if len(param) == 0 {
				fmt.Printf("Invalid empty parameter!\n")
				break
			}
			id, err := strconv.ParseInt(param, 10, 0)
			if err != nil {
				fmt.Printf("Error while parse: %v\n", err)
				break
			}
			s.RemoveBook(int32(id))
		default:
			fmt.Printf("Not found %v\n", cmd)
		}

	}
}

