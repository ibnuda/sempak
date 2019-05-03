package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// EmbuhHandler ya embuh ya.
func EmbuhHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ukuransempak := vars["ukuran"]
	fmt.Fprintf(w, "cari sempak ukuran %s", ukuransempak)
}
