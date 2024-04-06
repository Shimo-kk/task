package database

import (
	"task/app/domain/user"
	"task/app/domain/workspace"
)

type IRepositoryFactory interface {
	GetWorkspaceRepository() workspace.IWorkspaceRepository
	GetUserRepository() user.IUserRepository
}
