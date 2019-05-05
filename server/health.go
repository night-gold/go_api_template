package server

import(
	"net/http"
)

func healthz(w http.ResponseWriter, r *http.Request){
	message := "Ok!"
	w.Write([]byte(message))
}