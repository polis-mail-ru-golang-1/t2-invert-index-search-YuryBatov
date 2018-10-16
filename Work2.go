package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"./files"
	"./readfile"
	"./revers_index"
)

func sort2(m files.Allfiles) {
	for i := range m {
		for j := i; j > 0 && m[j-1].Quantity < m[j].Quantity; j-- {
			m[j-1], m[j] = m[j], m[j-1]
		}
	}
}

func main() {
	m := make(map[string]files.Allfiles)
	var word string
	var words []string
	files_name := make([]string, 0, 10)
	files_name = os.Args[1:]
	fmt.Printf("\n%s", "Введите слово: ")
	for i := range files_name {
		str1 := readfile.Readfile(files_name[i])
		m = revers_index.Revers_index(str1, files_name[i], m)
	}
	buf := bufio.NewScanner(os.Stdin)
	buf.Scan()
	word = buf.Text()
	words = strings.Split(word, " ")
	var m_1 files.Allfiles
	var m_11 files.Files
	for i := range words {
		l := 0
		for j := range m[words[i]] {
			for k := range m_1 {
				if m[words[i]][j].Name == m_1[k].Name {
					m_1[k].Quantity = m_1[k].Quantity + m[words[i]][j].Quantity
					l++
				}
			}
			if l == 0 {
				m_11.Name = m[words[i]][j].Name
				m_11.Quantity = m[words[i]][j].Quantity
				m_1 = append(m_1, m_11)
			}
			l = 0
		}
	}
	sort2(m_1)
	for i := range m_1 {
		fmt.Printf("%s : %d совпадений\n", m_1[i].Name, m_1[i].Quantity)
	}
	//m_1 = sort.Ints()
}
