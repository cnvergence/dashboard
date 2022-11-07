// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// DeleteUserFromProjectReader is a Reader for the DeleteUserFromProject structure.
type DeleteUserFromProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserFromProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUserFromProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteUserFromProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteUserFromProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteUserFromProjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteUserFromProjectOK creates a DeleteUserFromProjectOK with default headers values
func NewDeleteUserFromProjectOK() *DeleteUserFromProjectOK {
	return &DeleteUserFromProjectOK{}
}

/*
DeleteUserFromProjectOK describes a response with status code 200, with default header values.

User
*/
type DeleteUserFromProjectOK struct {
	Payload *models.User
}

// IsSuccess returns true when this delete user from project o k response has a 2xx status code
func (o *DeleteUserFromProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete user from project o k response has a 3xx status code
func (o *DeleteUserFromProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user from project o k response has a 4xx status code
func (o *DeleteUserFromProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete user from project o k response has a 5xx status code
func (o *DeleteUserFromProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user from project o k response a status code equal to that given
func (o *DeleteUserFromProjectOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteUserFromProjectOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/users/{user_id}][%d] deleteUserFromProjectOK  %+v", 200, o.Payload)
}

func (o *DeleteUserFromProjectOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/users/{user_id}][%d] deleteUserFromProjectOK  %+v", 200, o.Payload)
}

func (o *DeleteUserFromProjectOK) GetPayload() *models.User {
	return o.Payload
}

func (o *DeleteUserFromProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserFromProjectUnauthorized creates a DeleteUserFromProjectUnauthorized with default headers values
func NewDeleteUserFromProjectUnauthorized() *DeleteUserFromProjectUnauthorized {
	return &DeleteUserFromProjectUnauthorized{}
}

/*
DeleteUserFromProjectUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type DeleteUserFromProjectUnauthorized struct {
}

// IsSuccess returns true when this delete user from project unauthorized response has a 2xx status code
func (o *DeleteUserFromProjectUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user from project unauthorized response has a 3xx status code
func (o *DeleteUserFromProjectUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user from project unauthorized response has a 4xx status code
func (o *DeleteUserFromProjectUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user from project unauthorized response has a 5xx status code
func (o *DeleteUserFromProjectUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user from project unauthorized response a status code equal to that given
func (o *DeleteUserFromProjectUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteUserFromProjectUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/users/{user_id}][%d] deleteUserFromProjectUnauthorized ", 401)
}

func (o *DeleteUserFromProjectUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/users/{user_id}][%d] deleteUserFromProjectUnauthorized ", 401)
}

func (o *DeleteUserFromProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserFromProjectForbidden creates a DeleteUserFromProjectForbidden with default headers values
func NewDeleteUserFromProjectForbidden() *DeleteUserFromProjectForbidden {
	return &DeleteUserFromProjectForbidden{}
}

/*
DeleteUserFromProjectForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type DeleteUserFromProjectForbidden struct {
}

// IsSuccess returns true when this delete user from project forbidden response has a 2xx status code
func (o *DeleteUserFromProjectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user from project forbidden response has a 3xx status code
func (o *DeleteUserFromProjectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user from project forbidden response has a 4xx status code
func (o *DeleteUserFromProjectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user from project forbidden response has a 5xx status code
func (o *DeleteUserFromProjectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user from project forbidden response a status code equal to that given
func (o *DeleteUserFromProjectForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteUserFromProjectForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/users/{user_id}][%d] deleteUserFromProjectForbidden ", 403)
}

func (o *DeleteUserFromProjectForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/users/{user_id}][%d] deleteUserFromProjectForbidden ", 403)
}

func (o *DeleteUserFromProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserFromProjectDefault creates a DeleteUserFromProjectDefault with default headers values
func NewDeleteUserFromProjectDefault(code int) *DeleteUserFromProjectDefault {
	return &DeleteUserFromProjectDefault{
		_statusCode: code,
	}
}

/*
DeleteUserFromProjectDefault describes a response with status code -1, with default header values.

errorResponse
*/
type DeleteUserFromProjectDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete user from project default response
func (o *DeleteUserFromProjectDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete user from project default response has a 2xx status code
func (o *DeleteUserFromProjectDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete user from project default response has a 3xx status code
func (o *DeleteUserFromProjectDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete user from project default response has a 4xx status code
func (o *DeleteUserFromProjectDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete user from project default response has a 5xx status code
func (o *DeleteUserFromProjectDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete user from project default response a status code equal to that given
func (o *DeleteUserFromProjectDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteUserFromProjectDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/users/{user_id}][%d] deleteUserFromProject default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteUserFromProjectDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/users/{user_id}][%d] deleteUserFromProject default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteUserFromProjectDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteUserFromProjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}