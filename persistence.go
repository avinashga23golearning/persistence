package persistence

import (
	"log"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/avinashga23golearning/model"
)

var personCollection driver.Collection

func init() {
	personCollection, _ = getPersonCollection()
}

//PersonPersistenceManager type
type PersonPersistenceManager struct{}

//NewPersonPersistenceManager new person persistence manager
func NewPersonPersistenceManager() *PersonPersistenceManager {
	personPersistenceManager := PersonPersistenceManager{}

	return &personPersistenceManager
}

func getPersonCollection() (driver.Collection, error) {
	auth := driver.BasicAuthentication("root", "Mpv0DgupY7cgD53e")
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://192.168.100.10:8529"},
	})
	if err != nil {
		log.Panic(err)
	}

	conn.SetAuthentication(auth)

	client, err := driver.NewClient(driver.ClientConfig{
		Connection: conn,
	})
	if err != nil {
		log.Panic(err)
	}

	db, err := client.Database(nil, "_system")
	if err != nil {
		log.Panic(err)
	}

	return db.Collection(nil, "person")
}

//CreatePerson creates person
func (PersonPersistenceManager) CreatePerson(person model.Person) string {
	person.ID = ""

	meta, _ := personCollection.CreateDocument(nil, person)
	return meta.Key
}

//DeletePerson deletes person
func (PersonPersistenceManager) DeletePerson(id string) {
	personCollection.RemoveDocument(nil, id)
}

//UpdatePerson updates person
func (PersonPersistenceManager) UpdatePerson(person model.Person) {
	personCollection.UpdateDocument(nil, person.ID, person)
}

//GetPersonByID gets person by id
func (PersonPersistenceManager) GetPersonByID(id string) (model.Person, bool) {
	var person model.Person
	personNotFound := false

	meta, err := personCollection.ReadDocument(nil, id, &person)
	if err != nil {
		personNotFound = driver.IsNotFound(err)
		if personNotFound {
			log.Printf("person by id %s does not exists\n", id)
		} else {
			log.Fatal(err)
		}
	}

	person.ID = meta.Key
	return person, personNotFound
}
