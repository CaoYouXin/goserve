package migrate

import "github.com/CaoYouXin/goserve/orm"

func missionCreateBgHistoryTable() error {
	_, err := orm.Exec(`CREATE TABLE bg_history (
		id INT NOT NULL AUTO_INCREMENT,
		created_at DATETIME NOT NULL,
		deleted_at DATETIME NULL,
		uuid VARCHAR(100) NOT NULL,
		PRIMARY KEY (id))`)
	return err
}
