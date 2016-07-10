package main

import (
	"log"
	"net/http"
	"os"

	"goji.io"
	"goji.io/pat"

	"github.com/py150504/billingps/src/global"
	"github.com/py150504/billingps/src/people"
	"github.com/py150504/billingps/src/units"
	"github.com/rs/cors"
)

func main() {
	global.InitDB()
	global.InitGlobal(os.Stderr)
	people.InitPeople()
	errUnit := units.InitUnit()
	if errUnit != nil {
		global.LogError.Fatalf(errUnit.Error())
	}
	log.Println("Run on : http://localhost:8080")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://www.billing.com:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PATCH"},
	})

	mux := goji.NewMux()
	mux.Use(corsHandler.Handler)

	mux.HandleFuncC(pat.Get("/people"), people.Read)
	mux.HandleFuncC(pat.Get("/people/:id"), people.ReadDetail)
	mux.HandleFuncC(pat.Post("/people"), people.Create)
	mux.HandleFuncC(pat.Delete("/people/:id"), people.Delete)
	mux.HandleFuncC(pat.Patch("/people/:id"), people.Update)

	mux.HandleFuncC(pat.Get("/units"), units.Read)
	mux.HandleFuncC(pat.Get("/units/:id"), units.ReadDetail)
	mux.HandleFuncC(pat.Post("/units"), units.Create)
	mux.HandleFuncC(pat.Delete("/units/:id"), units.Delete)
	mux.HandleFuncC(pat.Patch("/units/:id"), units.Update)

	http.ListenAndServe(":8080", mux)
}
