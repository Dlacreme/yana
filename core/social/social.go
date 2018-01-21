package social

import (
	"github.com/Dlacreme/httpd/back/flight"
	"github.com/Dlacreme/httpd/back/wesult"
)

func GetNews(c *flight.Info) wesult.Result {
	return wesult.New(nil, nil)
}

func GetArticles(c *flight.Info) wesult.Result {
	return wesult.New(nil, nil)
}

func GetEvents(c *flight.Info) wesult.Result {
	return wesult.New(nil, nil)
}

func CreateNews(c *flight.Info, input CreateNewCtr) wesult.Result {
	return wesult.New(nil, nil)
}
