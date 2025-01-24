package asciigenerator

import (
	"fmt"
	"os"
	"strings"
)

func AsciiGenerator(text string, banner string) string {
	charset := CreateMap(banner)
	for _, v := range text {
		if v == 10 || v == 13 {
			continue
		}
		if v < 32 || v > 126 {
			fmt.Println("the input contains an inprintable character")
			os.Exit(0)
		}
	}
	if text == "" {
		os.Exit(0)
	}
	res := ""
	count := 0
	slice_by_line := strings.Split(text, "\r\n")
	fmt.Printf("%#v", slice_by_line)
	for i := range slice_by_line {
		if slice_by_line[i] == "" {
			count++
		}
	}
	if count == len(slice_by_line) {
		for range count - 1 {
			res += "\r\n"
		}
	}
	for i := range slice_by_line {
		if slice_by_line[i] == "" {
			res += "\r\n"
		} else {
			res += Compose(slice_by_line[i], charset)
		}
	}

	return res
}
