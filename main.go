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

// Application entry point
func main() {
	environment := strings.ToLower(env.MustGetenv("ENVIRONMENT"))
	// Initialize applicatoin configurations
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
	}
	logger.Setup(cfg)
	emailer.Setup(cfg)
	mysql := dbservice.MustConnect(cfg.DataSourceName)
	// Initializing application services user, asset,
	// timeline, and storage serivices
	userService := userservice.NewService(mysql)
	assetService := assetservice.NewService(mysql)
	timelineService := timelineservice.NewService(mysql)
	fs := filestorage.GetStorage(
		cfg.AwsSesKey,
		cfg.AwsSesSecret,
	)
	// connect with an instance of ethereum
	ethereum := ethereum.MustNewEthereum(cfg.EthereumHost)
	// Initilize service container
	sc := datatype.ServiceContainer{
		Config:          cfg,
		UserService:     userService,
		AssetService:    assetService,
		TimelineService: timelineService,
		Ethereum:        ethereum,
		FileStorage:     fs,
	}
	// Sertup API routing handler
	routes.SetupRouting(sc).Run(":" + cfg.ServerPort)
}

func getPort() string {
	port := env.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return port
}
