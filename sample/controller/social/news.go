package social

import (
	"fmt"
	"net/http"

	"github.com/Dlacreme/httpd/back/flight"
)

func GetNews(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	fmt.Printf("Not implemented: %v\n", c)

	return
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	fmt.Printf("Not implemented: %v\n", c)

	return
}

func GetEvents(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	fmt.Printf("Not implemented: %v\n", c)

	return
}
