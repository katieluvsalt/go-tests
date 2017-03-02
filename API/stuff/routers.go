package API

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	// Route{
	// 	"Index",
	// 	"GET",
	// 	"/amc/tbd/v0/applicants/",
	// 	Index,
	// },

	Route{
		"AddApplicant",
		"POST",
		"/amc/tbd/v0/applicants/",
		AddApplicant,
	},

	Route{
		"DeleteApplicant",
		"DELETE",
		"/amc/tbd/v0/applicants/{id}",
		DeleteApplicant,
	},

	Route{
		"FindApplicantById",
		"GET",
		"/amc/tbd/v0/applicants/{id}",
		FindApplicantById,
	},

	Route{
		"FindApplicants",
		"GET",
		"/amc/tbd/v0/applicants/",
		FindApplicants,
	},

}
