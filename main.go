package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Element struct {
	Name    string
	Surname string
	Id      string
}

var Data = make(map[string]Element)

func LookUp(k string) *Element {
	if _, ok := Data[k]; ok {
		n := Data[k]
		return &n
	}
	return nil
}

func Add(k string, n Element) bool {
	if k == "" {
		return false
	}

	if LookUp(k) == nil {
		Data[k] = n
		return true
	}
	return false
}

func Delete(k string) bool {
	if LookUp(k) != nil {
		delete(Data, k)
		return true
	}
	return false
}

func Change(k string, n Element) bool {
	Data[k] = n
	return true
}

func Print() {
	for key, data := range Data {
		fmt.Printf("Key: %s;\t Value: %v;\n", key, data)
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		text := sc.Text()
		text = strings.TrimSpace(text)
		tokens := strings.Fields(text)

		switch len(tokens) {
		case 0:
			continue
		case 1:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 2:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 3:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 4:
			tokens = append(tokens, "")
		}

		switch strings.ToLower(tokens[0]) {
		case "print":
			Print()
		case "stop":
			return
		case "delete":
			if !Delete(tokens[1]) {
				fmt.Println("Delete operation failed!")
			}
		case "add":
			n := Element{tokens[2], tokens[3], tokens[4]}
			if !Add(tokens[1], n) {
				fmt.Println("Add operation failed!")
			}
		case "lookup":
			n := LookUp(tokens[1])
			if n != nil {
				fmt.Printf("%v\n", *n)
			}
		case "change":
			n := Element{tokens[2], tokens[3], tokens[4]}
			if !Change(tokens[1], n) {
				fmt.Println("Update operation failed!")
			}
		default:
			fmt.Println("Unknown command - please try again!")
		}
	}
}
