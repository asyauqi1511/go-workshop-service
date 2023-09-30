package user

import (
	"context"
	"fmt"

	"github.com/asyauqi1511/go-workshop-service/internal/constants"
	"github.com/asyauqi1511/go-workshop-service/internal/entity"
	"github.com/asyauqi1511/go-workshop-service/internal/pkg/errorwrapper"
)

func (uc Usecase) GetRegistrantData(ctx context.Context, registrantId int64) (entity.UserDetailResult, error) {

	var result entity.UserDetailResult

	// Validation data
	value := ctx.Value(constants.UserIDTag)
	if value == nil {
		return result, errorwrapper.Wrap(fmt.Errorf("user not valid"))
	}

	userId, ok := value.(int64)
	if !ok {
		return result, errorwrapper.Wrap(fmt.Errorf("user not valid"))
	}

	value = ctx.Value(constants.UserRoleTag)
	if value == nil {
		return result, errorwrapper.Wrap(fmt.Errorf("user not valid"))
	}

	role, ok := value.(int)
	if !ok {
		return result, errorwrapper.Wrap(fmt.Errorf("user not valid"))
	}

	if userId != registrantId && role != constants.UserRoleAdmin {
		return result, errorwrapper.Wrap(fmt.Errorf("not authorized"))
	}

	// Get data.
	user, err := uc.userModule.GetUserByID(ctx, registrantId)
	if err != nil {
		return result, errorwrapper.Wrap(err)
	}
	if user.UserID == 0 || user.Role == constants.UserRoleAdmin {
		return result, errorwrapper.Wrap(fmt.Errorf("registrant id invalid"))
	}

	userDetail, err := uc.userModule.GetUserDetailByUserID(ctx, registrantId)
	if err != nil {
		return result, errorwrapper.Wrap(err)
	}
	if userDetail.DetailUserID == 0 {
		return result, errorwrapper.Wrap(fmt.Errorf("registrant id invalid"))
	}

	return entity.UserDetailResult{
		UserID:     registrantId,
		Username:   user.Username,
		Name:       userDetail.Name,
		Address:    userDetail.Address,
		SchoolName: userDetail.SchoolName,
	}, nil
}
