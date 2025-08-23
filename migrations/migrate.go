package migrations

import "Task-Management-Backend/internal/infrastructure"

func MigrateDB() {
	infrastructure.DB.AutoMigrate()
}
