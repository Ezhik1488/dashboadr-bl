package clients

import (
	"dashboard-bl/internal/database/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strings"
)

type ClientRepo interface {
	GetByID(db *gorm.DB, id int) (*models.Client, error)
	GetWithFilters(filter ClientFilter) ([]*models.Client, int64, error)
	Create(db *gorm.DB, client *models.Client) error
	Update(db *gorm.DB, client *models.Client) error
	Delete(db *gorm.DB, id int) error
}

type clientRepo struct {
	db     *gorm.DB
	logger *logrus.Entry
}

// NewClientRepo - Конструктор ClientRepo
func NewClientRepo(db *gorm.DB, logger *logrus.Logger) ClientRepo {
	return &clientRepo{db: db, logger: logger.WithFields(logrus.Fields{"repo": "client"})}
}

type ClientFilter struct {
	Name        string
	Inn         string
	Email       string
	IsKik       *bool
	IsPaymaster *bool
	Page        int
	PageSize    int
	SortBy      string
	Include     []string
}

func (r *clientRepo) GetByID(db *gorm.DB, id int) (*models.Client, error) {
	var client models.Client
	if err := db.First(&client, id).Error; err != nil {
		r.logger.Warnf("error getting client by id=%d: %v", id, err)
		return nil, err
	}
	r.logger.Debugf("client found by id=%d", id)
	return &client, nil
}

func (r *clientRepo) GetWithFilters(filter ClientFilter) ([]*models.Client, int64, error) {
	var clients []*models.Client
	var total int64

	query := r.db.Model(&models.Client{}) // Базовый запрос

	// Фильтры
	if filter.Name != "" {
		query = query.Where("name ILIKE ?", "%"+filter.Name+"%")
	}
	if filter.Email != "" {
		query = query.Where("email ILIKE ?", "%"+filter.Email+"%")
	}
	if filter.Inn != "" {
		query = query.Where("inn ILIKE ?", "%"+filter.Inn+"%")
	}
	if filter.IsKik != nil {
		query = query.Where("isKik = ?", *filter.IsKik)
	}
	if filter.IsPaymaster != nil {
		query = query.Where("isPaymaster = ?", *filter.IsPaymaster)
	}

	// Подсчёт общего количества данных по запросу
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Пагинация
	page := filter.Page
	if page < 1 {
		page = 1
	}

	pageSize := filter.PageSize
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	offset := (page - 1) * pageSize

	// Сортировка
	var sortBy string
	allowedSortFields := map[string]bool{
		"name":         true,
		"inn":          true,
		"email":        true,
		"is_kik":       true,
		"is_paymaster": true,
		"created_at":   true,
		"update_at":    true,
	}

	// Разбиваем SortBy на поле и направление
	parts := strings.Fields(filter.SortBy)
	if len(parts) > 0 && allowedSortFields[parts[0]] {
		sortBy = filter.SortBy
	} else {
		sortBy = "created_at desc" // Значение по умолчанию
	}

	// Добавление Preload в зависимости от значения Include
	for _, include := range filter.Include {
		switch strings.ToLower(include) {
		case "details":
			query = query.Preload("Details")
		case "services":
			query = query.Preload("Services")
		default:
			// Игнорируем неизвестные значения
		}
	}

	// Поиск клиентов по указанным параметрам
	err = query.Order(sortBy).Offset(offset).Limit(pageSize).Find(&clients).Error

	return clients, total, nil
}

func (r *clientRepo) Create(db *gorm.DB, client *models.Client) error {
	return db.Create(client).Error
}

func (r *clientRepo) Update(db *gorm.DB, client *models.Client) error { return nil }

func (r *clientRepo) Delete(db *gorm.DB, id int) error {
	return db.Delete(&models.Client{}, id).Error
}
