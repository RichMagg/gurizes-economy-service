package repositories

import (
	"context"

	"github.com/RichMagg/gurizes-economy-service/cmd/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return UserRepository{db: db}
}

func (r *UserRepository) GetUsers() ([]models.User, error) {
	rows, err := r.db.Query(context.Background(),
		`SELECT id, created_at, last_seen FROM users`,
	)

	if err != nil {
		return []models.User{}, err
	}

	defer rows.Close()

	users := []models.User{}

	for rows.Next() {
		var user models.User

		err := rows.Scan(
			&user.ID,
			&user.CreatedAt,
			&user.LastSeen,
		)

		if err != nil {
			return []models.User{}, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) CreateUser(user *models.User) error {
	_, err := r.db.Exec(context.Background(),
		`INSERT INTO users (id, created_at, last_seen)
		 VALUES ($1, $2, $3)`,
		user.ID,
		user.CreatedAt,
		user.LastSeen,
	)

	return err
}
