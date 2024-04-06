package user

type IUserRepository interface {
	Insert(entity *UserEntity) (*UserEntity, error)
	NotExists(workspaceId int, email string) (bool, error)
	FindById(id int) (*UserEntity, error)
	FindByEmail(workspaceId int, email string) (*UserEntity, error)
	Update(entity *UserEntity) (*UserEntity, error)
	DeleteById(id int) error
}
