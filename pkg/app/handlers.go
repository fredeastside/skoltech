package app

import (
	"encoding/json"
	"net/http"
)

//Handlers returs app server handler
func (s *Server) Handlers() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/", s.onlyPost(s.onlyJSON(s.handle())))

	return r
}

func (s *Server) handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		d := s.devicesService.CreateDevice()
		err := json.NewDecoder(r.Body).Decode(d)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = s.devicesService.Save(d)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) onlyPost(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		h(w, r)
	}
}

func (s *Server) onlyJSON(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}
		h(w, r)
	}
}
