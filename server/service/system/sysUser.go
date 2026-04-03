package system

import (
	"errors"
	"server/global"
	"server/model/system"
	"server/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct{}

func (UserService *UserService) Register(u *system.SysUser) (err error) {
	// 判断用户是否注册
	if !errors.Is(global.BIGO_DB.Where("username = ?", u.Username).First(u).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户已注册")
	}
	// 生成 uuid，加密密码
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.New()
	err = global.BIGO_DB.Create(u).Error
	return err
}
