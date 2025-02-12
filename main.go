package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	pq := NewPQ()

	for {
		fmt.Print("letter > ")
		if !scanner.Scan() {
			break
		}

		text := scanner.Text()
		if text == "exit" {
			break
		} else if text == "remove" {
			m := pq.Remove()
			fmt.Printf("max=%d. text=%q\n", m, string(rune(m)))
			continue
		}

		val := text[0]
		pq.Insert(int(val))
	}
	fmt.Println("bye!")
	fmt.Printf("pq=%+v\n", pq)
}
