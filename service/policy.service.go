package service

import (
	"context"
	"errors"

	"go-rest/driver"

	"github.com/gowok/ioc"
	"golang.org/x/exp/slices"
)

// PolicyService struct
type PolicyService struct {
	rbac *driver.RBAC
}

// NewPolicyService func
func NewPolicyService() PolicyService {
	return PolicyService{
		rbac: ioc.MustGet(driver.RBAC{}),
	}
}

// GetAllRoles func
func (u PolicyService) GetAllRoles(context.Context) []string {
	return u.rbac.GetAllRoles()
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

		_, err := u.rbac.AddGroupingPolicy(input[1:]...)
		if err != nil {
			return err
		}

		return nil
	}

	if input[0] != "p" || !slices.Contains(driver.RBACAll, input[3].(string)) {
		return errors.New("invalid policy data")
	}

	_, err := u.rbac.AddPolicy(input[1:]...)
	if err != nil {
		return err
	}

	return nil
}

func (u PolicyService) DeleteRole(c context.Context, name string) error {
	users, err := u.rbac.GetUsersForRole(name)
	if err != nil {
		return err
	}

	if len(users) > 0 {
		return errors.New("can't delete role with some users")
	}

	_, err = u.rbac.DeleteRole(name)
	if err != nil {
		return err
	}

	return nil
}
