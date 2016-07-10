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
		global.FailResponse(w, errGet)
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
	id, errParse := strconv.ParseInt(idString, 10, 64)
	if errParse != nil {
		global.LogError.Printf(errParse.Error())
		global.FailResponse(w, errParse)
		return
	}
	unit, errPerson := getUnit(id)
	if errPerson != nil {
		global.FailResponse(w, errPerson)
	}

	response := global.Response{
		Links: r.URL.Path,
		Data:  MapUnit(unit, true)}

	json.NewEncoder(w).Encode(response)
}

// Create : create people from input data
func Create(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/vnd.api+json")

	errParse := r.ParseForm()
	if errParse != nil {
		global.LogError.Printf(errParse.Error())
		return
	}

	unit := new(Unit)
	decoder := json.NewDecoder(r.Body)
	errDecode := decoder.Decode(&unit)
	if errDecode != nil {
		global.LogError.Printf(errDecode.Error())
		global.FailResponse(w, errDecode)
		return
	}

	errSave := unit.save()
	if errSave != nil {
		global.FailResponse(w, errSave)
		return
	}
	response := global.Response{
		Data: MapUnit(unit, true)}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// Delete : delete people from id
func Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/vnd.api+json")

	idString := pat.Param(ctx, "id")
	id, errParse := strconv.ParseInt(idString, 10, 64)
	if errParse != nil {
		global.LogError.Printf(errParse.Error())
		global.FailResponse(w, errParse)
		return
	}
	unit := new(Unit)
	unit.ID = id
	errDelete := unit.delete()
	if errDelete != nil {
		global.FailResponse(w, errDelete)
		return
	}
	response := global.Response{
		Data: MapUnit(unit, false)}

	json.NewEncoder(w).Encode(response)
}

// Update : update people from input data
func Update(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/vnd.api+json")

	errParse := r.ParseForm()
	if errParse != nil {
		global.LogError.Printf(errParse.Error())
		return
	}
	idString := pat.Param(ctx, "id")
	id, errParse := strconv.ParseInt(idString, 10, 64)
	if errParse != nil {
		global.LogError.Printf(errParse.Error())
		global.FailResponse(w, errParse)
		return
	}

	unit := new(Unit)
	decoder := json.NewDecoder(r.Body)
	errDecode := decoder.Decode(&unit)
	if errDecode != nil {
		global.LogError.Printf(errDecode.Error())
		global.FailResponse(w, errDecode)
		return
	}
	unit.ID = id
	errUpdate := unit.update()
	if errUpdate != nil {
		global.FailResponse(w, errUpdate)
		return
	}
	response := global.Response{
		Data: MapUnit(unit, true)}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
