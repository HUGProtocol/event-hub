package restful

import (
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gorilla/mux"
	"event-hub/db"
	"net/http"
	"time"
)

type Service struct {
	port       string
	db         *db.DBService
}

func InitRestService(port string, db *db.DBService) *Service {
	return &Service{
		port:       port,
		db:         db,
	}
}



func (c *Service) Start() error {
	log.Info("start queryer rpc port:" + c.port)
	address := "0.0.0.0:" + c.port
	r := mux.NewRouter()

	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

	})

	//jwt token name profile
	r.HandleFunc("/register", func(writer http.ResponseWriter, request *http.Request) {


	})

	r.HandleFunc("/profile_image", func(writer http.ResponseWriter, request *http.Request) {

	})

	r.HandleFunc("/create_event", func(writer http.ResponseWriter, request *http.Request) {

	})

	r.HandleFunc("/review_event", func(writer http.ResponseWriter, request *http.Request) {

	})


	go func() {
		err := http.ListenAndServe(address, r)
		if err != nil {
			utils.Fatalf("http listen error", err)
		}
	}()
	return nil
}

func PrintErrorStr(prefix string, detail string) string {
	if detail != "" {
		return time.Now().String() + " ERROR " + prefix + ":" + detail
	} else {
		return time.Now().String() + " ERROR " + prefix
	}
}


