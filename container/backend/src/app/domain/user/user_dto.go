package user

import "time"

type User struct {
	Id          int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	WorkspaceId int
	Name        string
	Email       string
	Password    string
	AdminFlag   bool
}

// エンティティをDTOに変換
func ToDtoFromEntity(e *UserEntity) *User {
	return &User{
		Id:          e.id,
		CreatedAt:   e.createdAt,
		UpdatedAt:   e.updatedAt,
		WorkspaceId: e.workspaceId,
		Name:        e.name,
		Email:       e.email,
		Password:    e.password,
		AdminFlag:   e.adminFlag,
	}
}

// DTOをエンティティに変換
func (d *User) ToEntity() *UserEntity {
	return &UserEntity{
		id:          d.Id,
		createdAt:   d.CreatedAt,
		updatedAt:   d.UpdatedAt,
		workspaceId: d.WorkspaceId,
		name:        d.Name,
		email:       d.Email,
		password:    d.Password,
		adminFlag:   d.AdminFlag,
	}
}
