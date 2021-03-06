package people

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/py150504/billingps/src/global"
	"goji.io/pat"
	"golang.org/x/net/context"
)

// Read : read people from id
func Read(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	people, errGet := getPeople()
	if errGet != nil {
		global.LogError.Printf(errGet.Error())
		global.FailResponse(w, errGet)
	}
	var data []interface{}
	for _, person := range people {
		data = append(data, MapPerson(person, true))
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
	person, errGet := getPerson(id)
	if errGet != nil {
		global.LogError.Printf(errGet.Error())
		global.FailResponse(w, errGet)
	}
	response := global.Response{
		Links: r.URL.Path,
		Data:  MapPerson(person, true)}

	json.NewEncoder(w).Encode(response)
}

// Create : create people from input data
func Create(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/vnd.api+json")

	errParse := r.ParseForm()
	if errParse != nil {
		log.Printf(errParse.Error())
		return
	}

	person := new(Person)
	decoder := json.NewDecoder(r.Body)
	errDecode := decoder.Decode(&person)
	if errDecode != nil {
		global.LogError.Printf(errDecode.Error())
		global.FailResponse(w, errDecode)
		return
	}

	errSave := person.save()
	if errSave != nil {
		global.LogError.Printf(errSave.Error())
		global.FailResponse(w, errSave)
		return
	}
	response := global.Response{
		Data: MapPerson(person, true)}

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
	person := new(Person)
	person.ID = id
	errDelete := person.delete()
	if errDelete != nil {
		global.LogError.Printf(errDelete.Error())
		global.FailResponse(w, errDelete)
		return
	}
	response := global.Response{
		Data: MapPerson(person, false)}

	json.NewEncoder(w).Encode(response)
}

// Update : update people from input data
func Update(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/vnd.api+json")

	errParse := r.ParseForm()
	if errParse != nil {
		log.Printf(errParse.Error())
		return
	}

	person := new(Person)
	decoder := json.NewDecoder(r.Body)
	errDecode := decoder.Decode(&person)
	if errDecode != nil {
		log.Printf(errDecode.Error())
		return
	}
	errUpdate := person.update()
	if errUpdate != nil {
		log.Printf(errUpdate.Error())
		return
	}
	response := global.Response{
		Data: MapPerson(person, true)}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
