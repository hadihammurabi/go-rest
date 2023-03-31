package service

import (
	"context"
	"errors"

	"github.com/gowok/gowok/policy"
	"github.com/gowok/ioc"
	"golang.org/x/exp/slices"
)

// PolicyService struct
type PolicyService struct {
	pol *policy.Policy
}

// NewPolicyService func
func NewPolicyService() PolicyService {
	return PolicyService{
		pol: ioc.MustGet(policy.Policy{}),
	}
}

// GetAllRoles func
func (u PolicyService) GetAllRoles(context.Context) []string {
	return u.pol.GetAllRoles()
}

// AddPolicy func
func (u PolicyService) AddPolicy(c context.Context, input []any) error {
	if len(input) < 3 {
		return errors.New("invalid policy data")
	}

	if len(input) == 3 {
		if input[0] != "g" {
			return errors.New("invalid policy data")
		}

		_, err := u.pol.AddGroupingPolicy(input[1:]...)
		if err != nil {
			return err
		}

		return nil
	}

	if input[0] != "p" || !slices.Contains(policy.Actions, input[3].(string)) {
		return errors.New("invalid policy data")
	}

	_, err := u.pol.AddPolicy(input[1:]...)
	if err != nil {
		return err
	}

	return nil
}

func (u PolicyService) DeleteRole(c context.Context, name string) error {
	users, err := u.pol.GetUsersForRole(name)
	if err != nil {
		return err
	}

	if len(users) > 0 {
		return errors.New("can't delete role with some users")
	}

	_, err = u.pol.DeleteRole(name)
	if err != nil {
		return err
	}

	return nil
}
