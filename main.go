package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf16"
)

func main() {
	var stdout []uint16
	if len(os.Args) == 1 {
		time.AfterFunc(time.Second*3, func() {
			log.Fatalf("Error: no input recieved after 3 seconds.")
		})
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			u, err := strconv.ParseUint(strings.Trim(scanner.Text(), ","), 0, 16)
			if err != nil {
				os.Exit(1)
			}
			stdout = append(stdout, uint16(u))
		}
		fmt.Println(string(utf16.Decode(stdout)))
	}
	if len(os.Args) > 1 {
		in := strings.Join(os.Args[1:], " ")
		if !strings.Contains(in, ",") {
			log.Fatalf("Error: did not contain a proper fromCharCode().  Usage: ./fromchar 116, 101, 115, 116")
		}
		split := strings.Split(in, ",")
		for _, codepoint := range split {
			u, _ := strconv.ParseUint(strings.TrimSpace(codepoint), 0, 16)
			stdout = append(stdout, uint16(u))
		}
		fmt.Println(string(utf16.Decode(stdout)))
	}
	os.Exit(0)
}
