package server

import(
	"github.com/rs/zerolog"

	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// NewServer is used to start the http server, it also starts logging
func NewServer(infos , errors zerolog.Logger) *chi.Mux{
	r := newRouter()

	if err := http.ListenAndServe("0.0.0.0:8080", r); err != nil {
		//panic(err)
		//errors.Panic().Err(err).Stack()
		errors.Panic().Msg("Big big big error")
	}

	return(r)
}

func newRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	return r
}