package handler

import (
	"net/http"
)

func Charge(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("charge"))
}
