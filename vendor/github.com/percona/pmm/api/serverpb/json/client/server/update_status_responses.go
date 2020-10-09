// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateStatusReader is a Reader for the UpdateStatus structure.
type UpdateStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateStatusOK creates a UpdateStatusOK with default headers values
func NewUpdateStatusOK() *UpdateStatusOK {
	return &UpdateStatusOK{}
}

/*UpdateStatusOK handles this case with default header values.

A successful response.
*/
type UpdateStatusOK struct {
	Payload *UpdateStatusOKBody
}

func (o *UpdateStatusOK) Error() string {
	return fmt.Sprintf("[POST /v1/Updates/Status][%d] updateStatusOk  %+v", 200, o.Payload)
}

func (o *UpdateStatusOK) GetPayload() *UpdateStatusOKBody {
	return o.Payload
}

func (o *UpdateStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateStatusOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateStatusDefault creates a UpdateStatusDefault with default headers values
func NewUpdateStatusDefault(code int) *UpdateStatusDefault {
	return &UpdateStatusDefault{
		_statusCode: code,
	}
}

/*UpdateStatusDefault handles this case with default header values.

An unexpected error response.
*/
type UpdateStatusDefault struct {
	_statusCode int

	Payload *UpdateStatusDefaultBody
}

// Code gets the status code for the update status default response
func (o *UpdateStatusDefault) Code() int {
	return o._statusCode
}

func (o *UpdateStatusDefault) Error() string {
	return fmt.Sprintf("[POST /v1/Updates/Status][%d] UpdateStatus default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateStatusDefault) GetPayload() *UpdateStatusDefaultBody {
	return o.Payload
}

func (o *UpdateStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateStatusDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateStatusBody update status body
swagger:model UpdateStatusBody
*/
type UpdateStatusBody struct {

	// Authentication token.
	AuthToken string `json:"auth_token,omitempty"`

	// Progress log offset.
	LogOffset int64 `json:"log_offset,omitempty"`
}

// Validate validates this update status body
func (o *UpdateStatusBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateStatusBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateStatusBody) UnmarshalBinary(b []byte) error {
	var res UpdateStatusBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateStatusDefaultBody update status default body
swagger:model UpdateStatusDefaultBody
*/
type UpdateStatusDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this update status default body
func (o *UpdateStatusDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateStatusDefaultBody) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("UpdateStatus default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateStatusDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateStatusDefaultBody) UnmarshalBinary(b []byte) error {
	var res UpdateStatusDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateStatusOKBody update status OK body
swagger:model UpdateStatusOKBody
*/
type UpdateStatusOKBody struct {

	// Progress log lines.
	LogLines []string `json:"log_lines"`

	// Progress log offset for the next request.
	LogOffset int64 `json:"log_offset,omitempty"`

	// True when update is done.
	Done bool `json:"done,omitempty"`
}

// Validate validates this update status OK body
func (o *UpdateStatusOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateStatusOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateStatusOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateStatusOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
