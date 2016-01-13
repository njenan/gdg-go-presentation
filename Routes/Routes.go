package Routes

import (
	"net/http"
	"log"
	"encoding/json"
	"math/rand"
	"strconv"

	"../DocumentDao"
)

var methodHandlers map[string]MethodHandler
var Dao *DocumentDao.DaoInstance

type MethodHandler func(id string, body map[string]interface{}) (interface{}, error)

func post(id string, body map[string]interface{}) (interface{}, error) {
	id = strconv.Itoa(rand.Intn(10000000))
	body["__id"] = id

	_, err := Dao.Create(id, body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func get(id string, body map[string]interface{}) (interface{}, error) {
	var err error
	var out interface{}

	//log.Println("body is {}", out)

	if id != "" {
		if out, err = Dao.Get(id); err != nil {
			log.Println("Error getting from dao")
			return nil, err
		}
	} else {
		if out, err = Dao.Search(); err != nil {
			log.Println("Error searching from dao")
			return nil, err
		}
	}

	//log.Println("body is now {}", out)

	return out, nil
}

func init() {
	Dao = DocumentDao.New()

	methodHandlers = make(map[string]MethodHandler)
	methodHandlers["POST"] = post
	methodHandlers["GET"] = get
}

func dispatch(writer http.ResponseWriter, req *http.Request) {
	log.Println("Starting dispatch")

	var out []byte
	var body map[string]interface{}

	method := req.Method

	id := req.URL.Path[1:]

	var err error

	log.Println(req.Body)

	if req.Method == "POST" || req.Method == "PUT" {
		err = json.NewDecoder(req.Body).Decode(&body)
	}

	if err != nil {
		log.Println("Error decoding body")
		writer.Write([]byte(err.Error()))
		return
	}

	var response interface{}

	if response, err = methodHandlers[method](id, body); err != nil {
		log.Println("Error executing method handler")
		writer.Write([]byte(err.Error()))
		return
	}

	if out, err = json.Marshal(response); err != nil {
		log.Println("Error mashalling response object")
		writer.Write([]byte(err.Error()))
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(out)
	log.Println("Dispatch ended normally")
}

func SetUpRoutes() {
	http.HandleFunc("/", dispatch)

	log.Print("Server listening at http://localhost:8080/")
	http.ListenAndServe(":8080", nil);

}