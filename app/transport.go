package app

import "net/http"

func (s *Service) Routes() *http.ServeMux {
	v0 := http.NewServeMux()

	v0.HandleFunc("/wishes", s.GetWishes)
	v0.HandleFunc("/wish", s.Wish)

	return v0
}
