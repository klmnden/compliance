package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// SentimentBatchResultV2 sentiment batch result v2
// swagger:model SentimentBatchResultV2
type SentimentBatchResultV2 struct {

	// documents
	Documents []*SentimentBatchResultItemV2 `json:"documents"`

	// errors
	Errors []*ErrorRecordV2 `json:"errors"`
}

// Validate validates this sentiment batch result v2
func (m *SentimentBatchResultV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocuments(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SentimentBatchResultV2) validateDocuments(formats strfmt.Registry) error {

	if swag.IsZero(m.Documents) { // not required
		return nil
	}

	for i := 0; i < len(m.Documents); i++ {

		if swag.IsZero(m.Documents[i]) { // not required
			continue
		}

		if m.Documents[i] != nil {

			if err := m.Documents[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *SentimentBatchResultV2) validateErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	for i := 0; i < len(m.Errors); i++ {

		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {

			if err := m.Errors[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
