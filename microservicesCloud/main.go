package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    //"strings"
    "github.com/katieluvsalt/microservicesCloud/dal"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome, %s!", r.URL.Path[1:])
}

func main() {
    dal.InitDB();
    http.HandleFunc("/", handler)
    http.HandleFunc("/about/", about)
    http.HandleFunc("/users/", users)
    http.ListenAndServe(":8080", nil)
}

type Message struct {
    Text string
}

func users(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
      case "GET":
        // Serve the resource.
        get, err := dal.GetUser()

        if err != nil {
      		fmt.Println(err.Error())
          w.WriteHeader(http.StatusInternalServerError)
          w.Write([]byte(err.Error()))
        }else{
          w.Header().Set("Content-Type", "application/json")
          w.Write([]byte(get))
        }
      case "POST":
        // Create a new record.
        dal.PostUser(r)
      case "PUT":
        // Update an existing record.
        w.WriteHeader(http.StatusMethodNotAllowed)
      case "DELETE":
        // Remove the record.
        w.WriteHeader(http.StatusMethodNotAllowed)
      default:
        // Give an error message.
        fmt.Println("There is an ERROR in request type")
      }
  }

func about(w http.ResponseWriter, r *http.Request) {
    m := Message{"Welcome to the SandovalEffect API, build v0.0.001.992, 6/22/2015 0340 UTC."}
    b, err := json.Marshal(m)
    if err != nil {
        panic(err)
    }
     w.Write(b)
}
