package psql

import (
	"database/sql"
	"github.com/dankru/practice-task1/internal/domain"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) GetAll() ([]domain.User, error) {
	rows, err := repo.db.Query("select * from users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]domain.User, 0)
	for rows.Next() {
		u := domain.User{}
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.RegisteredAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *Repository) GetById(id int64) (domain.User, error) {
	var u domain.User
	err := repo.db.QueryRow("select * from users WHERE id = $1", id).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.RegisteredAt)
	return u, err
}

func (repo *Repository) Create(user domain.User) error {
	_, err := repo.db.Exec("insert into users (name, email, password, registered_at) values ($1, $2, $3, $4)", user.Name, user.Email, user.Password, time.Now())
	return err
}

func (repo *Repository) Update(id int64, user domain.User) error {
	_, err := repo.db.Exec("update users set name = $1, email = $2, password = $3 WHERE id = $4", user.Name, user.Email, user.Password, id)
	return err
}

func (repo *Repository) Delete(id int64) error {
	_, err := repo.db.Exec("delete from users where id = $1", id)
	return err
}
