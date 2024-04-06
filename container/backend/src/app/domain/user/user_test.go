package user_test

import (
	"task/app/core"
	"task/app/domain/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEntity(t *testing.T) {
	workspaceId := 1
	name := "test user"
	email := "test@example.com"
	password := "testtest"
	adminFlag := true

	entity, err := user.NewEntity(workspaceId, name, email, password, adminFlag)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, 0, entity.GetId())
	assert.Equal(t, workspaceId, entity.GetWorkspaceId())
	assert.Equal(t, name, entity.GetName())
	assert.Equal(t, email, entity.GetEmail())
	assert.Equal(t, true, entity.VerifyPassword(password) == nil)
	assert.Equal(t, adminFlag, entity.GetAdminFlag())
}

func TestNewEntityNgWorkpaceIdValid(t *testing.T) {
	workspaceId := 0
	name := "test user"
	email := "test@example.com"
	password := "testtest"
	adminFlag := true

	entity, err := user.NewEntity(workspaceId, name, email, password, adminFlag)
	if entity != nil {
		t.Errorf("Error")
	}

	dstErr := core.AsAppError(err)
	assert.Equal(t, core.ValidationError, dstErr.Code())

	workspaceId = 1
	name = "testusertestusertestusertestusertestusertestusertestusertestusertestusertestusertestusertestusertestusertestusertestusertestuser"
	entity, err = user.NewEntity(workspaceId, name, email, password, adminFlag)
	if entity != nil {
		t.Errorf("Error")
	}

	dstErr = core.AsAppError(err)
	assert.Equal(t, core.ValidationError, dstErr.Code())

	name = "test user"
	email = "testmail"
	entity, err = user.NewEntity(workspaceId, name, email, password, adminFlag)
	if entity != nil {
		t.Errorf("Error")
	}

	dstErr = core.AsAppError(err)
	assert.Equal(t, core.ValidationError, dstErr.Code())

	email = "test@example.com"
	password = "testtestああ"
	entity, err = user.NewEntity(workspaceId, name, email, password, adminFlag)
	if entity != nil {
		t.Errorf("Error")
	}

	dstErr = core.AsAppError(err)
	assert.Equal(t, core.ValidationError, dstErr.Code())
}
