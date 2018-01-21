package social

import (
	"fmt"
	"net/http"

	"github.com/Dlacreme/httpd/back/flight"
)

func LoadCreateNews(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	fmt.Printf("Not implemented: %v\n", c)

	return
}

func CreateNews(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	fmt.Printf("Not implemented: %v\n", c)

	return
}
