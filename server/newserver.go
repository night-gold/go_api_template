package server

import(
	"github.com/rs/zerolog"

	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// NewServer is used to start the http server, it also starts 
func NewServer(infos , errors zerolog.Logger){
	r := newRouter()

	if err := http.ListenAndServe(":8080", r); err != nil {
		//panic(err)
		errors.Panic().Msg("panic")
	}
}

func newRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.HandleFunc("/", healthz)
	return r
}