package restapi

import (
	"bytes"
	"encoding/csv"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/klim0v/go-swagger-test/restapi/operations"
	"io/ioutil"
	"net/http"
)

func getFile(r *http.Request) middleware.Responder {
	return middleware.ResponderFunc(
		func(w http.ResponseWriter, pr runtime.Producer) {
			var buf bytes.Buffer
			writer := csv.NewWriter(&buf)
			writer.Comma = ';'

			err := writer.Write([]string{"1", "gopher conference", "$100"})
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			writer.Flush()

			_, _ = w.Write(buf.Bytes())
		})
}

func uploadFile(params operations.PostUploadParams) middleware.Responder {
	return middleware.ResponderFunc(
		func(w http.ResponseWriter, pr runtime.Producer) {
			if err := params.BindRequest(params.HTTPRequest, nil); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			data, err := ioutil.ReadAll(params.Upfile)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			defer params.Upfile.Close()
			err = ioutil.WriteFile("/tmp/dat1.csv", data, 0644)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			_, _ = w.Write([]byte{1, 2})
		})
}
