package seed

import (
	"fmt"
	"task/app/domain/user"
	"task/app/domain/workspace"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SeedTestData(url string) error {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("postgres://%s", url)), &gorm.Config{})
	if err != nil {
		return err
	}

	// workspaces
	if err := seedWorkspace(db); err != nil {
		return err
	}

	// users
	if err := seedUser(db); err != nil {
		return err
	}

	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		return err
	}

	return nil
}

func seedWorkspace(db *gorm.DB) error {
	testWorkspaces := []ModelInterface{
		&workspace.Workspace{
			Id:   1,
			Name: "test1.workspace",
		},
		&workspace.Workspace{
			Id:   2,
			Name: "test2.workspace",
		},
		&workspace.Workspace{
			Id:   3,
			Name: "test3.workspace",
		},
	}
	if err := seed(db, testWorkspaces); err != nil {
		return err
	}
	if err := db.Exec("SELECT setval('workspaces_id_seq', (SELECT MAX(id) FROM workspaces));").Error; err != nil {
		return err
	}

	return nil
}

func seedUser(db *gorm.DB) error {
	entity1, _ := user.NewEntity(1, "testuser1", "test1@example.com", "testtest", false)
	entity2, _ := user.NewEntity(1, "testuser2", "test2@example.com", "testtest", false)
	entity3, _ := user.NewEntity(1, "testuser3", "test3@example.com", "testtest", false)
	user1 := user.ToDtoFromEntity(entity1)
	user2 := user.ToDtoFromEntity(entity2)
	user3 := user.ToDtoFromEntity(entity3)
	user1.Id = 1
	user2.Id = 2
	user3.Id = 3
	testUsers := []ModelInterface{
		user1,
		user2,
		user3,
	}
	if err := seed(db, testUsers); err != nil {
		return err
	}
	if err := db.Exec("SELECT setval('users_id_seq', (SELECT MAX(id) FROM users));").Error; err != nil {
		return err
	}

	return nil
}
