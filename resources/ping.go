package resources

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type PingResources struct{}

func (rs PingResources) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello World!"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}) // GET /ping - Healthcheck of server

	return r
}
