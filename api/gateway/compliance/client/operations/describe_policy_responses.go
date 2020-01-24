// Code generated by go-swagger; DO NOT EDIT.

package operations

/**
 * Panther is a scalable, powerful, cloud-native SIEM written in Golang/React.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"

	models "github.com/panther-labs/panther/api/gateway/compliance/models"
)

// DescribePolicyReader is a Reader for the DescribePolicy structure.
type DescribePolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribePolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribePolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDescribePolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDescribePolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDescribePolicyOK creates a DescribePolicyOK with default headers values
func NewDescribePolicyOK() *DescribePolicyOK {
	return &DescribePolicyOK{}
}

/*DescribePolicyOK handles this case with default header values.

OK
*/
type DescribePolicyOK struct {
	Payload *models.PolicyResourceDetail
}

func (o *DescribePolicyOK) Error() string {
	return fmt.Sprintf("[GET /describe-policy][%d] describePolicyOK  %+v", 200, o.Payload)
}

func (o *DescribePolicyOK) GetPayload() *models.PolicyResourceDetail {
	return o.Payload
}

func (o *DescribePolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicyResourceDetail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribePolicyBadRequest creates a DescribePolicyBadRequest with default headers values
func NewDescribePolicyBadRequest() *DescribePolicyBadRequest {
	return &DescribePolicyBadRequest{}
}

/*DescribePolicyBadRequest handles this case with default header values.

Bad request
*/
type DescribePolicyBadRequest struct {
	Payload *models.Error
}

func (o *DescribePolicyBadRequest) Error() string {
	return fmt.Sprintf("[GET /describe-policy][%d] describePolicyBadRequest  %+v", 400, o.Payload)
}

func (o *DescribePolicyBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DescribePolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribePolicyInternalServerError creates a DescribePolicyInternalServerError with default headers values
func NewDescribePolicyInternalServerError() *DescribePolicyInternalServerError {
	return &DescribePolicyInternalServerError{}
}

/*DescribePolicyInternalServerError handles this case with default header values.

Internal server error
*/
type DescribePolicyInternalServerError struct {
}

func (o *DescribePolicyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /describe-policy][%d] describePolicyInternalServerError ", 500)
}

func (o *DescribePolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}