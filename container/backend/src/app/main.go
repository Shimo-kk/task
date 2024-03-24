package main

import (
	"os"
	"path/filepath"
	"task/app/infrastructure/database"
	"task/app/presentation/router"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

func main() {
	e := echo.New()

	// 環境変数の読み込み
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// ロガー設定の読み込み
	configYaml, err := os.ReadFile("./config/logger.yaml")
	if err != nil {
		panic(err)
	}
	var loggerConfig zap.Config
	if err := yaml.Unmarshal(configYaml, &loggerConfig); err != nil {
		panic(err)
	}

	// ログディレクトリの作成
	logDir := os.Getenv("LOG_OUTPUT_DIR")
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.MkdirAll(logDir, 0755); err != nil {
			panic(err)
		}
	}

	// ロガーの生成
	loggerConfig.OutputPaths = append(loggerConfig.OutputPaths, filepath.Join(logDir, "app.out.log"))
	loggerConfig.ErrorOutputPaths = append(loggerConfig.ErrorOutputPaths, filepath.Join(logDir, "app.err.log"))
	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}

	// データベースの設定
	databaseHandller := database.NewDatabaseHandller()
	if err := databaseHandller.OpenDB(os.Getenv("DATABASE_URL")); err != nil {
		panic(err)
	}
	defer databaseHandller.CloseDB()

	// ルーターの設定
	if err := router.SetUp(e, logger, databaseHandller); err != nil {
		panic(err)
	}

	// サーバーの起動
	port := os.Getenv("MY_PORT")
	e.Logger.Fatal(e.Start(":" + port))
}
