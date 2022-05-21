package main

import (
	"log"
	"net/http"

	"go.mod/main/users"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.GET("/", users.GetAllUsers).Methods("GET")
	router.POST("/", users.InsertUser).Methods("POST")
	router.GET("/url_post", users.GetPost).Methods("get")
	router.POST("/save_post", users.CreatePost).Methods("POST")
	router.HandleFunc("/change/{id:[0-9]+}", users.UpdateUser).Methods("PUT")
	router.HandleFunc("/change", users.GetUser).Methods("GET")
	router.HandleFunc("/delete/{id:[0-9]+}", users.DeleteUser).Methods("DELETE")

	log.Println("Start...")
	http.ListenAndServe(":4444", router)
}

// router.HandleFunc("/create", users.Create).Methods("GET")
// router.HandleFunc("/create_post", users.CreateArtHTML).Methods("GET")
// router.HandleFunc("/", users.GetAllUsers).Methods("GET")
// router.HandleFunc("/change/{id:[0-9]+}", users.EditUser).Methods("GET")
