package driver

import (
	sqladapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2/model"
	"github.com/gowok/gowok/policy"
	"github.com/gowok/ioc"
)

func NewPolicy() (*policy.Policy, error) {
	model, err := model.NewModelFromString(`
		[request_definition]
		r = sub, obj, act

		[policy_definition]
		p = sub, obj, act

		[role_definition]
		g = _, _

		[policy_effect]
		e = some(where (p.eft == allow))

		[matchers]
		m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
	`)
	if err != nil {
		return nil, err
	}

	db, err := GetSQL().DB()
	if err != nil {
		return nil, err
	}

	a, err := sqladapter.NewAdapter(db, "postgres", "policies")
	if err != nil {
		return nil, err
	}

	e, err := policy.NewPolicy(model, a)
	if err != nil {
		return nil, err
	}

	e.DeleteRole("superuser")
	e.AddPolicy("superuser", "policies", policy.ActionRead)
	e.AddPolicy("superuser", "policies", policy.ActionCreate)
	e.AddPolicy("superuser", "policies", policy.ActionUpdate)
	e.AddPolicy("superuser", "policies", policy.ActionDelete)
	e.AddPolicy("superuser", "users", policy.ActionRead)
	e.AddPolicy("superuser", "users", policy.ActionCreate)
	e.AddPolicy("superuser", "users", policy.ActionUpdate)
	e.AddPolicy("superuser", "users", policy.ActionDelete)
	e.AddGroupingPolicy("root", "superuser")

	return e, nil
}

func initPolicy() {
	e, err := NewPolicy()
	if err != nil {
		panic(err)
	}

	ioc.Set(func() policy.Policy { return *e })
}
