package people

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/py150504/billingps/src/global"
)

// Person : data type people
type Person struct {
	ID       int64     `schema:"-" json:"-"`
	Name     string    `schema:"name" json:"name"`
	Phone    string    `schema:"phone" json:"phone"`
	JoinDate time.Time `schema:"-" json:"-"`
	Status   int       `schema:"-" json:"-"`
}

var queryPerson preparedQueryPerson

type preparedQueryPerson struct {
	selectPeople *sql.Stmt
	selectPerson *sql.Stmt
	insertPerson *sql.Stmt
	deletePerson *sql.Stmt
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
	queryPerson.insertPerson, errPrepared = db.Prepare(`
		INSERT INTO 
			people (name, phone, join_date, status)
		VALUES 
			(?, ?, ?, ?)`)

	if errPrepared != nil {
		log.Printf("Error prepare insert person : %s", errPrepared.Error())
		log.Fatal("App exit, fail Init People")
	}
	queryPerson.deletePerson, errPrepared = db.Prepare(`
		UPDATE people
		SET
			status = 0
		WHERE
			id = ?`)

	if errPrepared != nil {
		log.Printf("Error prepare delete person : %s", errPrepared.Error())
		log.Fatal("App exit, fail Init People")
	}
	return nil
}

func (p *Person) save() error {
	p.JoinDate = time.Now()
	p.Status = 1
	resultInsert, errInsert := queryPerson.insertPerson.Exec(p.Name, p.Phone, p.JoinDate, p.Status)
	if errInsert != nil {
		log.Printf(errInsert.Error())
		return nil
	}
	lastID, errResult := resultInsert.LastInsertId()
	if errResult != nil {
		log.Printf(errResult.Error())
		return nil
	}
	p.ID = lastID

	return nil
}

func (p *Person) load() error {
	errSelect := queryPerson.selectPerson.QueryRow(p.ID).Scan(
		&p.ID,
		&p.Name,
		&p.Phone,
		&p.JoinDate)

	if errSelect != nil {
		log.Printf(errSelect.Error())
		return nil
	}

	return nil
}

func (p *Person) delete() error {
	resultDelete, errDelete := queryPerson.deletePerson.Exec(p.ID)
	if errDelete != nil {
		log.Printf(errDelete.Error())
		return nil
	}
	affectedRow, errResult := resultDelete.RowsAffected()
	if errResult != nil {
		log.Printf(errResult.Error())
		return nil
	}
	if affectedRow == 0 {
		log.Printf("affected row: %d", affectedRow)
		return nil
	}
	return nil
}

func getPerson(id int64) *Person {
	person := new(Person)
	person.ID = id
	person.load()

	return person
}

func getPeople() []*Person {
	people := []*Person{}
	// db := global.DB.Core
	// selectQuery, errPrepared := db.Prepare(`
	// 	SELECT
	// 		id, name, phone, join_date
	// 	FROM
	// 		people
	// 	WHERE
	// 		status = 1`)
	// if errPrepared != nil {
	// 	log.Printf(errPrepared.Error())
	// }
	rows, errSelect := queryPerson.selectPeople.Query()

	if errSelect != nil {
		log.Printf(errSelect.Error())
	}
	defer rows.Close()
	for rows.Next() {
		person := new(Person)
		errScan := rows.Scan(
			&person.ID,
			&person.Name,
			&person.Phone,
			&person.JoinDate)
		if errScan != nil {
			log.Printf(errScan.Error())
		}
		people = append(people, person)
	}
	// queryPerson.selectPeople.Close()
	return people
}

// Read : read people from id
func Read(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	people := getPeople()
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
func ReadDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	vars := mux.Vars(r)
	idString := vars["id"]
	id, _ := strconv.ParseInt(idString, 10, 64)
	person := getPerson(id)

	response := global.Response{
		Links: r.URL.Path,
		Data:  MapPerson(person, true)}

	json.NewEncoder(w).Encode(response)
}

// Create : create people from input data
func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	errParse := r.ParseForm()
	if errParse != nil {
		log.Printf(errParse.Error())
		return
	}

	people := new(Person)
	decoder := schema.NewDecoder()
	errDecode := decoder.Decode(people, r.PostForm)
	if errDecode != nil {
		log.Printf(errDecode.Error())
		return
	}

	errSave := people.save()
	if errSave != nil {
		log.Printf(errSave.Error())
		return
	}
	response := MapPerson(people, true)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// Delete : delete people from id
func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	vars := mux.Vars(r)
	idString := vars["id"]
	id, _ := strconv.ParseInt(idString, 10, 64)
	person := new(Person)
	person.ID = id
	person.delete()

	response := global.Response{
		Data: MapPerson(person, false)}

	json.NewEncoder(w).Encode(response)
}

// MapPerson : map person output as jsonapi.org
func MapPerson(p *Person, detail bool) interface{} {
	var attributes interface{}
	if detail {
		attributes = map[string]interface{}{
			"name":      p.Name,
			"phone":     p.Phone,
			"join_date": p.JoinDate.Format("02 January 2006, 15:04"),
		}
	}
	person := map[string]interface{}{
		"id":         strconv.FormatInt(p.ID, 10),
		"type":       "person",
		"attributes": attributes,
	}
	return person
}
