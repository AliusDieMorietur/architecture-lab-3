package data

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
)

type Forum struct {
	Id 					 int64
	Name         string
	TopicKeyword string
	Users        []string
}

type User struct {
	Id 				int64
	Name      string
	Interests []string
}

type Repository struct {
	Db *sql.DB
}

func (r *Repository) GetForumsList() ([]*Forum, error) {
	fmt.Printf("Get forums list")
	rows, err :=  r.Db.Query(`
		SELECT 
			f.*,
			(
				SELECT array_agg("name") 
				FROM "User" u
				WHERE u."id" IN (
					SELECT uf."userId"
					FROM "UserForum" uf
					WHERE uf."forumId" = f."id"
				)
			) users
		FROM "Forum" f
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var res []*Forum
	for rows.Next() {
		var f Forum
		if err := rows.Scan(&f.Id, &f.Name, &f.TopicKeyword, pq.Array(&f.Users)); err != nil {
			return nil, err
		}
		res = append(res, &f)
	}
	if res == nil {
		res = make([]*Forum, 0)
	}
	return res, nil
}

func (r *Repository) AddUser(user User) {
	fmt.Printf("Add User")
	rows, err :=  r.Db.Query(`
		INSERT INTO "User" ("name") VALUES ($1)
	`, user.Name)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
}
