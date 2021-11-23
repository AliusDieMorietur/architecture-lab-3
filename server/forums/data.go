package forums

import "fmt"

type Forum struct {
	TopicKeyword string
	Name         string
	Users        []string
}

type User struct {
	Name      string
	Interests []string
}

type Repository struct {
}

func (r *Repository) GetForumsList() ([]*Forum, error) {
	fmt.Printf("GetForumsList stub")

	// Мои мысли:
	//  Лезем в БД за списком форумов и людей, смотрим, какие у кого интересы, создаём список, готово.
	//  Псевдокод:

	// 	forums = map { keyword : Forum }
	//  заполнить forums с БД
	//	usersFromDb.foreach { user => user.interests.foreach { interest => forums[interest].add(user.name) } }
	// 	return forums.values

	// Другие варианты приветствуются.

	return []*Forum{
		{
			"stub_forum_1",
			"Stub Name 1",
			[]string{"stub_user_1", "stub_user_2"},
		},
		{
			"stub_forum_2",
			"Stub Name 2",
			[]string{"stub_user_1", "stub_user_3"},
		},
	}, nil
}

func (r *Repository) AddUser(user User) {
	// Засунуть юзера в бд, вроде не сложно.
}
