package asciigenerator

func Compose(str string, charset map[rune][]string) string {
	res := ""
	for i := 0; i < 8; i++ {
		for _, v := range str {
			// fmt.Println(charset[v][i])s
			res += charset[v][i]
		}
		res += "\r\n"
	}
	return res
}
