package DocumentDao

import (
	"os"
	"encoding/json"
	"io/ioutil"
	"log"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(2)
}

type DaoInstance struct{}

func (dao DaoInstance) Get(id string) (map[string]interface{}, error) {
	log.Println("Starting get from file system")

	var contents []byte
	var err error

	if contents, err = ioutil.ReadFile("jsons/" + id + ".json"); err != nil {
		log.Println("Error reading file")
		return nil, err
	}

	var out map[string]interface{}

	if err = json.Unmarshal(contents, &out); err != nil {
		log.Println("Error unmashalling file contents")
		log.Print("Contents were: ")
		log.Println(contents)
		return nil, err
	}

	log.Println("Get from file system ended normally")

	return out, nil
}

func (dao DaoInstance) Search() ([]map[string]interface{}, error) {
	log.Println("Starting search")

	var out = []map[string]interface{}{}

	var files []os.FileInfo
	var err error

	if files, err = ioutil.ReadDir("jsons/"); err != nil {
		log.Println("Error reading directory")
		return nil, err
	}

	size := len(files)

	channel := make(chan map[string]interface{})

	for _, value := range files {
		go func (value os.FileInfo, channel chan map[string]interface{}) {
			log.Println("Trying to read file ", value.Name())

			var contents []byte
			if contents, err = ioutil.ReadFile("jsons/" + value.Name()); err != nil {
				log.Println("Error reading file", value.Name())
			}

			var current map[string]interface{}
			if err = json.Unmarshal(contents, &current); err != nil {
				log.Println("Error unmarshalling file", value.Name())
			}

			channel <- current

			log.Println("File ", value.Name(), " was read normally")
		}(value, channel)
	}

	for i := 0; i < size; i ++ {
		next := <- channel
		out = append(out, next)
	}

	log.Println("Search ended normally")
	return out, nil
}

func (dao DaoInstance) Create(id string, doc map[string]interface{}) (map[string]interface{}, error) {
	var file *os.File
	var err error

	if file, err = os.Create("jsons/" + id + ".json"); err != nil {
		return nil, err
	}

	var jsonBytes []byte

	if jsonBytes, err = json.Marshal(doc); err != nil {
		return nil, err
	}

	if _, err = file.Write(jsonBytes); err != nil {
		return nil, err;
	}


	return doc, nil
}

func New() *DaoInstance {
	return new(DaoInstance)
}