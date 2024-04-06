package usecase

import "task/app/application/schema"

type IWorkspaceUsecase interface {
	CreateWorkspace(data schema.WorkspaceCreateModel) error
}
