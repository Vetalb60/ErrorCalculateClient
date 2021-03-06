// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPostUploadParams creates a new PostUploadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostUploadParams() *PostUploadParams {
	return &PostUploadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostUploadParamsWithTimeout creates a new PostUploadParams object
// with the ability to set a timeout on a request.
func NewPostUploadParamsWithTimeout(timeout time.Duration) *PostUploadParams {
	return &PostUploadParams{
		timeout: timeout,
	}
}

// NewPostUploadParamsWithContext creates a new PostUploadParams object
// with the ability to set a context for a request.
func NewPostUploadParamsWithContext(ctx context.Context) *PostUploadParams {
	return &PostUploadParams{
		Context: ctx,
	}
}

// NewPostUploadParamsWithHTTPClient creates a new PostUploadParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostUploadParamsWithHTTPClient(client *http.Client) *PostUploadParams {
	return &PostUploadParams{
		HTTPClient: client,
	}
}

/* PostUploadParams contains all the parameters to send to the API endpoint
   for the post upload operation.

   Typically these are written to a http.Request.
*/
type PostUploadParams struct {

	/* Upfile.

	   The file to upload.
	*/
	Upfile runtime.NamedReadCloser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post upload params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUploadParams) WithDefaults() *PostUploadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post upload params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUploadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post upload params
func (o *PostUploadParams) WithTimeout(timeout time.Duration) *PostUploadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post upload params
func (o *PostUploadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post upload params
func (o *PostUploadParams) WithContext(ctx context.Context) *PostUploadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post upload params
func (o *PostUploadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post upload params
func (o *PostUploadParams) WithHTTPClient(client *http.Client) *PostUploadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post upload params
func (o *PostUploadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpfile adds the upfile to the post upload params
func (o *PostUploadParams) WithUpfile(upfile runtime.NamedReadCloser) *PostUploadParams {
	o.SetUpfile(upfile)
	return o
}

// SetUpfile adds the upfile to the post upload params
func (o *PostUploadParams) SetUpfile(upfile runtime.NamedReadCloser) {
	o.Upfile = upfile
}

// WriteToRequest writes these params to a swagger request
func (o *PostUploadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Upfile != nil {

		if o.Upfile != nil {
			// form file param upfile
			if err := r.SetFileParam("upfile", o.Upfile); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
