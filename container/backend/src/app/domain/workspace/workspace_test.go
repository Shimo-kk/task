package workspace_test

import (
	"task/app/core"
	"task/app/domain/workspace"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEntityOk(t *testing.T) {
	name := "test.workspace"
	entity1, err1 := workspace.NewEntity(name)
	if err1 != nil {
		t.Errorf(err1.Error())
	}
	assert.Equal(t, 0, entity1.GetId())
	assert.Equal(t, name, entity1.GetName())
}

func TestNewEntityNgNameValid(t *testing.T) {
	name := "test.test.test.test.test.test.test.test.test.test.workspace"
	entity, err := workspace.NewEntity(name)
	if entity != nil {
		t.Errorf("Error")
	}

	dstErr := core.AsAppError(err)
	assert.Equal(t, core.ValidationError, dstErr.Code())

	name = "テストワークスペース"
	entity, err = workspace.NewEntity(name)
	if entity != nil {
		t.Errorf("Error")
	}

	dstErr = core.AsAppError(err)
	assert.Equal(t, core.ValidationError, dstErr.Code())
}
