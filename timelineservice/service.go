package timelineservice

import (
	"github.com/qjouda/dignity-platform/backend/dbservice"
)

// Service defines service type
type Service struct {
	*dbservice.DB
}

// NewService factory for Service
func NewService(db *dbservice.DB) *Service {
	return &Service{db}
}
