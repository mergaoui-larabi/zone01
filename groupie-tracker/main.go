package main

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/model"
	"net/http"
)

const URL = "https://groupietrackers.herokuapp.com/api/artists/1"

type htmlComponent []string

func main() {

	howa, _ := http.Get(URL) //get the response
	// resp, _ := io.ReadAll(howa.Body) // parse the &r.Body

	getArtist := model.Artist{}

	decoder := json.NewDecoder(howa.Body)

	decoder.DisallowUnknownFields()
	decoder.Decode(&getArtist)

	h, err := json.MarshalIndent(howa, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(h))
	// fmt.Println(string(resp))
	// json.Unmarshal(resp, &getArtist)
	fmt.Printf("howa : %+v \n", getArtist)
	// os.Exit(0)
	mux := http.NewServeMux()

	// mux.HandleFunc("/", indexHandler)
	// mux.HandleFunc("/", fetchHandler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println(err.Error())
	}
}

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	temp, _ := template.ParseFiles("./index.html")

// 	temp.Execute()
// }
