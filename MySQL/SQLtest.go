package main

import (
  "database/sql";
  _ "github.com/go-sql-driver/mysql";
  "net/http";
  "fmt"
  )


var db *sql.DB
var err error

func signupPage(res http.ResponseWriter, req *http.Request) {
    if req.Method != "POST" {
        http.ServeFile(res, req, "signup.html")
        return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	var user string

    err := db.QueryRow("SELECT username FROM user WHERE username=?", username).Scan(&user)

    switch {
    case err == sql.ErrNoRows:
        _, err = db.Exec("INSERT INTO user(username, password) VALUES(?, ?)", username, password)
        if err != nil {
            panic(err.Error())
            http.Error(res, "Server error, unable to create your account.", 500)
            return
        }

        res.Write([]byte("User created!"))
        return
    case err != nil:
        http.Error(res, "Server error, unable to create your account.", 500)
        return
    default:
        http.Redirect(res, req, "/", 301)
    }
}

func loginPage(res http.ResponseWriter, req *http.Request) {
  fmt.Println("login")
	if req.Method != "POST" {
		http.ServeFile(res, req, "login.html")
    fmt.Println("8 - get login")
		return
	}
fmt.Println("9-login post")
	username := req.FormValue("username")
	//password := req.FormValue("password")

	var databaseUsername string
	var databasePassword string

	err := db.QueryRow("SELECT username, password FROM user WHERE username=?", username).Scan(&databaseUsername, &databasePassword)
  fmt.Println("8a - username", err)

	if err != nil {
    panic(err.Error())
		http.Redirect(res, req, "/login", 301)
    fmt.Println("8b - redirect")
		return
	}

	res.Write([]byte("Hello" + databaseUsername))
  fmt.Println("8c - new")

}

/*type Rows struct {
  username string
  password string
}*/

//Test if can return data fields as string/other type from MySQL
func robFunction() {
  fmt.Println("robFunction starts")

  rows, err := db.Query("SELECT username, password FROM user")
  defer rows.Close()
  for rows.Next() {
    var user string
    var pswd string
    err = rows.Scan(&user, &pswd)
    fmt.Println("User: %s Pswd: %s", user, pswd)
  }
  err = rows.Err() //should be <nil>
  fmt.Println(err, "robFunction end")
}

func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
  fmt.Println("9 - get homepage")
}


func main() {
	db, err = sql.Open("mysql", "root:MySQL123@tcp(127.0.0.2:3306)/new_schema")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

  robFunction()
	http.HandleFunc("/signup", signupPage)
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}
