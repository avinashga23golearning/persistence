package persistence

import (
	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/avinashga23golearning/model"
)

var client driver.Client

func init() {
	client, _ = newDriver()
}

func newDriver() (driver.Client, error) {
	conn, _ := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://192.168.100.10:8259"},
	})

	return driver.NewClient(driver.ClientConfig{
		Connection: conn,
	})
}

//CreatePerson creates person
func CreatePerson(person model.Person) string {
	person.ID = ""

	db, _ := client.Database(nil, "_system")
	col, _ := db.Collection(nil, "person")

	meta, _ := col.CreateDocument(nil, person)
	return meta.Key
}

//DeletePerson deletes person
func DeletePerson(id string) {
	db, _ := client.Database(nil, "_system")
	col, _ := db.Collection(nil, "person")

	col.RemoveDocument(nil, id)
}

//UpdatePerson updates person
func UpdatePerson(person model.Person) {
	db, _ := client.Database(nil, "_system")
	col, _ := db.Collection(nil, "person")

	col.UpdateDocument(nil, person.ID, person)
}

//GetPersonByID gets person by id
func GetPersonByID(id string) model.Person {
	db, _ := client.Database(nil, "_system")
	col, _ := db.Collection(nil, "person")
	var person model.Person

	col.ReadDocument(nil, id, &person)
	return person
}
