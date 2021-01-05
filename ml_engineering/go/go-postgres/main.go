// https://medium.com/swlh/building-a-restful-api-with-go-and-postgresql-494819f51810
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq" // https://stackoverflow.com/a/52791919/256002
)

type Person struct {
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
}

const (
	host     = "172.17.0.2"
	port     = 5432
	user     = "postgres"
	password = "mypassword"
	dbname   = "postgres"
)

func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	fmt.Println(psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func GETHandler(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()
	rows, err := db.Query("SELECT * FROM person")
	if err != nil {
		log.Fatal(err)
	}
	var people []Person
	for rows.Next() {
		var person Person
		rows.Scan(&person.Name, &person.Nickname)
		people = append(people, person)
	}
	peopleBytes, _ := json.MarshalIndent(people, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(peopleBytes)
	defer rows.Close()
	defer db.Close() // so this is imho wrong, cause we should not open connection for each method
}

func POSTHandler(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	var p Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sqlStatement := `INSERT INTO person(name, nickname) VALUES ($1, $2)`
	_, err = db.Exec(sqlStatement, p.Name, p.Nickname)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	defer db.Close()
}
func main() {
	http.HandleFunc("/", GETHandler)
	http.HandleFunc("/insert", POSTHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// docker pull postgres:latest
// docker run --name go_postgres -e POSTGRES_PASSWORD=mypassword -d postgres
// docker exec -it go_postgres psql -U postgres
