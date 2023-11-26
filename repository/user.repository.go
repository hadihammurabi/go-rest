package repository

import (
	"context"
	"fmt"
	"go-rest/api/dto"
	"go-rest/entity"
	"go-rest/repository/table"

	"go-rest/driver"

	"github.com/gowok/qry"
	"gorm.io/gorm"
)

// UserSQL struct
type UserSQL struct {
	db *gorm.DB
}

// NewUser func
func NewUser() *UserSQL {
	return &UserSQL{
		db: driver.GetSQL(),
	}
}

// All func
func (r UserSQL) All(c context.Context, pagination dto.PaginationReq) (dto.PaginationRes[entity.User], error) {
	res := dto.NewPaginationRes(dto.PaginationRes[entity.User]{PaginationReq: &pagination})

	q := qry.Select("id", "email", "password").
		From(table.NameUsers)

	if filter, ok := pagination.Filter["search"]; ok {
		q = q.Where(fmt.Sprintf("email like '%%%s%%'", filter))
	}

	for k, v := range pagination.Sort {
		if k == "email" && v != "" {
			q = q.OrderBy(k, v)
		}
	}

	qCount := qry.Select("COUNT(datum.id)").
		From("(" + q.SQL() + ") datum").SQL()
	row := r.db.Raw(qCount)
	err := row.Scan(&res.Count).Error
	if err != nil {
		return res, err
	}

	q = q.Limit(int(pagination.Perpage)).
		Offset(int((pagination.Page * pagination.Perpage) - pagination.Perpage))

	err = r.db.Raw(q.SQL()).Scan(&res.Items).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

// Create func
func (r UserSQL) Create(c context.Context, user *entity.User) (*entity.User, error) {
	q := qry.Insert(table.NameUsers)
	q.Column("email", "password")
	q.Values("$1", "$2")
	q.Suffix("RETURNING id")

	err := r.db.Raw(q.SQL(), user.Email, user.Password).Scan(&user.ID).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// FindByEmail func
func (r UserSQL) FindByEmail(c context.Context, email string) (*entity.User, error) {
	userTable := table.User{}
	q := qry.Select("id", "email", "password").
		From(table.NameUsers).
		Where("email = $1")
	err := r.db.Raw(q.SQL(), email).Scan(&userTable).Error
	if err != nil {
		return nil, err
	}

	return userTable.ToEntity(), err
}

// FindByID func
func (r UserSQL) FindByID(c context.Context, id string) (*entity.User, error) {
	userTable := table.User{}
	q := qry.Select("id", "email", "password").
		From(table.NameUsers).
		Where("id = $1")
	err := r.db.Raw(q.SQL(), id).Scan(&userTable).Error
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

	q := qry.Update(table.NameUsers)
	q.Set("password", password)
	q.Where("id = $1")

	err = r.db.Exec(q.SQL(), id).Error
	if err != nil {
		return user, err
	}

	return userTable.ToEntity(), err
}
