package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Prazno ha ha")
	//http.ServeFile(w, r, "crazy-awesom/fe/public/index.html")
}
func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}
