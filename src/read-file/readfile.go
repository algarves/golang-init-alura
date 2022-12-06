package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	sites := readFile()

	for _, site := range sites {
		fmt.Println(site)
	}
}

func readFile() []string {
	var sites []string
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("An error ocurred:", err)
	}

	scanner := bufio.NewReader(file)
	for {
		line, err := scanner.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}
