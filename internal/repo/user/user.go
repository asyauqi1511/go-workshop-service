package user

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/asyauqi1511/go-workshop-service/internal/constants"
	"github.com/asyauqi1511/go-workshop-service/internal/entity"
	"github.com/asyauqi1511/go-workshop-service/internal/pkg/errorwrapper"
)

func (m Module) GetAuth(ctx context.Context, username, password string) (entity.User, error) {
	var result entity.User

	row := m.db.QueryRowContext(ctx, queryGetUserAuth, username, password)

	err := row.Err()
	if row.Err() != nil && row.Err() != sql.ErrNoRows {
		return result, errorwrapper.Wrap(err)
	}

	err = row.Scan(&result.UserID, &result.Username, &result.Role, &result.CreateTime, &result.UpdateTime)
	if err != nil && err != sql.ErrNoRows {
		return result, errorwrapper.Wrap(err)
	}

	return result, nil
}

func (m Module) GetUserByID(ctx context.Context, userId int64) (entity.User, error) {
	var result entity.User

	row := m.db.QueryRowContext(ctx, queryGetUserByID, userId)

	err := row.Err()
	if row.Err() != nil && row.Err() != sql.ErrNoRows {
		return result, errorwrapper.Wrap(err)
	}

	err = row.Scan(&result.UserID, &result.Username, &result.Role, &result.CreateTime, &result.UpdateTime)
	if err != nil {
		return result, errorwrapper.Wrap(err)
	}

	return result, nil
}

func (m Module) RegisterUser(ctx context.Context, data entity.UserRegisterParam) error {
	tx, err := m.db.Begin()
	if err != nil {
		return errorwrapper.Wrap(err)
	}

	defer func() {
		if err != nil {
			errTx := tx.Rollback()
			log.Print(errTx)
		} else {
			errTx := tx.Commit()
			log.Print(errTx)
		}
	}()

	userId, err := m.insertUser(ctx, data)
	if err != nil {
		return errorwrapper.Wrap(err)
	}

	_, err = m.insertUserDetail(ctx, userId, data)
	if err != nil {
		return errorwrapper.Wrap(err)
	}

	return nil
}

func (m Module) insertUser(ctx context.Context, data entity.UserRegisterParam) (int64, error) {
	result, err := m.db.ExecContext(ctx, queryInsertUser, data.Username, data.Password, constants.UserRoleRegistrant, time.Now(), time.Now())
	if err != nil {
		return 0, errorwrapper.Wrap(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, errorwrapper.Wrap(err)
	}

	return id, nil
}
