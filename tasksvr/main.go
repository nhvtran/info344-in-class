package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nhvtran/info344-in-class/tasksvr/handlers"
	"github.com/nhvtran/info344-in-class/tasksvr/models/tasks"

	mgo "gopkg.in/mgo.v2"
)

const defaultAddr = ":80"

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = defaultAddr
	}

	//TODO: make connection to the DBMS
	//construct the appropriate tasks.Store
	//construct the handlers.Context

	// note: because I'm using Windows Home, use 192.168.99.100 instead of localhost because of the Linux VIM
	// another note: mongoDB is at 192.168.99.100:27017;
	// in reality, you're supposed to use localhost for running Go API server
	mongoSess, err := mgo.Dial("192.168.99.100")
	if err != nil {
		log.Fatalf("error dialing mongo: %v", err)
	}
	mongoStore := tasks.NewMongoStore(mongoSess, "tasks", "tasks")

	handlerCtx := handlers.NewHandlerContext(mongoStore)

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/tasks", handlerCtx.TasksHandler)
	mux.HandleFunc("/v1/tasks/", handlerCtx.SpecificTaskHandler)

	fmt.Printf("server is listening at http://%s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
