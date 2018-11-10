package handler

import (
	"net/http"
)

func Reserve(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("reserve"))
}
