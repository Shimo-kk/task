package workspace

type IWorkspaceRepository interface {
	Insert(entity *WorkspaceEntity) (*WorkspaceEntity, error)
	NotExists(name string) (bool, error)
	FindById(id int) (*WorkspaceEntity, error)
	FindByName(name string) (*WorkspaceEntity, error)
	DeleteById(id int) error
}
