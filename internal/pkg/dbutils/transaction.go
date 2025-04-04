package dbutils

import (
	"gorm.io/gorm"
	"log"
)

// WithTransaction оборачивает произвольную функцию в транзакцию.
// При ошибке внутри fn происходит rollback, иначе — commit.
func WithTransaction(db *gorm.DB, fn func(tx *gorm.DB) error) error {
	tx := db.Begin()
	if tx.Error != nil {
		log.Println("Ошибка при начале транзакции:", tx.Error)
		return tx.Error
	}

	if err := fn(tx); err != nil {
		if rbErr := tx.Rollback().Error; rbErr != nil {
			log.Println("Ошибка при rollback:", rbErr)
		}
		return err
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Ошибка при commit:", err)
		return err
	}

	return nil
}
