package data

import (
	"database/sql"
	"fmt"
	"log"
)

type Forum struct {
	Id int64
	TopicKeyword string
	Name         string
	Users        []string
}

type User struct {
	Id int64
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
				SELECT "name" 
				FROM "User" u
				WHERE u."id" IN (
					SELECT * 
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
		if err := rows.Scan(&f.Id, &f.Name); err != nil {
			return nil, err
		}
		res = append(res, &f)
	}
	if res == nil {
		res = make([]*Forum, 0)
	}
	return res, nil
	// Мои мысли:
	//  Лезем в БД за списком форумов и людей, смотрим, какие у кого интересы, создаём список, готово.
	//  Псевдокод:

	// 	forums = map { keyword : Forum }
	//  заполнить forums с БД
	//	usersFromDb.foreach { user => user.interests.foreach { interest => forums[interest].add(user.name) } }
	// 	return forums.values

	// Другие варианты приветствуются.

	// return []*Forum{
	// 	{
	// 		"stub_forum_1",
	// 		"Stub Name 1",
	// 		[]string{"stub_user_1", "stub_user_2"},
	// 	},
	// 	{
	// 		"stub_forum_2",
	// 		"Stub Name 2",
	// 		[]string{"stub_user_1", "stub_user_3"},
	// 	},
	// }, nil
}

func (r *Repository) AddUser(user User) {
	// Засунуть юзера в бд, вроде не сложно.
}
