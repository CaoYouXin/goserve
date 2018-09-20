package migrate

import (
	"fmt"
	"strconv"

	"github.com/huandu/go-sqlbuilder"

	"github.com/CaoYouXin/goserve/orm"
)

// Init migrate
func Init() error {
	rows, err := orm.Query("show tables")
	if err != nil {
		return err
	}
	defer rows.Close()

	var tableName string
	for rows.Next() {
		rows.Scan(&tableName)
		if tableName == "configs" {
			break
		}
	}

	if tableName != "configs" {
		if _, err = orm.Exec(`CREATE TABLE configs (
			c_key VARCHAR(100) NOT NULL,
			c_value VARCHAR(1000) NULL,
			PRIMARY KEY (c_key))`); err != nil {
			return err
		}
	}

	row := orm.Row("SELECT c_value FROM configs WHERE c_key = 'last_index'")

	var lastIndexStr string
	row.Scan(&lastIndexStr)

	lastIndex, err := strconv.Atoi(lastIndexStr)
	if err != nil {
		fmt.Printf("%v", err.Error())
		lastIndex = 0
		_, err = orm.Exec("INSERT INTO configs (c_key, c_value) VALUES (?, ?)", "last_index", "0")
		if err != nil {
			return err
		}
	}

	var ret error

	for lastIndex < len(configs) {
		err = configs[lastIndex]()
		if err != nil {
			ret = err
			break
		}
		lastIndex++
	}

	if ret != nil {
		return ret
	}

	sb := sqlbuilder.Buildf("UPDATE configs SET c_value = %v WHERE c_key = 'last_index'", fmt.Sprintf("%v", lastIndex))
	sql, args := sb.BuildWithFlavor(sqlbuilder.MySQL)
	_, ret = orm.Exec(sql, args...)

	return ret
}
