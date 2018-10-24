package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"./readfile"
	"./reversindex"
)

func sort2(m reversindex.Allfiles) {
	for i := range m {
		for j := i; j > 0 && m[j-1].Quantity < m[j].Quantity; j-- {
			m[j-1], m[j] = m[j], m[j-1]
		}
	}
}

func main() {
	m := make(map[string]reversindex.Allfiles)
	var word string
	var words []string
	var filesname []string
	filesname = os.Args[1:]
	fmt.Printf("\n%s", "Введите слово: ")
	for i := range filesname {
		str1 := readfile.ReadFile(filesname[i])
		m = reversindex.ReversIndex(str1, filesname[i], m)
	}
	buf := bufio.NewScanner(os.Stdin)
	buf.Scan()
	word = buf.Text()
	words = strings.Split(word, " ")
	var m1 reversindex.Allfiles
	m1 = reversindex.Search(words, m, m1)
	sort2(m1)
	for i := range m1 {
		fmt.Printf("%s : %d совпадений\n", m1[i].Name, m1[i].Quantity)
	}
}
