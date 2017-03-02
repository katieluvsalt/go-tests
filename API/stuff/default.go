package API

import (
	"net/http"
	"fmt"
	"github.com/katieluvsalt/microservicesCloud/dal"
)

type Default struct {

}

func AddApplicant(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

func DeleteApplicant(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

func FindApplicantById(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

func FindApplicants(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		get, err := dal.GetUser()

		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}else{
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(get))
		}
		w.WriteHeader(http.StatusOK)
}
