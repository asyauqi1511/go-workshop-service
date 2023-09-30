package user

import (
	"database/sql"

	"github.com/go-redis/redis/v8"
)

type Module struct {
	db    *sql.DB
	redis *redis.Client
}

func New(db *sql.DB, redis *redis.Client) Module {
	return Module{
		db:    db,
		redis: redis,
	}
}
