package main

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pkg/errors"
)

type Data struct {
	db *sql.DB
}

func NewData() *Data {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	// mock sql
	columns := []string{"id", "name"}
	rows := sqlmock.NewRows(columns).AddRow(1, "Tom")
	mock.ExpectQuery("select (.+) from user where name = (.+) limit 1").
		WithArgs("Tom").
		WillReturnRows(rows).
		WillReturnError(sql.ErrNoRows)

	return &Data{db}
}

type userRepo struct {
	data *Data
}

func NewUserRepo(d *Data) UserRepo {
	return &userRepo{d}
}

func (r *userRepo) FindUser(targetName string) (*User, error) {
	var id int
	var name string
	err := r.data.db.QueryRow("select * from user where name = ? limit 1", targetName).Scan(&id, &name)

	// sql.ErrNoRows 降级处理
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "find user by sql query error")
	}
	return &User{id, name}, nil
}
