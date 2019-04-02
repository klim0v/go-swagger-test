// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// NewPostUploadParams creates a new PostUploadParams object
// no default values defined in spec.
func NewPostUploadParams() PostUploadParams {

	return PostUploadParams{}
}

// PostUploadParams contains all the bound params for the post upload operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostUpload
type PostUploadParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The file to upload.
	  In: formData
	*/
	Upfile io.ReadCloser
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostUploadParams() beforehand.
func (o *PostUploadParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}

	upfile, upfileHeader, err := r.FormFile("upfile")
	if err != nil && err != http.ErrMissingFile {
		res = append(res, errors.New(400, "reading file %q failed: %v", "upfile", err))
	} else if err == http.ErrMissingFile {
		// no-op for missing but optional file parameter
	} else if err := o.bindUpfile(upfile, upfileHeader); err != nil {
		res = append(res, err)
	} else {
		o.Upfile = &runtime.File{Data: upfile, Header: upfileHeader}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindUpfile binds file parameter Upfile.
//
// The only supported validations on files are MinLength and MaxLength
func (o *PostUploadParams) bindUpfile(file multipart.File, header *multipart.FileHeader) error {
	return nil
}
