package main

import (
	"strings"

	"github.com/qjouda/dignity-platform/backend/assetservice"
	"github.com/qjouda/dignity-platform/backend/datatype"
	"github.com/qjouda/dignity-platform/backend/dbservice"
	"github.com/qjouda/dignity-platform/backend/emailer"
	"github.com/qjouda/dignity-platform/backend/env"
	"github.com/qjouda/dignity-platform/backend/ethereum"
	"github.com/qjouda/dignity-platform/backend/filestorage"
	"github.com/qjouda/dignity-platform/backend/logger"
	"github.com/qjouda/dignity-platform/backend/routes"
	"github.com/qjouda/dignity-platform/backend/timelineservice"
	"github.com/qjouda/dignity-platform/backend/userservice"
)

func main() {
	environment := strings.ToLower(env.MustGetenv("ENVIRONMENT"))

	cfg := datatype.Config{
		Environment:     environment,
		DBMigrationPath: env.MustGetenv("DB_MIGRATIONS_PATH"),
		ServerPort:      getPort(),
		DataSourceName:  env.MustGetenv("DATA_SOURCE_NAME"),
		AwsSesKey:       env.Getenv("AWS_SES_KEY"),
		AwsSesSecret:    env.Getenv("AWS_SES_SECRET"),
		AwsEmailFrom:    env.Getenv("EMAIL_FROM"),
		SlackWebhookURL: env.Getenv("SLACK_WEBHOOK_URL"),
		EthereumHost:    env.Getenv("RINKEBY_ETH_HOST"),
		MongoDBHost:     env.Getenv("MongoHost"),
		MongoDBName:     env.Getenv("MongoDBName"),
	}

	logger.Setup(cfg)

	emailer.Setup(cfg)

	mysql := dbservice.MustConnect(cfg.DataSourceName)
	// mongoDB := dbservice.NewMongoDB(cfg)
	userService := userservice.NewService(mysql)
	assetService := assetservice.NewService(mysql)
	timelineService := timelineservice.NewService(mysql)
	fs := filestorage.GetStorage(
		env.Getenv("AWS_SES_KEY"),
		env.Getenv("AWS_SES_SECRET"),
	)
	// connect with an instance of ethereum
	ethereum := ethereum.MustNewEthereum(cfg.EthereumHost)
	sc := datatype.ServiceContainer{
		Config:          cfg,
		UserService:     userService,
		AssetService:    assetService,
		TimelineService: timelineService,
		Ethereum:        ethereum,
		FileStorage:     fs,
	}

	routes.SetupRouting(sc).Run(":" + cfg.ServerPort)
}

func getPort() string {
	port := env.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return port
}
