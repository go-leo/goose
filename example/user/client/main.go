package main

import (
	"net/http"

	"github.com/go-leo/goose/example/user"
)

func main() {
	router := http.NewServeMux()
	router = user.AppendUserGooseRoute(router, &MockUserService{})
	server := http.Server{Addr: ":8000", Handler: router}
	server.ListenAndServe()
}
