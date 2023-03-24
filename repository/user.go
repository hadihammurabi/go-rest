package repository

import (
	"context"
	"fmt"
	"go-rest/api/dto"
	"go-rest/entity"
	"go-rest/repository/table"

	"go-rest/driver"

	"github.com/gowok/ioc"
)

// UserSQL struct
type UserSQL struct {
	db *driver.DB
}

// NewUser func
func NewUser() *UserSQL {
	return &UserSQL{
		db: ioc.Get(driver.DB{}),
	}
}

// All func
func (r UserSQL) All(c context.Context, pagination dto.PaginationReq) (dto.PaginationRes[entity.User], error) {
	res := dto.NewPaginationRes(dto.PaginationRes[entity.User]{PaginationReq: &pagination})
	usersTable := []*table.User{}
	q := r.db.NewSelect().
		Model(&usersTable).
		Column("id", "email", "password").
		Limit(int(pagination.Perpage)).
		Offset(int(pagination.Offset()))

	if filter, ok := pagination.Filter["search"]; ok {
		q.Where(fmt.Sprintf("email like '%%%s%%'", filter))
	}

	for k, v := range pagination.Sort {
		if k == "email" && v != "" {
			q.Order(fmt.Sprintf("%s %s", k, v))
		}
	}

	count, err := q.ScanAndCount(c)
	if err != nil {
		return res, err
	}

	res.Count = uint(count)
	res.Items = make([]entity.User, 0)
	for _, v := range usersTable {
		res.Items = append(res.Items, *v.ToEntity())
	}

	return res, err
}

// Create func
func (r UserSQL) Create(c context.Context, user *entity.User) (*entity.User, error) {
	userTable := table.UserFromEntity(user)
	// userTable.ID = uuid.NewString()
	_, err := r.db.NewInsert().Model(userTable).Exec(c)
	if err != nil {
		return user, err
	}

	return user, err
}

// FindByEmail func
func (r UserSQL) FindByEmail(c context.Context, email string) (*entity.User, error) {
	userTable := table.User{}
	err := r.db.NewSelect().
		Column("id", "email", "password").
		Where("email", email).
		Scan(c)
	if err != nil {
		return nil, err
	}

	return userTable.ToEntity(), err
}

// FindByID func
func (r UserSQL) FindByID(c context.Context, id string) (*entity.User, error) {
	userTable := table.User{}
	err := r.db.NewSelect().
		Model(&userTable).
		Column("id", "email", "password").
		Where("id = ?", id).
		Scan(c)
	if err != nil {
		return nil, err
	}

	return userTable.ToEntity(), err
}

// ChangePassword func
func (r UserSQL) ChangePassword(c context.Context, id string, password string) (*entity.User, error) {
	user, err := r.FindByID(c, id)
	if err != nil {
		return nil, err
	}

	userTable := table.UserFromEntity(user)

	_, err = r.db.NewUpdate().
		Model(userTable).
		Set("password = ?", password).
		WherePK().
		Exec(c)
	fmt.Println(userTable, err)
	return userTable.ToEntity(), err
}
