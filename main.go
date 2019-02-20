package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("\n./logCount log-path url-regex limit \n\nor check\n\n https://github.com/maurodelazeri/logCount/blob/master/README.md\n\n")
		os.Exit(1)
	}
	limit, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatalf("limit convertion: %s", err)
	}
	lines, err := readFile(os.Args[1])
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	m := make(map[string]int)
	for _, text := range lines {
		if strings.Contains(text, os.Args[2]) {
			data := strings.Split(text, " ")
			action := data[5][1:]
			m[data[6]+":"+action]++
		}
	}
	n := map[int][]string{}
	var a []int
	for k, v := range m {
		n[v] = append(n[v], k)
	}
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	total := 0
	for _, k := range a {
		for _, s := range n[k] {
			if total > limit {
				return
			}
			values := strings.Split(s, ":")
			fmt.Printf("%s %s %d\n", values[0], values[1], k)
			total++
		}
	}
}
