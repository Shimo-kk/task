package workspace

import "time"

type WorkspaceEntity struct {
	id        int
	createdAt time.Time
	updatedAt time.Time
	name      string
}

// エンティティの作成
func NewEntity(name string) (*WorkspaceEntity, error) {
	// バリデーション
	if err := validateName(name); err != nil {
		return nil, err
	}

	return &WorkspaceEntity{name: name}, nil
}

// ゲッター　ID
func (e *WorkspaceEntity) GetId() int {
	return e.id
}

// ゲッター　作成日時
func (e *WorkspaceEntity) GetCreatedAt() time.Time {
	return e.createdAt
}

// ゲッター　更新日時
func (e *WorkspaceEntity) GetUpdatedAt() time.Time {
	return e.updatedAt
}

// ゲッター　名称
func (e *WorkspaceEntity) GetName() string {
	return e.name
}
