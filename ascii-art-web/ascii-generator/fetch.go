package asciigenerator

import (
	"fmt"
	"os"
	"strings"
)

func CreateMap(banner string) map[rune][]string {
	size_of_byte := 0
	fileName := banner + ".txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(0)
	}
	slice := []string{}
	charset := map[rune][]string{}

	switch banner {
	case "shadow":
		size_of_byte = 7463
		slice = strings.Split(string(file), "\n")
	case "standard":
		size_of_byte = 6623
		slice = strings.Split(string(file), "\n")
	case "thinkertoy":
		size_of_byte = 5558
		slice = strings.Split(string(file), "\r\n")
	}
	if len(file) != size_of_byte { // stand 6623  shadow 7463 thinkertoy 5558 byte
		fmt.Println("hoooooooooooooooowa")
		os.Exit(0)
	}
	for i := ' '; i <= '~'; i++ {
		buffer := []string{}
		index := (int(i)-32)*9 + 1
		for i := index; i < index+8; i++ {
			buffer = append(buffer, slice[i])
		}
		charset[i] = buffer
	}
	// fmt.Println(charset)
	return charset
}
