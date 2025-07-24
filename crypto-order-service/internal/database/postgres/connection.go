package postgres

import (
	"log"

	"github.com/Amir-Sadati/crypto-microservice/crypto-order-service/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func Connect(cfg *config.PostgresConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", cfg.Url)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return db, nil
}
