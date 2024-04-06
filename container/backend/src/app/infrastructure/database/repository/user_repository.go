package repository

import (
	"task/app/core"
	"task/app/domain/user"

	"gorm.io/gorm"
)

type userRepository struct {
	tx *gorm.DB
}

func NewUserRepository(tx *gorm.DB) user.IUserRepository {
	return &userRepository{tx: tx}
}

// 挿入
func (r *userRepository) Insert(entity *user.UserEntity) (*user.UserEntity, error) {
	dto := user.ToDtoFromEntity(entity)
	if err := r.tx.Create(dto).Error; err != nil {
		return nil, core.NewError(core.SystemError, "ユーザーの作成に失敗しました。->"+err.Error())
	}
	result := dto.ToEntity()
	return result, nil
}

// 存在していないか確認
func (r *userRepository) NotExists(workspaceId int, email string) (bool, error) {
	dto := user.User{}
	err := r.tx.Where(&user.User{WorkspaceId: workspaceId, Email: email}).First(&dto).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return true, nil
		} else {
			return true, core.NewError(core.SystemError, "ワークスペースの取得に失敗しました。->"+err.Error())
		}
	}

	return false, nil
}

// IDで取得
func (r *userRepository) FindById(id int) (*user.UserEntity, error) {
	dto := user.User{}
	if err := r.tx.Where(&user.User{Id: id}).First(&dto).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			return nil, core.NewError(core.SystemError, "ユーザーの取得に失敗しました。->"+err.Error())
		}
	}
	result := dto.ToEntity()
	return result, nil
}

// Emailで取得
func (r *userRepository) FindByEmail(workspaceId int, email string) (*user.UserEntity, error) {
	dto := user.User{}
	if err := r.tx.Where(&user.User{WorkspaceId: workspaceId, Email: email}).First(&dto).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			return nil, core.NewError(core.SystemError, "ユーザーの取得に失敗しました。->"+err.Error())
		}
	}
	result := dto.ToEntity()
	return result, nil
}

// 更新
func (r *userRepository) Update(entity *user.UserEntity) (*user.UserEntity, error) {
	dto := user.ToDtoFromEntity(entity)
	if err := r.tx.Save(dto).Error; err != nil {
		return nil, core.NewError(core.SystemError, "ユーザーの更新に失敗しました。->"+err.Error())
	}
	result := dto.ToEntity()
	return result, nil
}

// 削除
func (r *userRepository) DeleteById(id int) error {
	if err := r.tx.Delete(&user.User{Id: id}).Error; err != nil {
		return core.NewError(core.SystemError, "ユーザーの削除に失敗しました。->"+err.Error())
	}
	return nil
}
