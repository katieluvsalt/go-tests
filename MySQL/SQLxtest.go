package main

import (
  //"database/sql";
  "github.com/jmoiron/sqlx";
  _ "github.com/go-sql-driver/mysql";
  "net/http";
  "fmt"
  )


var db *sqlx.DB
var err error

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

type Rows struct {
  Username string
  Password string
}


func katieFunction() {
  rows, _ := db.Query("SELECT username, password FROM user")
  cols, _ := rows.Columns()
  defer rows.Close()
  for rows.Next() {
    //var user string
    //var pswd string
    // Create a slice of interface{}'s to represent each column,
    // and a second slice to contain pointers to each item in the columns slice.
    columns := make([]interface{}, len(cols))
    columnPointers := make([]interface{}, len(cols))
    for i, _ := range columns {
        columnPointers[i] = &columns[i]
    }
    // Scan the result into the column pointers...
    if err := rows.Scan(columnPointers...); err != nil {
        panic(err.Error())
    }
    // Create our map, and retrieve the value for each column from the pointers slice,
    // storing it in the map with the name of the column as the key.
    m := make(map[string]interface{})
    for i, colName := range cols {
        val := columnPointers[i].(*interface{})
        m[colName] = *val
    }
    //err = rows.Scan(&user, &pswd)
    //fmt.Println("User: %s Pswd: %s", user, pswd)
    //fmt.Printf("%s", m)
    for _, value := range m {
      fmt.Printf("%s\n", value)
    }
  }
  // Outputs: map[columnName:value columnName2:value2 columnName3:value3 ...]
  fmt.Println("katieFunction end")
}



func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
  fmt.Println("9 - get homepage")
}


func main() {
  fmt.Println("1 - main")
	db, err = sqlx.Open("mysql", "root:MySQL123@tcp(127.0.0.2:3306)/new_schema")
  fmt.Println("2 - sqlx")
	if err != nil {
		panic(err.Error())
    fmt.Println("3 - error")
	}
	defer db.Close()
  fmt.Println("4 - close")

	err = db.Ping()
  fmt.Println("5 - ping")
	if err != nil {
		panic(err.Error())
    fmt.Println("6 - error")
	}

  katieFunction()

	/*http.HandleFunc("/signup", signupPage)
  fmt.Println("7 - signup")*/
	http.HandleFunc("/login", loginPage)
  fmt.Println("8 - login")
	http.HandleFunc("/", homePage)
  fmt.Println("9 - home")
	http.ListenAndServe(":8080", nil)
  fmt.Println("10 - port")
}
