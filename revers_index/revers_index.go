package revers_index

import "../files"

func sort1(m map[string]files.Allfiles, str string) map[string]files.Allfiles {
	for i := range m[str] {
		for j := i; j > 0 && m[str][j-1].Quantity < m[str][j].Quantity; j-- {
			m[str][j-1], m[str][j] = m[str][j], m[str][j-1]
		}
	}
	return m
}

func Revers_index(str []string, files_name string, m map[string]files.Allfiles) map[string]files.Allfiles {
	//var data []files
	var data1 files.Files
	data1.Name = files_name
	data1.Quantity = 1
	//var k int
	for i := range str {
		k := 0
		if m[str[i]] != nil {
			for j := range m[str[i]] {
				if m[str[i]][j].Name == files_name {
					m[str[i]][j].Quantity++
					k++
				}
			}
			if k == 0 {
				m[str[i]] = append(m[str[i]], data1)
			}
		} else {
			m[str[i]] = append(m[str[i]], data1)
		}
		k = 0
		m = sort1(m, str[i])
	}

	return m
}
