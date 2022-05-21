package users

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"go.mod/main/connecting"
)

type Users struct {
	Id      int
	Name    string
	Surname string
}

type Articles struct {
	Id    int
	Title string
	Text  string
}

var Db = connecting.InitDB()

func GetUser(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	w.Write([]byte("hello"))
}

func GetPost(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	w.Write([]byte("wlcome to posting"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	// Db := connecting.InitDB()
	vars := mux.Vars(r)
	id := vars["id"]
	name := r.FormValue("name")
	surname := r.FormValue("surname")
	update := `UPDATE "User" set  "name"=$1, "surname"=$2 where "id"=$3`
	_, err := Db.Exec(update, name, surname, id)
	check(err)
	if err != nil {
		http.NotFound(w, r)
	}
	http.Redirect(w, r, "/", 302)
}

func DeleteUser(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	// Db := connecting.InitDB()
	vars := mux.Vars(r)
	id := vars["id"]
	_, err := Db.Exec(`DELETE  FROM "User" where id = $1`, id)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/", 302)
}

func check(err error) {
	if err != nil {
		log.Println(err)
	}
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Db := connecting.InitDB()
	var err error
	rows, err := Db.Query(`SELECT * FROM "User"`)
	check(err)

	// users := []Users{}
	for rows.Next() {
		u := Users{}
		err := rows.Scan(&u.Id, &u.Name, &u.Surname)
		check(err)
		// users = append(users, u)
		w.Write([]byte(fmt.Sprintf("Id:%d\nName:%s\nSurname:%s\n", u.Id, u.Name, u.Surname)))
	}

	w.Write([]byte("\n<h1>НОВОСТИ</h1>\n"))

	rowsPost, _ := Db.Query(`SELECT * FROM "Articles"`)

	for rowsPost.Next() {
		a := Articles{}
		err = rowsPost.Scan(&a.Id, &a.Title, &a.Text)
		check(err)
		w.Write([]byte(fmt.Sprintf("Id:%d\nTitle:%s\nText:%s\n", a.Id, a.Title, a.Text)))

	}

}

// func EditUser(w http.ResponseWriter, r *http.Request) {
// Db := connecting.InitDB()

// 	vars := mux.Vars(r)
// 	id := vars["id"]
// 	getting := `select from "User" where id=$1`

// 	row := Db.QueryRow(getting, &id)
// 	users := Users{}
// 	err := row.Scan(&users.Id, &users.Name, &users.Surname)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	filepath := filepath.Join("template", "edit.html")
// 	tmpl, err := template.ParseFiles(filepath)
// 	log.Println(err)

// 	tmpl.Execute(w, users)
// }
