package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func check(e error) {
	if e != nil {
		fmt.Print("error occoured while reading from file")
		panic(e)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	var data string
	data = readFile()
	fmt.Fprintf(w, "Congratulations! this is your secret from hashi vault "+data)
}

func isAlivehandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Application is running")
}

func imReadyhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Application is ready")
}

func readFile() string {
	dat, err := ioutil.ReadFile("/mnt/app.properties")
	check(err)
	return string(dat)
}

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/isAlive", isAlivehandler)
	http.HandleFunc("/imReady", imReadyhandler)
	http.ListenAndServe(":3000", nil)
}
