package repository

import (
	"context"
	"fmt"
	"go-rest/api/dto"
	"go-rest/entity"
	"go-rest/repository/table"

	"go-rest/driver"

	"github.com/gowok/ioc"
	"github.com/gowok/qry"
)

// UserSQL struct
type UserSQL struct {
	db *driver.DB
}

// NewUser func
func NewUser() *UserSQL {
	return &UserSQL{
		db: ioc.MustGet(driver.DB{}),
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
	row := r.db.QueryRow(qCount)
	err := row.Scan(&res.Count)
	if err != nil {
		return res, err
	}

	q = q.Limit(int(pagination.Perpage)).
		Offset(int((pagination.Page * pagination.Perpage) - pagination.Perpage))

	rows, err := r.db.Query(q.SQL())
	if err != nil {
		return res, err
	}

	res.Items = make([]entity.User, 0)
	for rows.Next() {
		userTable := table.User{}
		rows.Scan(
			&userTable.ID,
			&userTable.Email,
			&userTable.Password,
		)

		res.Items = append(res.Items, *userTable.ToEntity())
	}

	return res, nil
}

// Create func
func (r UserSQL) Create(c context.Context, user *entity.User) (*entity.User, error) {
	q := qry.Insert(table.NameUsers).
		Column("email", "password").
		Values("$1", "$2").
		Suffix("RETURNING id")
	row := r.db.QueryRow(q.SQL(), user.Email, user.Password)

	err := row.Scan(&user.ID)
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
	row := r.db.QueryRow(q.SQL(), email)
	err := row.Err()
	if err != nil {
		return nil, err
	}
	err = row.Scan(
		&userTable.ID,
		&userTable.Email,
		&userTable.Password,
	)
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
	row := r.db.QueryRow(q.SQL(), id)
	err := row.Scan(
		&userTable.ID,
		&userTable.Email,
		&userTable.Password,
	)
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

	q := qry.Update(table.NameUsers).
		Set("password", password).
		Where("id = $1")

	_, err = r.db.Exec(q.SQL(), id)
	if err != nil {
		return user, err
	}

	return userTable.ToEntity(), err
}
