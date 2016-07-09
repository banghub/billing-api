package units

import (
	"strconv"

	"github.com/py150504/billingps/src/global"
)

// Unit : data type unit
type Unit struct {
	ID     int64   `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status int64   `json:"-"`
}

func (u *Unit) save() error {
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
	return nil
}
func (u *Unit) update() error {
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
		return units, errSelect
	}
	for rows.Next() {
		unit := new(Unit)
		errScan := rows.Scan(
			&unit.ID,
			&unit.Name,
			&unit.Price)
		if errScan != nil {
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
