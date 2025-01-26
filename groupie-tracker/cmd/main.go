package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const URL = "http://groupietrackers.herokuapp.com/api/artists/1"

type artist struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type htmlComponent []string

func (hC *htmlComponent) Write(p []byte) (int, error) {
	*hC = append(*hC, string(p))
	return len(p), nil
}

// var heros []Hero

func main() {
	fmt.Println("howa")
	newar := artist{}
	resp, _ := http.Get(URL)

	err := json.NewDecoder(resp.Body).Decode(&newar)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newar)

	mux := http.NewServeMux()

	// t, _ := template.ParseGlob("../temp/*.html")

	// var buffer htmlComponent

	// for i := range heros {
	// 	t.ExecuteTemplate(&buffer, "li.html", heros[i])
	// }

	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	t.ExecuteTemplate(w, "atag.html", buffer)
	// 	fmt.Fprintf(w, "<h1>d</h1>")
	// })
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}

// func createHero(name string, age uint) {
// 	newhero := Hero{
// 		Name: name,
// 		Age:  age,
// 	}

// 	heros = append(heros, newhero)
// }
