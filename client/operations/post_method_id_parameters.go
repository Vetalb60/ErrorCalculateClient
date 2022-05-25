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
	"github.com/go-openapi/swag"

	"client/models"
)

// NewPostMethodIDParams creates a new PostMethodIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostMethodIDParams() *PostMethodIDParams {
	return &PostMethodIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostMethodIDParamsWithTimeout creates a new PostMethodIDParams object
// with the ability to set a timeout on a request.
func NewPostMethodIDParamsWithTimeout(timeout time.Duration) *PostMethodIDParams {
	return &PostMethodIDParams{
		timeout: timeout,
	}
}

// NewPostMethodIDParamsWithContext creates a new PostMethodIDParams object
// with the ability to set a context for a request.
func NewPostMethodIDParamsWithContext(ctx context.Context) *PostMethodIDParams {
	return &PostMethodIDParams{
		Context: ctx,
	}
}

// NewPostMethodIDParamsWithHTTPClient creates a new PostMethodIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostMethodIDParamsWithHTTPClient(client *http.Client) *PostMethodIDParams {
	return &PostMethodIDParams{
		HTTPClient: client,
	}
}

/* PostMethodIDParams contains all the parameters to send to the API endpoint
   for the post method ID operation.

   Typically these are written to a http.Request.
*/
type PostMethodIDParams struct {

	/* ID.

	   Identificator of decode file.
	*/
	ID int64

	/* Method.

	   The method of predict.
	*/
	Method *models.Method

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post method ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostMethodIDParams) WithDefaults() *PostMethodIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post method ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostMethodIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post method ID params
func (o *PostMethodIDParams) WithTimeout(timeout time.Duration) *PostMethodIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post method ID params
func (o *PostMethodIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post method ID params
func (o *PostMethodIDParams) WithContext(ctx context.Context) *PostMethodIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post method ID params
func (o *PostMethodIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post method ID params
func (o *PostMethodIDParams) WithHTTPClient(client *http.Client) *PostMethodIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post method ID params
func (o *PostMethodIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post method ID params
func (o *PostMethodIDParams) WithID(id int64) *PostMethodIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post method ID params
func (o *PostMethodIDParams) SetID(id int64) {
	o.ID = id
}

// WithMethod adds the method to the post method ID params
func (o *PostMethodIDParams) WithMethod(method *models.Method) *PostMethodIDParams {
	o.SetMethod(method)
	return o
}

// SetMethod adds the method to the post method ID params
func (o *PostMethodIDParams) SetMethod(method *models.Method) {
	o.Method = method
}

// WriteToRequest writes these params to a swagger request
func (o *PostMethodIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}
	if o.Method != nil {
		if err := r.SetBodyParam(o.Method); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
