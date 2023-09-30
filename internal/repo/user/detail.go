package user

import (
	"context"
	"database/sql"
	"time"

	"github.com/asyauqi1511/go-workshop-service/internal/entity"
	"github.com/asyauqi1511/go-workshop-service/internal/pkg/errorwrapper"
)

func (m Module) GetUserDetailByUserID(ctx context.Context, userId int64) (entity.UserDetail, error) {
	var result entity.UserDetail

	row := m.db.QueryRowContext(ctx, queryGetUserDetailByUserID, userId)

	err := row.Err()
	if row.Err() != nil && row.Err() != sql.ErrNoRows {
		return result, errorwrapper.Wrap(err)
	}

	err = row.Scan(&result.DetailUserID, &result.UserID, &result.Name, &result.Address, &result.SchoolName, &result.CreateTime, &result.UpdateTime)
	if err != nil {
		return result, errorwrapper.Wrap(err)
	}

	return result, nil
}

func (m Module) insertUserDetail(ctx context.Context, userId int64, data entity.UserRegisterParam) (int64, error) {

	result, err := m.db.ExecContext(ctx, queryInsertUserDetail, userId, data.Name, data.Address, data.SchoolName, time.Now(), time.Now())
	if err != nil {
		return 0, errorwrapper.Wrap(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, errorwrapper.Wrap(err)
	}

	return id, nil
}
