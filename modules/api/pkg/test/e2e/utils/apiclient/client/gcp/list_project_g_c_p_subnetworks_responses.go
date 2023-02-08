// Code generated by go-swagger; DO NOT EDIT.

package gcp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListProjectGCPSubnetworksReader is a Reader for the ListProjectGCPSubnetworks structure.
type ListProjectGCPSubnetworksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectGCPSubnetworksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectGCPSubnetworksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListProjectGCPSubnetworksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectGCPSubnetworksOK creates a ListProjectGCPSubnetworksOK with default headers values
func NewListProjectGCPSubnetworksOK() *ListProjectGCPSubnetworksOK {
	return &ListProjectGCPSubnetworksOK{}
}

/*
ListProjectGCPSubnetworksOK describes a response with status code 200, with default header values.

GCPSubnetworkList
*/
type ListProjectGCPSubnetworksOK struct {
	Payload models.GCPSubnetworkList
}

// IsSuccess returns true when this list project g c p subnetworks o k response has a 2xx status code
func (o *ListProjectGCPSubnetworksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project g c p subnetworks o k response has a 3xx status code
func (o *ListProjectGCPSubnetworksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project g c p subnetworks o k response has a 4xx status code
func (o *ListProjectGCPSubnetworksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project g c p subnetworks o k response has a 5xx status code
func (o *ListProjectGCPSubnetworksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project g c p subnetworks o k response a status code equal to that given
func (o *ListProjectGCPSubnetworksOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListProjectGCPSubnetworksOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gcp/{dc}/subnetworks][%d] listProjectGCPSubnetworksOK  %+v", 200, o.Payload)
}

func (o *ListProjectGCPSubnetworksOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gcp/{dc}/subnetworks][%d] listProjectGCPSubnetworksOK  %+v", 200, o.Payload)
}

func (o *ListProjectGCPSubnetworksOK) GetPayload() models.GCPSubnetworkList {
	return o.Payload
}

func (o *ListProjectGCPSubnetworksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectGCPSubnetworksDefault creates a ListProjectGCPSubnetworksDefault with default headers values
func NewListProjectGCPSubnetworksDefault(code int) *ListProjectGCPSubnetworksDefault {
	return &ListProjectGCPSubnetworksDefault{
		_statusCode: code,
	}
}

/*
ListProjectGCPSubnetworksDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListProjectGCPSubnetworksDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list project g c p subnetworks default response
func (o *ListProjectGCPSubnetworksDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list project g c p subnetworks default response has a 2xx status code
func (o *ListProjectGCPSubnetworksDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list project g c p subnetworks default response has a 3xx status code
func (o *ListProjectGCPSubnetworksDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list project g c p subnetworks default response has a 4xx status code
func (o *ListProjectGCPSubnetworksDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list project g c p subnetworks default response has a 5xx status code
func (o *ListProjectGCPSubnetworksDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list project g c p subnetworks default response a status code equal to that given
func (o *ListProjectGCPSubnetworksDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListProjectGCPSubnetworksDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gcp/{dc}/subnetworks][%d] listProjectGCPSubnetworks default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectGCPSubnetworksDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gcp/{dc}/subnetworks][%d] listProjectGCPSubnetworks default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectGCPSubnetworksDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListProjectGCPSubnetworksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
