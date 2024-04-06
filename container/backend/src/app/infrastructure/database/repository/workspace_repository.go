package repository

import (
	"task/app/core"
	"task/app/domain/workspace"

	"gorm.io/gorm"
)

type workspaceRepository struct {
	tx *gorm.DB
}

// ワークスペースリポジトリの作成
func NewWorkspaceRepository(tx *gorm.DB) workspace.IWorkspaceRepository {
	return &workspaceRepository{tx: tx}
}

// 挿入
func (r *workspaceRepository) Insert(entity *workspace.WorkspaceEntity) (*workspace.WorkspaceEntity, error) {
	dto := workspace.ToDtoFromEntity(entity)
	if err := r.tx.Create(dto).Error; err != nil {
		return nil, core.NewError(core.SystemError, "ワークスペースの作成に失敗しました。->"+err.Error())
	}
	result := dto.ToEntity()
	return result, nil
}

// 存在していないか確認
func (r *workspaceRepository) NotExists(name string) (bool, error) {
	dto := workspace.Workspace{}
	err := r.tx.Where(&workspace.Workspace{Name: name}).First(&dto).Error
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
func (r *workspaceRepository) FindById(id int) (*workspace.WorkspaceEntity, error) {
	dto := workspace.Workspace{}
	if err := r.tx.Where(&workspace.Workspace{Id: id}).First(&dto).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			return nil, core.NewError(core.SystemError, "ワークスペースの取得に失敗しました。->"+err.Error())
		}
	}
	result := dto.ToEntity()
	return result, nil
}

// 名称で取得
func (r *workspaceRepository) FindByName(name string) (*workspace.WorkspaceEntity, error) {
	dto := workspace.Workspace{}
	if err := r.tx.Where(&workspace.Workspace{Name: name}).First(&dto).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			return nil, core.NewError(core.SystemError, "ワークスペースの取得に失敗しました。->"+err.Error())
		}
	}
	result := dto.ToEntity()
	return result, nil
}

// 削除
func (r *workspaceRepository) DeleteById(id int) error {
	if err := r.tx.Delete(&workspace.Workspace{Id: id}).Error; err != nil {
		return core.NewError(core.SystemError, "ワークスペースの削除に失敗しました。->"+err.Error())
	}
	return nil
}
