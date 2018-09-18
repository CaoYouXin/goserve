package bghis

import (
	"github.com/CaoYouXin/goserve/orm"
	sqlbuilder "github.com/huandu/go-sqlbuilder"
)

// SNTable for test
type SNTable struct {
	Number       int `db:"number"`
	SquareNumber int `db:"squareNumber"`
}

var snStruct = sqlbuilder.NewStruct(new(SNTable))

// GetSquareNumber get square number
func GetSquareNumber(number int) (*SNTable, error) {
	sb := snStruct.SelectFrom("squareNum")
	sb.Where(sb.E("number", number))

	sql, args := sb.Build()
	rows, err := orm.Query(sql, args...)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var sn SNTable
	for rows.Next() {
		rows.Scan(snStruct.Addr(&sn)...)
	}

	return &sn, nil
}
