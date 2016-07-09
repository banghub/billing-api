package units

import (
	"encoding/json"
	"net/http"
	"strconv"

	"goji.io/pat"

	"github.com/py150504/billingps/src/global"
	"golang.org/x/net/context"
)

// Read : read units list
func Read(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	units, errGet := getUnits()
	if errGet != nil {

	}
	var data []interface{}
	for _, unit := range units {
		data = append(data, MapUnit(unit, true))
	}
	response := global.Response{
		Links: r.URL.Path,
		Data:  data}

	json.NewEncoder(w).Encode(response)
}

// ReadDetail : read spesific people from id
func ReadDetail(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/vnd.api+json")

	idString := pat.Param(ctx, "id")
	id, _ := strconv.ParseInt(idString, 10, 64)
	unit, errPerson := getUnit(id)
	if errPerson != nil {

	}

	response := global.Response{
		Links: r.URL.Path,
		Data:  MapUnit(unit, true)}

	json.NewEncoder(w).Encode(response)
}
