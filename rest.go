package main

import (
	"golang.org/x/net/context"
	rest "github.com/AlexanderChen1989/go-json-rest/rest"
	"log"
	"net/http"
)

// curl -i http://127.0.0.1:8080/


func main() {
	api := rest.NewAPI()
	api.Use(rest.DefaultDevStack...)
	api.SetApp(rest.AppSimple(func(ctx context.Context, w rest.ResponseWriter, r *rest.Request) {
		w.WriteJSON(map[string]string{"Body": "Hello World!"})
	}))
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
