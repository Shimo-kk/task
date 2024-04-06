package workspace

import "time"

type Workspace struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

// エンティティをDTOに変換
func ToDtoFromEntity(e *WorkspaceEntity) *Workspace {
	return &Workspace{
		Id:        e.id,
		CreatedAt: e.createdAt,
		UpdatedAt: e.updatedAt,
		Name:      e.name,
	}
}

// DTOをエンティティに変換
func (d *Workspace) ToEntity() *WorkspaceEntity {
	return &WorkspaceEntity{
		id:        d.Id,
		createdAt: d.CreatedAt,
		updatedAt: d.UpdatedAt,
		name:      d.Name,
	}
}
