package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	database "github.com/3D-ShowCaser/backend/internal"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var PUBLIC_URL = []string{"/", "/auth"}

func Is_URL_Public(url string) bool {
	for _, i := range PUBLIC_URL {
		if strings.Contains(url, i) {
			return true
		}
	}
	return false
}

type DbAPi struct {
	queries *database.Queries
}

var db = DbAPi{}

func main() {
	godotenv.Load()
	queries, err := ConnectToDB()
	db = DbAPi{
		queries: queries,
	}
	if err != nil {
		log.Fatal("Unable to connect to database")
	}
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler)
	router.Handle("/auth", GetAuthRouter(router, &db))
	router.Use(db.authenticatedMiddleWare)
	router.HandleFunc("/profile", ProfileHandler)
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router)
}
