package main

import (
	"fmt"
	"nc/folder"
)

func main() {
	folder.SetUp()
	tcp := folder.Newserver("8989")
	err := tcp.ListenToConn()
	if err != nil {
		fmt.Println(err)
	}
}
