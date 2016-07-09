package people

import (
	"database/sql"
	"log"

	"github.com/py150504/billingps/src/global"
)

var queryPerson preparedQueryPerson

type preparedQueryPerson struct {
	selectPeople *sql.Stmt
	selectPerson *sql.Stmt
	insert       *sql.Stmt
	delete       *sql.Stmt
	update       *sql.Stmt
}

// InitPeople : prepare query
func InitPeople() error {
	db := global.DB.Core
	var errPrepared error
	queryPerson.selectPeople, errPrepared = db.Prepare(`
		SELECT
			id, name, phone, join_date
		FROM
			people
		WHERE
			status = 1`)
	if errPrepared != nil {
		log.Printf("Error prepare select people : %s", errPrepared.Error())
		log.Fatal("App exit, fail Init People")
	}
	queryPerson.selectPerson, errPrepared = db.Prepare(`
		SELECT
			id, name, phone, join_date
		FROM
			people
		WHERE
			status = 1 AND
			id = ?`)
	if errPrepared != nil {
		log.Printf("Error prepare select person : %s", errPrepared.Error())
		log.Fatal("App exit, fail Init People")
	}
	queryPerson.insert, errPrepared = db.Prepare(`
		INSERT INTO 
			people (name, phone, join_date, status)
		VALUES 
			(?, ?, ?, ?)`)

	if errPrepared != nil {
		log.Printf("Error prepare insert person : %s", errPrepared.Error())
		log.Fatal("App exit, fail Init People")
	}
	queryPerson.delete, errPrepared = db.Prepare(`
		UPDATE people
		SET
			status = 0
		WHERE
			id = ?`)

	if errPrepared != nil {
		log.Printf("Error prepare delete person : %s", errPrepared.Error())
		log.Fatal("App exit, fail Init People")
	}
	queryPerson.update, errPrepared = db.Prepare(`
		UPDATE people
		SET
			name = ?,
            phone = ?
		WHERE
			id = ?`)

	if errPrepared != nil {
		log.Printf("Error prepare update person : %s", errPrepared.Error())
		log.Fatal("App exit, fail Init People")
	}
	return nil
}
