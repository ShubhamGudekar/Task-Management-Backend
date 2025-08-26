package migrations

import (
	"Task-Management-Backend/internal/infrastructure"
	"Task-Management-Backend/internal/model"
	"fmt"
)

func MigrateDB() {
	if err := infrastructure.DB.AutoMigrate(&model.User{}, &model.Task{}); err != nil {
		panic(fmt.Sprintf("Failed to migrate tables: %v", err))
	}
}
