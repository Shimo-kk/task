package usecase

import (
	"task/app/application/interface/database"
	"task/app/application/interface/usecase"
	"task/app/application/schema"
	"task/app/core"
	"task/app/domain/user"
	"task/app/domain/workspace"
)

type workspaceUsecase struct {
	databaseHandller database.IDatabaseHandller
}

// ワークスペースユースケースの作成
func NewWorkspaceUsecase(databaseHandller database.IDatabaseHandller) usecase.IWorkspaceUsecase {
	return &workspaceUsecase{databaseHandller: databaseHandller}
}

// ワークスペースの作成
func (u *workspaceUsecase) CreateWorkspace(data schema.WorkspaceCreateModel) error {
	return u.databaseHandller.Transaction(func(rf database.IRepositoryFactory) error {
		workspaceRepository := rf.GetWorkspaceRepository()
		userRepository := rf.GetUserRepository()

		// 同名のワークスペースが存在しないか確認
		notExists, err := workspaceRepository.NotExists(data.Name)
		if err != nil {
			return err
		}
		if !notExists {
			return core.NewError(core.BadRequestError, "同名のワークスペースが既に存在しています。")
		}

		// ワークスペースのエンティティを作成
		workspaceEntity, err := workspace.NewEntity(data.Name)
		if err != nil {
			return err
		}

		// ワークスペースを挿入
		insertedWorkspace, err := workspaceRepository.Insert(workspaceEntity)
		if err != nil {
			return err
		}

		// 管理者ユーザーのエンティティを作成
		userEntity, err := user.NewEntity(insertedWorkspace.GetId(), data.Name, data.UserEmail, data.UserPassword, true)
		if err != nil {
			return err
		}

		// ユーザーを挿入
		_, err = userRepository.Insert(userEntity)
		if err != nil {
			return err
		}

		return nil
	})
}
