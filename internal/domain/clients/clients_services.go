package clients

import (
	"dashboard-bl/internal/database/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ClientServices interface {
	GetClients(filter ClientFilter) ([]*models.Client, int64, error)
}

type clientServices struct {
	db         *gorm.DB
	logger     *logrus.Entry
	clientRepo ClientRepo
}

func NewClientServices(db *gorm.DB, logger *logrus.Logger, clientRepo ClientRepo) ClientServices {
	return &clientServices{db: db, logger: logger.WithField("service", "clients"), clientRepo: clientRepo}
}

func (s *clientServices) GetClients(filter ClientFilter) ([]*models.Client, int64, error) {
	return s.clientRepo.GetWithFilters(filter)
}
