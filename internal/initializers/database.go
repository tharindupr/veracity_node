package initializers

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Define a model (table schema)
type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string
}

func ConnectToDB() {

	// Connect to SQLite database
	// dsn := os.GetEnv("DB_URL")
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	// Create a new record
	db.Create(&User{Name: "John Doe", Email: "john@example.com"})

}
