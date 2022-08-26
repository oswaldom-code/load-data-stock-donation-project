package services

import (
	"io/ioutil"
	"os"

	"github.com/oswaldom-code/load-data-stock-donation-project/src/adapters/persistence/repository"
	"github.com/oswaldom-code/load-data-stock-donation-project/src/services/ports"
)

type LoadDataServices interface {
	TestDbConnection() error
	LoadDataFromJsonFile(jsonFile string) (int64, error)
	LoadDataFromDirectory(path string) (int64, error)
}

type loadDataServices struct {
	r ports.Repository
}

func NewLoadDataServices() LoadDataServices {
	return &loadDataServices{r: repository.NewRepository()}
}

func (s *loadDataServices) TestDbConnection() error {
	return s.r.TestDb()
}

func (s *loadDataServices) LoadDataFromJsonFile(jsonFile string) (int64, error) {
	// open json file
	f, err := os.Open(jsonFile)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	// read json file
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return 0, err
	}
	// unmarshal json file
	fields, err := UnmarshalField(content)
	if err != nil {
		return 0, err
	}
	// insert data into database
	return s.r.InsertData(fields)
}

func (s *loadDataServices) LoadDataFromDirectory(path string) (int64, error) {
	// TODO: implement the load for each json file in the directory
	return 0, nil
}
