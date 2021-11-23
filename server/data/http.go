package data

import (
	"SamuraiLab3/server/tools"
	"fmt"
	"log"
	"net/http"
)

func HttpHandler(repository *Repository) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleGetForumsList(repository, rw)
		} else if r.Method == "POST" {
			handleAddUser(r, rw, repository)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleAddUser(r *http.Request, rw http.ResponseWriter, repository *Repository) {
	fmt.Println("handleAddUser invoke")
	// Достать из r.Body шнягу
	var user = User{}
	repository.AddUser(user)
}

func handleGetForumsList(repository *Repository, rw http.ResponseWriter) {
	fmt.Println("handleGetForumsList invoke")

	res, err := repository.GetForumsList()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}
