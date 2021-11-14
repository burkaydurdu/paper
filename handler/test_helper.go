package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"paper.com/paper/model"
)

func CreateHTTPReq(method, endpoint string, body io.Reader) (w *httptest.ResponseRecorder, req *http.Request) {
	if body == nil {
		const size = 512
		body = bytes.NewBuffer(make([]byte, size))
	}
	req = httptest.NewRequest(method, endpoint, body)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	rec := httptest.NewRecorder()
	return rec, req
}

func ParseBody(body []byte) model.ColorResponse {
	var actualResponseBody model.ColorResponse
	err := json.Unmarshal(body, &actualResponseBody)
	if err != nil {
		fmt.Println(err)
	}
	return actualResponseBody
}
