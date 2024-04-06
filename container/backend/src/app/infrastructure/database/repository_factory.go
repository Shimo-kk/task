package database

import (
	"task/app/application/interface/database"
	"task/app/domain/user"
	"task/app/domain/workspace"
	"task/app/infrastructure/database/repository"

	"gorm.io/gorm"
)

type repositoryFactory struct {
	db *gorm.DB
}

func NewRepositoryFactory(db *gorm.DB) database.IRepositoryFactory {
	return &repositoryFactory{db: db}
}

func (rf *repositoryFactory) GetWorkspaceRepository() workspace.IWorkspaceRepository {
	return repository.NewWorkspaceRepository(rf.db)
}

func (rf *repositoryFactory) GetUserRepository() user.IUserRepository {
	return repository.NewUserRepository(rf.db)
}
