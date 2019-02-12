package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main (){
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter your command: ")
		scanner.Scan()
		text := scanner.Text()
		cmd := strings.Split(text, " ")
		switch cmd[0] {
		case "close":
			fmt.Println("cmd close")
			return
		case "list":
			fmt.Println("cmd list")
			getAllBook()
		case "add":
			fmt.Println("cmd add")
			params := strings.Split(text[4:], ";")
			if len(params) != 3{
				fmt.Printf("Invalid parameters!%v %v\n", len(params), params)
				break
			}
			year, err := strconv.ParseInt(strings.Trim(params[2], " "), 10, 0)
			if err != nil {
				fmt.Printf("Error while parse %v\n", err)
			}
			addBook(params[0], params[1], int16(year))
		case "remove":
			fmt.Println("cmd remove")
			id, err := strconv.ParseInt(strings.Trim(cmd[1], " "), 10, 0)
			if err != nil {
				fmt.Printf("Error while parse %v\n", err)
			}
			removeBook(int32(id))
		default:
			fmt.Printf("Not found %v\n", cmd)
		}

	}
}

