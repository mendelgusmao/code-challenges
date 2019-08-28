package middleware

import (
	"mime"
	"net/http"

	"github.com/gorilla/mux"
)

const expectedMIMEType = "application/x-www-form-urlencoded"

func IsURLEncodedForm(r *http.Request, rm *mux.RouteMatch) bool {
	headerValue := r.Header.Get("Content-Type")
	contentType, _, _ := mime.ParseMediaType(headerValue)

	return contentType == expectedMIMEType
}
