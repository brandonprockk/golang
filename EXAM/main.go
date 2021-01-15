package main

import (
	"io"
	"net/http"
	"log"
)


type Task struct {
        Description string
       	Done bool
}

type List struct {
	ID int
	Task string
}
var task []Task

func list(rw http.ResponseWriter, _ *http.Request) {
        //io.WriteString(w, "Fonction List\n")
	for i,val := range task{
		if val.Done == false {
			list = append(list, List{strconv.Itoa(id),val.Description})
		}
	rw.WriteHeader(http.StatusOK)

}

func done(rw http.ResponseWriter, _ *http.Request) {
        io.WriteString(w, "Fonction Done\n")
}

func add(rw http.ResponseWriter, _ *http.Request) {
        io.WriteString(w, "Fonction Add\n")
}

func main() {

	http.HandleFunc("/", list)
	http.HandleFunc("/done", done)
	http.HandleFunc("/add", add)

	log.Printf("About to listen on 8080, Go to http://127.0.0.0:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
