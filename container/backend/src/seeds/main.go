package main

import (
	"os"
	"task/seeds/seed"

	"github.com/joho/godotenv"
)

func main() {
	// 環境変数の読み込み
	if err := godotenv.Load(""); err != nil {
		panic(err)
	}

	// シードの実行
	if err := seed.SeedData(os.Getenv("DATABASE_URL")); err != nil {
		panic(err)
	}
}
