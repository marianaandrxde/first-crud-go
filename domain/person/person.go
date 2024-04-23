package person

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/marianaandrxde/first-crud-go/domain"
)

type Service struct {
	dbFilePath string
	people     domain.People
}

func NewService(dbFilePath string) (Service, error) {
	// verifica se o arquivo existe
	_, err := os.Stat(dbFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			//cria um arquivo vazio
			createEmptyFile(dbFilePath)
		}
	}

	// se não existir, cria o arquivo vazio
	//se existir, leio o arquivo e atualizo a variavel people
}

func createEmptyFile(dbFilePath string) error {
	var people domain.People = domain.People{
		People: []domain.Person{},
	}

	peopleJSON, err := json.Marshal(people)
	if err != nil {
		return fmt.Errorf("Error trying to encode people as JSON?: %s", err.Error())
	}

	err = ioutil.WriteFile(dbFilePath, peopleJSON, 0755)
	if err != nil {
		return fmt.Errorf("Error trying to write to file. Error: %s", err.Error())
	}

	return nil
}
