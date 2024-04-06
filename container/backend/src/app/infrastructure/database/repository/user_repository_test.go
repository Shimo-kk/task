package repository_test

import (
	"task/app/application/interface/database"
	"task/app/domain/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserInsertOk(t *testing.T) {
	databaseHandller := up()
	defer down(databaseHandller)

	name := "testuser"
	email := "test@example.com"
	password := "testtest"
	adminFlag := false
	entity, err := user.NewEntity(1, name, email, password, adminFlag)
	if err != nil {
		t.Errorf(err.Error())
	}

	databaseHandller.Transaction(func(rf database.IRepositoryFactory) error {
		repository := rf.GetUserRepository()

		inserted, err := repository.Insert(entity)
		if err != nil {
			t.Errorf(err.Error())
			return err
		}

		assert.Equal(t, name, inserted.GetName())
		assert.Equal(t, email, inserted.GetEmail())
		assert.Equal(t, true, inserted.VerifyPassword(password) == nil)
		assert.Equal(t, adminFlag, inserted.GetAdminFlag())
		return nil
	})
}

func TestUserNotExistsOk(t *testing.T) {
	databaseHandller := up()
	defer down(databaseHandller)

	repositoryFactory := databaseHandller.GetRepositoryFactory()
	repository := repositoryFactory.GetUserRepository()

	exists1, _ := repository.NotExists(1, "test1@example.com")
	exists2, _ := repository.NotExists(2, "test1@example.com")
	exists3, _ := repository.NotExists(1, "test@example.com")

	assert.Equal(t, false, exists1)
	assert.Equal(t, true, exists2)
	assert.Equal(t, true, exists3)
}

func TestUserFindByIdOk(t *testing.T) {
	databaseHandller := up()
	defer down(databaseHandller)

	repositoryFactory := databaseHandller.GetRepositoryFactory()
	repository := repositoryFactory.GetUserRepository()

	entity1, _ := repository.FindById(1)
	entity2, _ := repository.FindById(2)
	entity3, _ := repository.FindById(3)
	entity4, _ := repository.FindById(4)

	if entity4 != nil {
		t.Errorf("Error")
	}

	assert.Equal(t, 1, entity1.GetId())
	assert.Equal(t, "testuser1", entity1.GetName())

	assert.Equal(t, 2, entity2.GetId())
	assert.Equal(t, "testuser2", entity2.GetName())

	assert.Equal(t, 3, entity3.GetId())
	assert.Equal(t, "testuser3", entity3.GetName())
}

func TestWorkspaceFindByEmailOk(t *testing.T) {
	databaseHandller := up()
	defer down(databaseHandller)

	repositoryFactory := databaseHandller.GetRepositoryFactory()
	repository := repositoryFactory.GetUserRepository()

	entity1, _ := repository.FindByEmail(1, "test1@example.com")
	entity2, _ := repository.FindByEmail(1, "test2@example.com")
	entity3, _ := repository.FindByEmail(1, "test3@example.com")
	entity4, _ := repository.FindByEmail(1, "test4@example.com")

	if entity4 != nil {
		t.Errorf("Error")
	}

	assert.Equal(t, 1, entity1.GetId())
	assert.Equal(t, "testuser1", entity1.GetName())

	assert.Equal(t, 2, entity2.GetId())
	assert.Equal(t, "testuser2", entity2.GetName())

	assert.Equal(t, 3, entity3.GetId())
	assert.Equal(t, "testuser3", entity3.GetName())
}

func TestUserUpdate(t *testing.T) {
	databaseHandller := up()
	defer down(databaseHandller)

	name := "testuser"
	adminFlag := true

	databaseHandller.Transaction(func(rf database.IRepositoryFactory) error {
		repository := rf.GetUserRepository()

		entity, _ := repository.FindById(1)
		entity.ChangeName(name)
		entity.ChangeAdminFlag(adminFlag)

		updated, err := repository.Update(entity)
		if err != nil {
			t.Errorf(err.Error())
			return err
		}

		assert.Equal(t, name, updated.GetName())
		assert.Equal(t, adminFlag, updated.GetAdminFlag())
		return nil
	})
}

func TestUserDeleteByIdOk(t *testing.T) {
	databaseHandller := up()
	defer down(databaseHandller)

	repositoryFactory := databaseHandller.GetRepositoryFactory()
	repository := repositoryFactory.GetUserRepository()

	repository.DeleteById(1)
	entity, _ := repository.FindById(1)

	if entity != nil {
		t.Errorf("Error")
	}
}
