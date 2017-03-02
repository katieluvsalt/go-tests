package dal

import (
	"database/sql"
  "fmt"
	_ "github.com/go-sql-driver/mysql"
  "github.com/katieluvsalt/microservicesCloud/model"
  "encoding/json"
  "net/http"
)

var db *sql.DB

func InitDB() error {
  var err error
	db, err = sql.Open("mysql",
		"admin:JVEGNYSQVKTDDYNP@tcp(sl-us-dal-9-portal.2.dblayer.com:19876)/microservicesTest")
	if err != nil {
		fmt.Println(err.Error())
    return err
	}
  /*if err = db.Ping(); err != nil {
      fmt.Println(err.Error())
      return err
  }*/
  return nil
}

func PostUser(r *http.Request) (model.User,error){
  decoder := json.NewDecoder(r.Body)
  var u model.User
  err := decoder.Decode(&u)
  if err != nil {
     fmt.Println(err.Error())
     return  u, err
   }
  err = insertUser(u)
  if err != nil {
     fmt.Println(err.Error())
     return  u, err
   }
  defer r.Body.Close()
  return u, err
}

func insertUser(user model.User) error{
  fmt.Printf("The following will be added to the database: %v\n" , user)
  //db.Ping()
  _, err := db.Exec("INSERT INTO user (username, password, email) VALUES (?, ?, ?)",user.Username, user.Password, user.Email )
  if err != nil {
      fmt.Println(err.Error())
      return err;
  }
  //defer db.Close()
  return nil
}

func GetUser1() error{
  rows, _ := db.Query("SELECT * FROM user")
  cols, _ := rows.Columns()
  defer rows.Close()
  m := make(map[string]interface{})
  for rows.Next() {
    columns := make([]interface{}, len(cols))
    columnPointers := make([]interface{}, len(cols))
    for i, _ := range columns {
        columnPointers[i] = &columns[i]
    }
    if err := rows.Scan(columnPointers...); err != nil {
      fmt.Println(err.Error())
      return err;
    }
    for i, colName := range cols {
        val := columnPointers[i].(*interface{})
        m[colName] = *val
		//test
		fmt.Println(m)
    }
    for _, value := range m {
      jsonData, err := json.Marshal(value)
      fmt.Printf("%s\n", jsonData)
      if err != nil {
        fmt.Println(err.Error())
        return err;
      }
    }
  }
  return nil
}

func GetUser() (string, error){
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		fmt.Println(err.Error())
		return "", err;
	}
	defer rows.Close()
	//create empty slice
	var a []model.User
	for rows.Next() {
		foundUser := new(model.User)
		if err := rows.Scan(
			&foundUser.Username, &foundUser.Password, &foundUser.Email,
		); err != nil {
			fmt.Println(err.Error())
			return "",err;
		}
		//append row to slice
		a = append(a, *foundUser)
	}
	//print database
	fmt.Println(a)
	//encode database
	k, err := encodeUser(a)
	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return "",err;
	}
	return k, nil
}

func encodeUser(data []model.User) (string,error){
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return "", err;
	}
	fmt.Printf("%s",b)
	return string(b), nil
}
