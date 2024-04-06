package repository_test

import (
	"task/app/application/interface/database"
	"task/app/domain/workspace"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkspaceInsertOk(t *testing.T) {
	databaseHandller := up()
	defer down(databaseHandller)

	name := "test.workspace"
	entity, err := workspace.NewEntity(name)
	if err != nil {
		t.Errorf(err.Error())
	}

	databaseHandller.Transaction(func(rf database.IRepositoryFactory) error {
		repository := rf.GetWorkspaceRepository()

		inserted, err := repository.Insert(entity)
		if err != nil {
			t.Errorf(err.Error())
			return err
		}

		assert.Equal(t, name, inserted.GetName())
		return nil
	})
}

func TestWorkspaceNotExistsOk(t *testing.T) {
	databaseHandller := up()
	defer down(databaseHandller)

	repositoryFactory := databaseHandller.GetRepositoryFactory()
	repository := repositoryFactory.GetWorkspaceRepository()

	exists1, _ := repository.NotExists("test1.workspace")
	exists2, _ := repository.NotExists("test.workspace")

	assert.Equal(t, false, exists1)
	assert.Equal(t, true, exists2)
}

func TestWorkspaceFindByIdOk(t *testing.T) {
	databaseHandller := up()
	defer down(databaseHandller)

	repositoryFactory := databaseHandller.GetRepositoryFactory()
	repository := repositoryFactory.GetWorkspaceRepository()

	entity1, _ := repository.FindById(1)
	entity2, _ := repository.FindById(2)
	entity3, _ := repository.FindById(3)
	entity4, _ := repository.FindById(4)

	if entity4 != nil {
		t.Errorf("Error")
	}

	assert.Equal(t, 1, entity1.GetId())
	assert.Equal(t, "test1.workspace", entity1.GetName())

	assert.Equal(t, 2, entity2.GetId())
	assert.Equal(t, "test2.workspace", entity2.GetName())

	assert.Equal(t, 3, entity3.GetId())
	assert.Equal(t, "test3.workspace", entity3.GetName())
}

func TestWorkspaceFindByNameOk(t *testing.T) {
	databaseHandller := up()
	defer down(databaseHandller)

	repositoryFactory := databaseHandller.GetRepositoryFactory()
	repository := repositoryFactory.GetWorkspaceRepository()

	entity1, _ := repository.FindByName("test1.workspace")
	entity2, _ := repository.FindByName("test2.workspace")
	entity3, _ := repository.FindByName("test3.workspace")
	entity4, _ := repository.FindByName("test4.workspace")

	if entity4 != nil {
		t.Errorf("Error")
	}

	assert.Equal(t, 1, entity1.GetId())
	assert.Equal(t, "test1.workspace", entity1.GetName())

	assert.Equal(t, 2, entity2.GetId())
	assert.Equal(t, "test2.workspace", entity2.GetName())

	assert.Equal(t, 3, entity3.GetId())
	assert.Equal(t, "test3.workspace", entity3.GetName())
}

func TestWorkspaceDeleteByIdOk(t *testing.T) {
	databaseHandller := up()
	defer down(databaseHandller)

	repositoryFactory := databaseHandller.GetRepositoryFactory()
	repository := repositoryFactory.GetWorkspaceRepository()

	repository.DeleteById(1)
	entity, _ := repository.FindById(1)

	if entity != nil {
		t.Errorf("Error")
	}
}
