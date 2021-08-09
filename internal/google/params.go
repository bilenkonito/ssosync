package google

import (
	"errors"
)

type Params interface {
	Validate() error
}

type GetParams struct {
	Id string
}
func (r *GetParams) Validate() error {
	if r.Id == "" {
		return errors.New("google.client: empty Id provided")
	}

	return nil
}

type AccountParams struct {
	Customer string
	Domain string
}
type DeletedParams = AccountParams
func (r *AccountParams) Validate() error {
	if r.Customer == "" && r.Domain == "" {
		return errors.New("google.client: empty Customer and Domain provided")
	}

	return nil
}

type QueryParams struct {
	AccountParams
	Query string
	ShowDeleted bool
}
func (r *QueryParams) Validate() error {
	return r.AccountParams.Validate()
}

type GroupMembersParams struct {
	GetParams
	Parent *Group
}
func (r *GroupMembersParams) Validate() error {
	err := r.GetParams.Validate()
	if err != nil && r.Parent == nil {
		return errors.New("google.client: empty Parent and Id provided")
	}

	return nil
}

type ClientParams struct {
	AdminEmail string
	ServiceAccountKey []byte
}
func (r *ClientParams) Validate() error {
	if r.AdminEmail == "" {
		return errors.New("google.client: empty AdminEmail provided")
	}
	if r.ServiceAccountKey == nil || len(r.ServiceAccountKey) < 1 {
		return errors.New("google.client: empty ServiceAccountKey provided")
	}

	return nil
}
