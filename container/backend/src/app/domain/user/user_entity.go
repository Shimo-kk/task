package user

import (
	"task/app/core"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserEntity struct {
	id          int
	createdAt   time.Time
	updatedAt   time.Time
	workspaceId int
	name        string
	email       string
	password    string
	adminFlag   bool
}

// エンティティの作成
func NewEntity(workspaceId int, name string, email string, password string, adminFlg bool) (*UserEntity, error) {
	// バリデーション
	if err := validateWorkspaceId(workspaceId); err != nil {
		return nil, err
	}
	if err := validateName(name); err != nil {
		return nil, err
	}
	if err := validateEmail(email); err != nil {
		return nil, err
	}
	if err := validatePassword(password); err != nil {
		return nil, err
	}

	// パスワードの暗号化
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, core.NewError(core.SystemError, "パスワードの暗号化に失敗しました。->"+err.Error())
	}

	return &UserEntity{workspaceId: workspaceId, name: name, email: email, password: string(hashed), adminFlag: adminFlg}, nil
}

// パスワードの検証
func (e *UserEntity) VerifyPassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(e.password), []byte(password)); err != nil {
		return core.NewError(core.BadRequestError, "パスワードが正しくありません。")
	}

	return nil
}

// ユーザー名の変更
func (e *UserEntity) ChangeName(name string) error {
	// バリデーション
	if err := validateName(name); err != nil {
		return err
	}

	e.name = name
	return nil
}

// パスワードの変更
func (e *UserEntity) ChangePassword(password string) error {
	// バリデーション
	if err := validatePassword(password); err != nil {
		return err
	}

	// パスワードの暗号化
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return core.NewError(core.SystemError, "パスワードの暗号化に失敗しました。->"+err.Error())
	}

	e.password = string(hashed)
	return nil
}

// 管理者フラグの変更
func (e *UserEntity) ChangeAdminFlag(adminFlag bool) error {
	e.adminFlag = adminFlag
	return nil
}

// ゲッター　ID
func (e *UserEntity) GetId() int {
	return e.id
}

// ゲッター　作成日時
func (e *UserEntity) GetCreatedAt() time.Time {
	return e.createdAt
}

// ゲッター　更新日時
func (e *UserEntity) GetUpdatedAt() time.Time {
	return e.updatedAt
}

// ゲッター　ワークスペースID
func (e *UserEntity) GetWorkspaceId() int {
	return e.workspaceId
}

// ゲッター　名称
func (e *UserEntity) GetName() string {
	return e.name
}

// ゲッター　E-mailアドレス
func (e *UserEntity) GetEmail() string {
	return e.email
}

// ゲッター　パスワード
func (e *UserEntity) GetPassword() string {
	return e.password
}

// ゲッター　管理者フラグ
func (e *UserEntity) GetAdminFlag() bool {
	return e.adminFlag
}
