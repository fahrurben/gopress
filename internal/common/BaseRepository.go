package common

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type BaseRepository struct {
	db *sqlx.DB
}

func (b BaseRepository) GetPagingDetails(querystr string, page int, pageSize int, params []any) (int, int, error) {
	var totalCount int
	offset := (page - 1) * pageSize
	params = append(params, pageSize)
	params = append(params, offset)
	queryPagination := fmt.Sprintf("select count(parent_table.id) from (%s) as parent_table", querystr)
	rowPagination, _ := b.db.Queryx(queryPagination, params...)
	rowPagination.Next()
	err := rowPagination.Scan(&totalCount)

	if err != nil {
		return 0, 0, err
	}

	var totalPage int = totalCount / pageSize
	if totalCount%pageSize > 0 {
		totalPage++
	}

	return totalCount, totalPage, err
}
