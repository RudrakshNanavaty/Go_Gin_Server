package database

import "web-server/entities"

type UserDB struct {
	users []*entities.User
}

func NewUserDB() *UserDB {
	return &UserDB{
		users: []*entities.User{
			{
				ID:   1,
				Name: "ABC",
			},
			{
				ID:   2,
				Name: "DEF",
			},
			{
				ID:   3,
				Name: "GHI",
			},
		},
	}
}

func (repo *UserDB) GetAll() ([]*entities.User, error) {
	return repo.users, nil
}

func (repo *UserDB) GetByID(id int) (*entities.User, error) {
	for _, user := range repo.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, nil // or an error indicating user not found
}

func (repo *UserDB) CreateNew(user *entities.User) error {
	repo.users = append(repo.users, user)
	return nil
}

func (repo *UserDB) DeleteByID(id int) error {
	for i, user := range repo.users {
		if user.ID == id {
			repo.users = append(repo.users[:i], repo.users[i+1:]...)
			return nil
		}
	}
	return nil // or an error indicating user not found
}
