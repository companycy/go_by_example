package main

import (
	"golang.org/x/net/context"
	rest "github.com/AlexanderChen1989/go-json-rest/rest"
	"log"
	"net/http"
	"sync"
)

// curl demo:

// curl -i -H 'Content-Type: application/json' -d '{"Code":"FR","Name":"France"}' http://127.0.0.1:8080/countries
// curl -i -H 'Content-Type: application/json' -d '{"Code":"US","Name":"United States"}' http://127.0.0.1:8080/countries
// curl -i http://127.0.0.1:8080/countries/FR
// curl -i http://127.0.0.1:8080/countries/US
// curl -i http://127.0.0.1:8080/countries
// curl -i -X DELETE http://127.0.0.1:8080/countries/FR
// curl -i http://127.0.0.1:8080/countries
// curl -i -X DELETE http://127.0.0.1:8080/countries/US
// curl -i http://127.0.0.1:8080/countries

func main() {
	api := rest.NewAPI()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/countries", GetAllCountries),
		rest.Post("/countries", PostCountry),
		// rest.Get("/countries/:code", GetCountry),
		// rest.Delete("/countries/:code", DeleteCountry),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

type Country struct {
	Code string
	Name string
}

var store = map[string]*Country{}

var lock = sync.RWMutex{}

// func GetCountry(ctx context.Context, w rest.ResponseWriter, r *rest.Request) {
// 	code := r.PathParam("code") // ?

// 	lock.RLock()
// 	var country *Country
// 	if store[code] != nil {
// 		country = &Country{}
// 		*country = *store[code]
// 	}
// 	lock.RUnlock()

// 	if country == nil {
// 		rest.NotFound(w, r)
// 		return
// 	}
// 	w.WriteJSON(country)
// }

func GetAllCountries(ctx context.Context, w rest.ResponseWriter, r *rest.Request) {
	lock.RLock()
	countries := make([]Country, len(store))
	i := 0
	for _, country := range store {
		countries[i] = *country
		i++
	}
	lock.RUnlock()
	w.WriteJSON(&countries)
}

func PostCountry(ctx context.Context, w rest.ResponseWriter, r *rest.Request) {
	country := Country{}
	err := r.DecodeJSONPayload(&country)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if country.Code == "" {
		rest.Error(w, "country code required", 400)
		return
	}
	if country.Name == "" {
		rest.Error(w, "country name required", 400)
		return
	}
	lock.Lock()
	store[country.Code] = &country
	lock.Unlock()
	w.WriteJSON(&country)
}

// func DeleteCountry(ctx context.Context, w rest.ResponseWriter, r *rest.Request) {
// 	code := r.PathParam("code")
// 	lock.Lock()
// 	delete(store, code)
// 	lock.Unlock()
// 	w.WriteHeader(http.StatusOK)
// }
