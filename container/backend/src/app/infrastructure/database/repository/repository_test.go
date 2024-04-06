package repository_test

import (
	"os"
	"task/app/application/interface/database"
	databaseimpl "task/app/infrastructure/database"
	"task/migrations/migrate"
	"task/seeds/seed"

	"github.com/joho/godotenv"
)

func up() database.IDatabaseHandller {
	if err := godotenv.Load("../../../../.env"); err != nil {
		panic(err)
	}

	if err := migrate.UpMigration("../../../../migrations/sql", os.Getenv("DATABASE_URL_TEST")); err != nil {
		panic(err)
	}

	if err := seed.SeedTestData(os.Getenv("DATABASE_URL_TEST")); err != nil {
		panic(err)
	}

	databaseHandller := databaseimpl.NewDatabaseHandller()
	if err := databaseHandller.OpenDB(os.Getenv("DATABASE_URL_TEST")); err != nil {
		panic(err)
	}

	return databaseHandller
}

func down(databaseHandller database.IDatabaseHandller) {
	databaseHandller.CloseDB()
	migrate.DownMigration("../../../../migrations/sql", os.Getenv("DATABASE_URL_TEST"))
}
