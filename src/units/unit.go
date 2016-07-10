package units

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/py150504/billingps/src/global"
)

// Unit : data type unit
type Unit struct {
	ID     int64   `json:"-"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status int64   `json:"-"`
}

func (u *Unit) save() error {
	u.Status = 1
	resultInsert, errInsert := queryUnit.insert.Exec(u.Name, u.Price)
	if errInsert != nil {
		global.LogError.Printf(errInsert.Error())
		return errInsert
	}
	lastID, errResult := resultInsert.LastInsertId()
	if errResult != nil {
		global.LogError.Printf(errResult.Error())
		return errResult
	}
	u.ID = lastID
	return nil
}
func (u *Unit) load() error {
	errSelect := queryUnit.selectUnit.QueryRow(u.ID).Scan(
		&u.ID,
		&u.Name,
		&u.Price)
	if errSelect != nil {
		global.LogError.Printf(errSelect.Error())
		return errSelect
	}
	return nil
}
func (u *Unit) delete() error {
	resultDelete, errDelete := queryUnit.delete.Exec(u.ID)
	if errDelete != nil {
		global.LogError.Printf(errDelete.Error())
		return errDelete
	}
	affectedRow, errResult := resultDelete.RowsAffected()
	if errResult != nil {
		global.LogError.Printf(errResult.Error())
		return errResult
	}
	if affectedRow == 0 {
		noRow := errors.New("No row affected")
		global.LogError.Printf(noRow.Error())
		return noRow
	}
	return nil
}
func (u *Unit) update() error {
	resultUpdate, errUpdate := queryUnit.update.Exec(u.Name, u.Price, u.ID)
	if errUpdate != nil {
		global.LogError.Printf(errUpdate.Error())
		return errUpdate
	}
	affectedRow, errResult := resultUpdate.RowsAffected()
	if errResult != nil {
		global.LogError.Printf(errResult.Error())
		return errResult
	}
	if affectedRow == 0 {
		global.LogError.Printf(fmt.Sprintf("%d", affectedRow))
		return nil
	}
	return nil
}

func getUnit(id int64) (*Unit, error) {
	unit := new(Unit)
	unit.ID = id
	errLoad := unit.load()
	if errLoad != nil {
		return unit, errLoad
	}
	return unit, nil
}

func getUnits() ([]*Unit, error) {
	units := []*Unit{}
	rows, errSelect := queryUnit.selectUnits.Query()
	defer rows.Close()
	if errSelect != nil {
		global.LogError.Printf(errSelect.Error())
		return units, errSelect
	}
	for rows.Next() {
		unit := new(Unit)
		errScan := rows.Scan(
			&unit.ID,
			&unit.Name,
			&unit.Price)
		if errScan != nil {
			global.LogError.Printf(errScan.Error())
			return units, errScan
		}
		units = append(units, unit)
	}
	return units, nil
}

// MapUnit : map unit
func MapUnit(u *Unit, detail bool) interface{} {
	var attributes interface{}
	if detail {
		attributes = map[string]interface{}{
			"name":  u.Name,
			"price": u.Price,
		}
	}
	unit := map[string]interface{}{
		"id":         strconv.FormatInt(u.ID, 10),
		"type":       "unit",
		"attributes": attributes,
	}
	return unit
}
