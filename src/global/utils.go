package global

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// LogError : error logger
var LogError *log.Logger

// Response : data type responses
type Response struct {
	Links  interface{} `json:"links"`
	Data   interface{} `json:"data"`
	Errors interface{} `json:"errors"`
}

// NotFound : handle 404.
func NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "custom 404")
}

// InitGlobal : init global func
func InitGlobal(errorHandle io.Writer) {
	LogError = log.New(errorHandle,
		"ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// FailResponse : respond with error
func FailResponse(w http.ResponseWriter, err error) {
	response := Response{
		Errors: err.Error()}
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(response)
}
