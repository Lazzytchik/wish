package app

import (
	"log"
	"net/http"
)

type Service struct {
	Server *http.Server
	Api    Usecaser
	Logger *log.Logger
}

func (s *Service) GetWishes(w http.ResponseWriter, r *http.Request) {
	JsonAnswer(w, r, "get wishes")
}

func (s *Service) Wish(w http.ResponseWriter, r *http.Request) {
	JsonAnswer(w, r, "wish")
}
