package stock

import (
	"net/http"
	"yana/core/stock/yimport"

	"github.com/Dlacreme/httpd/back/flight"
	"github.com/Dlacreme/httpd/back/werror"
	"github.com/Dlacreme/httpd/back/wesult"
)

func Import(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	_, id, err := c.Config.Form.UploadFile(r, "file", 10000)

	if err != nil {
		res := wesult.New(nil, werror.New(http.StatusInternalServerError, err.Error()))
		v := c.View.New("status/error")
		res.RenderView(w, r, v)
		return
	}

	res := yimport.ImportFile(&c, c.Config.Form.FileStorageFolder+"/"+id)

	if res.Error != nil {
		res := wesult.New(nil, werror.New(http.StatusInternalServerError, err.Error()))
		v := c.View.New("status/error")
		res.RenderView(w, r, v)
	}
	c.FlashSuccess("Import success")
	Index(w, r)
}
