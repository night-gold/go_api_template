package server

import(
	"net/http"
)

// Healthz is here to return a response from load balancer and kubernetes readiness and liveliness probe
func Healthz(w http.ResponseWriter, r *http.Request){
	message := "Ok!"
	w.Write([]byte(message))
}