package datatype

import "github.com/qjouda/dignity-platform/backend/ethereum"

//ServiceContainer defines our service container type
type ServiceContainer struct {
	Config          Config
	AssetService    AssetService
	UserService     UserService
	FileStorage     FileStorage
	TimelineService TimelineService
	Ethereum        *ethereum.Ethereum
}
