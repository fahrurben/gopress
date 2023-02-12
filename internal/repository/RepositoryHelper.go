package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func FindAll[T any](db *sqlx.DB, model T, querystr string, params []any, page int, pageSize int) ([]T, int, int, error) {
	offset := (page - 1) * pageSize
	params = append(params, pageSize)
	params = append(params, offset)
	rows, err := db.Queryx(querystr, params...)

	if err != nil {
		panic(err)
	}

	var results []T
	for rows.Next() {
		_ = rows.StructScan(&model)
		results = append(results, model)
	}

	var totalCount int
	queryPagination := fmt.Sprintf("select count(parent_table.id) from (%s) as parent_table", querystr)
	rowPagination, _ := db.Queryx(queryPagination, params...)
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
