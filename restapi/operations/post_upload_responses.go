// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostUploadOKCode is the HTTP code returned for type PostUploadOK
const PostUploadOKCode int = 200

/*PostUploadOK File Uploaded

swagger:response postUploadOK
*/
type PostUploadOK struct {
}

// NewPostUploadOK creates PostUploadOK with default headers values
func NewPostUploadOK() *PostUploadOK {

	return &PostUploadOK{}
}

// WriteResponse to the client
func (o *PostUploadOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
