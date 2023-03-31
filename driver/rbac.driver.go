package driver

import (
	sqladapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/gowok/ioc"
)

const (
	RBACCreate = "CREATE"
	RBACRead   = "READ"
	RBACUpdate = "UPDATE"
	RBACDelete = "DELETE"
)

var RBACAll = []string{
	RBACCreate,
	RBACRead,
	RBACUpdate,
	RBACDelete,
}

type RBAC struct {
	*casbin.Enforcer
}

func NewRBAC() (*RBAC, error) {
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

	a, err := sqladapter.NewAdapter(ioc.Get(DB{}).DB, "postgres", "policies")
	if err != nil {
		return nil, err
	}

	e, err := casbin.NewEnforcer(model, a)
	if err != nil {
		return nil, err
	}

	e.DeleteRole("superuser")
	e.AddPolicy("superuser", "policies", RBACRead)
	e.AddPolicy("superuser", "policies", RBACCreate)
	e.AddPolicy("superuser", "policies", RBACUpdate)
	e.AddPolicy("superuser", "policies", RBACDelete)
	e.AddPolicy("superuser", "users", RBACRead)
	e.AddPolicy("superuser", "users", RBACCreate)
	e.AddPolicy("superuser", "users", RBACUpdate)
	e.AddPolicy("superuser", "users", RBACDelete)
	e.AddGroupingPolicy("root", "superuser")

	ee := &RBAC{e}
	return ee, nil
}

func initRBAC() {
	rbac, err := NewRBAC()
	if err != nil {
		panic(err)
	}
	ioc.Set(func() RBAC { return *rbac })

}
