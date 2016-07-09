package global

import (
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
