// Code generated by go-swagger; DO NOT EDIT.

package rest_model_zrok

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EnvironmentAndResources environment and resources
//
// swagger:model environmentAndResources
type EnvironmentAndResources struct {

	// environment
	Environment *Environment `json:"environment,omitempty"`

	// frontends
	Frontends Frontends `json:"frontends,omitempty"`

	// shares
	Shares Shares `json:"shares,omitempty"`
}

// Validate validates this environment and resources
func (m *EnvironmentAndResources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrontends(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShares(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnvironmentAndResources) validateEnvironment(formats strfmt.Registry) error {
	if swag.IsZero(m.Environment) { // not required
		return nil
	}

	if m.Environment != nil {
		if err := m.Environment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("environment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("environment")
			}
			return err
		}
	}

	return nil
}

func (m *EnvironmentAndResources) validateFrontends(formats strfmt.Registry) error {
	if swag.IsZero(m.Frontends) { // not required
		return nil
	}

	if err := m.Frontends.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("frontends")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("frontends")
		}
		return err
	}

	return nil
}

func (m *EnvironmentAndResources) validateShares(formats strfmt.Registry) error {
	if swag.IsZero(m.Shares) { // not required
		return nil
	}

	if err := m.Shares.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("shares")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("shares")
		}
		return err
	}

	return nil
}

// ContextValidate validate this environment and resources based on the context it is used
func (m *EnvironmentAndResources) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEnvironment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFrontends(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShares(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnvironmentAndResources) contextValidateEnvironment(ctx context.Context, formats strfmt.Registry) error {

	if m.Environment != nil {
		if err := m.Environment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("environment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("environment")
			}
			return err
		}
	}

	return nil
}

func (m *EnvironmentAndResources) contextValidateFrontends(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Frontends.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("frontends")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("frontends")
		}
		return err
	}

	return nil
}

func (m *EnvironmentAndResources) contextValidateShares(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Shares.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("shares")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("shares")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnvironmentAndResources) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnvironmentAndResources) UnmarshalBinary(b []byte) error {
	var res EnvironmentAndResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
