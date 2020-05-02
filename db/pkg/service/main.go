package dbservice

import (
	"gopkg.in/mgo.v2"
	"net/http"
)

func main() {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	dbService := NewDbService(session)
	endpoints := NewEndpoints(dbService)
	handler := NewHTTPHandler(endpoints)

	// Replace ":8080" with your desired port number
	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		return
	}
}
