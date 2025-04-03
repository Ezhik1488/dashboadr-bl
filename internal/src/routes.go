package src

import (
	"dashboard-bl/internal/database"
	"fmt"
	"github.com/sirupsen/logrus"
)

func SetupRoutes(logger *logrus.Logger) {
	db := database.GetConnection()
	fmt.Println(db)

}
