package user

import (
	"context"
	"fmt"
	"github.com/fahrurben/gopress/internal/repository"
	"github.com/jmoiron/sqlx"
	"time"
)

type RepositoryImpl struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *RepositoryImpl {
	return &RepositoryImpl{db: db}
}

func (r *RepositoryImpl) Save(user User) (*int64, error) {
	tx, err := r.db.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	exec, err := tx.Exec(InsertUser, user.Email, user.Name, user.Password, user.Type, user.CreatedAt, user.UpdatedAt)
	insertId, err := exec.LastInsertId()
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &insertId, err
}

func (r *RepositoryImpl) Update(user User) (bool, error) {
	_, err := r.db.Exec(UpdateUser, user.Name, time.Now(), user.Id)
	if err != nil {
		return false, err
	}
	return true, err
}

func (r *RepositoryImpl) FindById(id int) (*User, error) {
	row := r.db.QueryRowx(FindUserById, id)
	result := User{}
	err := row.StructScan(&result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

func (r *RepositoryImpl) FindAll(page int, pageSize int) ([]User, int, int, error) {
	offset := (page - 1) * pageSize
	rows, err := r.db.Queryx(FindAllUser, pageSize, offset)

	if err != nil {
		panic(err)
	}

	var results []User
	for rows.Next() {
		user := User{}
		_ = rows.StructScan(&user)
		results = append(results, user)
	}

	var totalCount int
	queryPagination := fmt.Sprintf("select count(parent_table.id) from (%s) as parent_table", FindAllUser)
	rowPagination, _ := r.db.Queryx(queryPagination, pageSize, offset)
	rowPagination.Next()
	err = rowPagination.Scan(&totalCount)

	if err != nil {
		return nil, 0, 0, err
	}

	var totalPage int = totalCount / pageSize
	if totalCount%pageSize > 0 {
		totalPage++
	}

	return results, totalCount, totalPage, err
}

func (r *RepositoryImpl) FindAllPagination(page int, pageSize int) ([]User, int, int, error) {
	return repository.FindAll[User](r.db, User{}, FindAllUser, nil, 1, 10)
}

func (r *RepositoryImpl) DeleteById(id int) error {
	tx, err := r.db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer tx.Rollback() // nolint
	_, err = tx.Exec(deleteUserById, id)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
