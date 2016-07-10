package units

import (
	"database/sql"

	"github.com/py150504/billingps/src/global"
)

var queryUnit preparedQueryUnit

type preparedQueryUnit struct {
	selectUnits *sql.Stmt
	selectUnit  *sql.Stmt
	insert      *sql.Stmt
	delete      *sql.Stmt
	update      *sql.Stmt
}

// InitUnit : prepare query
func InitUnit() error {
	db := global.DB.Core
	var errPrepared error
	queryUnit.selectUnits, errPrepared = db.Prepare(`
        SELECT
            id, name, price
        FROM 
            units
        WHERE
            status = 1`)
	if errPrepared != nil {
		global.LogError.Println(errPrepared.Error())
		return errPrepared
	}
	queryUnit.selectUnit, errPrepared = db.Prepare(`
        SELECT
            id, name, price
        FROM 
            units
        WHERE
            status = 1 AND
            id = ?`)
	if errPrepared != nil {
		global.LogError.Println(errPrepared.Error())
		return errPrepared
	}
	queryUnit.insert, errPrepared = db.Prepare(`
		INSERT INTO 
			units (name, price, status)
		VALUES 
			(?, ?, 1)`)
	if errPrepared != nil {
		global.LogError.Println(errPrepared.Error())
		return errPrepared
	}
	queryUnit.delete, errPrepared = db.Prepare(`
		UPDATE units
		SET
			status = 0
		WHERE
			id = ?`)
	if errPrepared != nil {
		global.LogError.Println(errPrepared.Error())
		return errPrepared
	}
	queryUnit.update, errPrepared = db.Prepare(`
		UPDATE units
		SET
			name = ?,
            price = ?
		WHERE
			id = ?`)
	if errPrepared != nil {
		global.LogError.Println(errPrepared.Error())
		return errPrepared
	}
	return nil
}
