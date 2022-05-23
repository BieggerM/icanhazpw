package api

import (
	"encoding/json"
	"github.com/BieggerM/icanhazpw/service"
	"github.com/gorilla/mux"
	"net/http"
)

type PasswordParameters struct {
	NumCharacters int  `json:"numchar"`
	Digits        bool `json:"digits"`
	Symbols       bool `json:"symbols"`
}

type Server struct {
	*mux.Router
	param PasswordParameters
}

func NewServer() *Server {
	server := &Server{
		Router: mux.NewRouter(),
		param:  PasswordParameters{},
	}
	server.routes()
	return server
}

func (s *Server) routes() {
	s.HandleFunc("/", s.getPassword()).Methods("GET")
}

func (s *Server) getPassword() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		response, err := service.GeneratePassword(16, true, true)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
		}
		writer.Write([]byte(response))
	}
}

func returnJson(writer http.ResponseWriter, response string) {
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) getPasswordWithParams() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var param PasswordParameters
		err := json.NewDecoder(request.Body).Decode(param)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
		}
		response, err := service.GeneratePassword(param.NumCharacters, param.Digits, param.Symbols)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
		}
		returnJson(writer, response)
	}
}
