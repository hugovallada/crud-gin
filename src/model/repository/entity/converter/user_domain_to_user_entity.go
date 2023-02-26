package converter

import (
	"github.com/hugovallada/crud-gin/src/model"
	"github.com/hugovallada/crud-gin/src/model/repository/entity"
)

func FromUserDomainToUserEntity(user model.UserDomainInterface) *entity.UserEntity {
	return & entity.UserEntity{
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
		Name:     user.GetName(),
		Age:      user.GetAge(),
	}
}
