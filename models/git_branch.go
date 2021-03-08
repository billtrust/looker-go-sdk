// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GitBranch git branch
//
// swagger:model GitBranch
type GitBranch struct {

	// Number of commits the local branch is ahead of the remote
	// Read Only: true
	AheadCount int64 `json:"ahead_count,omitempty"`

	// Number of commits the local branch is behind the remote
	// Read Only: true
	BehindCount int64 `json:"behind_count,omitempty"`

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// UNIX timestamp at which this branch was last committed.
	// Read Only: true
	CommitAt int64 `json:"commit_at,omitempty"`

	// Name of error
	// Read Only: true
	Error string `json:"error,omitempty"`

	// Whether or not a local ref exists for the branch
	// Read Only: true
	IsLocal *bool `json:"is_local,omitempty"`

	// Whether or not this is the production branch
	// Read Only: true
	IsProduction *bool `json:"is_production,omitempty"`

	// Whether or not a remote ref exists for the branch
	// Read Only: true
	IsRemote *bool `json:"is_remote,omitempty"`

	// Message describing an error if present
	// Read Only: true
	Message string `json:"message,omitempty"`

	// The short name on the local. Updating `name` results in `git checkout <new_name>`
	Name string `json:"name,omitempty"`

	// Name of the owner of a personal branch
	// Read Only: true
	OwnerName string `json:"owner_name,omitempty"`

	// Whether or not this branch is a personal branch - readonly for all developers except the owner
	// Read Only: true
	Personal *bool `json:"personal,omitempty"`

	// Whether or not this branch is readonly
	// Read Only: true
	Readonly *bool `json:"readonly,omitempty"`

	// The resolved ref of this branch. Updating `ref` results in `git reset --hard <new_ref>``.
	Ref string `json:"ref,omitempty"`

	// The name of the remote
	// Read Only: true
	Remote string `json:"remote,omitempty"`

	// The short name on the remote
	// Read Only: true
	RemoteName string `json:"remote_name,omitempty"`

	// The resolved ref of this branch remote.
	// Read Only: true
	RemoteRef string `json:"remote_ref,omitempty"`
}

// Validate validates this git branch
func (m *GitBranch) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this git branch based on the context it is used
func (m *GitBranch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAheadCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBehindCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCommitAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsLocal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsProduction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsRemote(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwnerName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePersonal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReadonly(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRemote(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRemoteName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRemoteRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitBranch) contextValidateAheadCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "ahead_count", "body", int64(m.AheadCount)); err != nil {
		return err
	}

	return nil
}

func (m *GitBranch) contextValidateBehindCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "behind_count", "body", int64(m.BehindCount)); err != nil {
		return err
	}

	return nil
}

func (m *GitBranch) contextValidateCan(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *GitBranch) contextValidateCommitAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "commit_at", "body", int64(m.CommitAt)); err != nil {
		return err
	}

	return nil
}

func (m *GitBranch) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "error", "body", string(m.Error)); err != nil {
		return err
	}

	return nil
}

func (m *GitBranch) contextValidateIsLocal(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "is_local", "body", m.IsLocal); err != nil {
		return err
	}

	return nil
}

func (m *GitBranch) contextValidateIsProduction(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "is_production", "body", m.IsProduction); err != nil {
		return err
	}

	return nil
}

func (m *GitBranch) contextValidateIsRemote(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "is_remote", "body", m.IsRemote); err != nil {
		return err
	}

	return nil
}

func (m *GitBranch) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "message", "body", string(m.Message)); err != nil {
		return err
	}

	return nil
}

func (m *GitBranch) contextValidateOwnerName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "owner_name", "body", string(m.OwnerName)); err != nil {
		return err
	}

	return nil
}

func (m *GitBranch) contextValidatePersonal(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "personal", "body", m.Personal); err != nil {
		return err
	}

	return nil
}

func (m *GitBranch) contextValidateReadonly(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "readonly", "body", m.Readonly); err != nil {
		return err
	}

	return nil
}

func (m *GitBranch) contextValidateRemote(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "remote", "body", string(m.Remote)); err != nil {
		return err
	}

	return nil
}

func (m *GitBranch) contextValidateRemoteName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "remote_name", "body", string(m.RemoteName)); err != nil {
		return err
	}

	return nil
}

func (m *GitBranch) contextValidateRemoteRef(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "remote_ref", "body", string(m.RemoteRef)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GitBranch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitBranch) UnmarshalBinary(b []byte) error {
	var res GitBranch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
