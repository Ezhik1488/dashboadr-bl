package client_details

import (
	"dashboard-bl/internal/database/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ClientDetailsRepo interface {
	GetByID(db *gorm.DB, id int) (*models.ClientDetail, error)
	GetWithFilter(filter ClientDetailsFilter) error
	Create(item *models.ClientDetail) error
	Update(item *models.ClientDetail) error
	Delete(id int) error
}

type clientDetails struct {
	db     *gorm.DB
	logger *logrus.Entry
}

type ClientDetailsFilter struct {
}

func New(db *gorm.DB, logger *logrus.Logger) ClientDetailsRepo {
	return &clientDetails{
		db:     db,
		logger: logger.WithFields(logrus.Fields{"repo": "client_details"}),
	}
}

func (r *clientDetails) GetByID(db *gorm.DB, id int) (*models.ClientDetail, error) {
	var result models.ClientDetail
	if err := db.First(&result, id).Error; err != nil {
		r.logger.Warnf("cannot get client details by id: %v", err)
		return nil, err
	}
	return &result, nil
}

func (r *clientDetails) GetWithFilter(filter ClientDetailsFilter) error { return nil }

func (r *clientDetails) Create(item *models.ClientDetail) error {
	return r.db.Create(item).Error
}

func (r *clientDetails) Update(item *models.ClientDetail) error {
	return r.db.Model(&models.Client{}).Updates(item).Error
}

func (r *clientDetails) Delete(id int) error {
	return r.db.Delete(&models.ClientDetail{}, id).Error
}
