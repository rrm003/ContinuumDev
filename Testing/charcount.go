package test

func CharCount(s string) int {
	count := 0
	for _, i := range s {
		if i != ' ' {
			count++
			//fmt.Println(string(i), "\t", count)
		}
	}
	return count
}
