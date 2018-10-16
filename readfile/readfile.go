package readfile

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Readfile(files_name string) []string {
	var str1 []string
	data, err := ioutil.ReadFile(files_name)
	if err != nil {
		fmt.Println("Ошибка чтения файла")
		fmt.Println(err)
	}
	str := string(data)
	str1 = strings.Split(str, " ")
	return str1
}
