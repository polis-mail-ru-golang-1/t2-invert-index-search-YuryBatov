package reversindex

type Files struct {
	Name     string
	Quantity int
}
type Allfiles []Files

func sort1(m map[string]Allfiles, str string) map[string]Allfiles {
	for i := range m[str] {
		for j := i; j > 0 && m[str][j-1].Quantity < m[str][j].Quantity; j-- {
			m[str][j-1], m[str][j] = m[str][j], m[str][j-1]
		}
	}
	return m
}
func ReversIndex(str []string, filesname string, m map[string]Allfiles) map[string]Allfiles {
	var data1 Files
	data1.Name = filesname
	data1.Quantity = 1
	for i := range str {
		k := 0
		if m[str[i]] != nil {
			for j := range m[str[i]] {
				if m[str[i]][j].Name == filesname {
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
func Search(words []string, m map[string]Allfiles, m1 Allfiles) Allfiles {
	var m11 Files
	for i := range words {
		l := 0
		for j := range m[words[i]] {
			for k := range m1 {
				if m[words[i]][j].Name == m1[k].Name {
					m1[k].Quantity = m1[k].Quantity + m[words[i]][j].Quantity
					l++
				}
			}
			if l == 0 {
				m11.Name = m[words[i]][j].Name
				m11.Quantity = m[words[i]][j].Quantity
				m1 = append(m1, m11)
			}
			l = 0
		}
	}
	return m1
}
