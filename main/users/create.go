package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go.mod/main/connecting"
)

var DB = connecting.InitDB()

func InsertUser(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	var err error

	name := r.FormValue("name")
	surname := r.FormValue("surname")

	insert := `insert into "User"("name","surname")values($1,$2)`
	_, err = DB.Exec(insert, name, surname)
	if err != nil {
		http.NotFound(w, r)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var err error
	title := r.FormValue("title")
	text := r.FormValue("text")
	insertArticles := `insert into "Articles"("title","text")values($1,$2)`
	_, err = DB.Exec(insertArticles, title, text)
	if err != nil {
		http.NotFound(w, r)
	}
	// http.Redirect(w, r, "/", http.StatusFound)
}

// func CreateArtHTML(w http.ResponseWriter, r *http.Request) {
// 	filepath := filepath.Join("template", "newArticls.html")
// 	html, err := template.ParseFiles(filepath)
// 	if err != nil {
// 		http.NotFound(w, r)
// 	}
// 	html.Execute(w, nil)
// }

// func Create(w http.ResponseWriter, r *http.Request) {
// 	filepath := filepath.Join("template", "create.html")
// 	html, err := template.ParseFiles(filepath)
// 	if err != nil {
// 		http.NotFound(w, r)
// 	}
// 	html.Execute(w, nil)
// }
