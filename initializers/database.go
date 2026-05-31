package initializers

import (
	"log"

	"github.com/glebarez/sqlite" // pure-Go SQLite driver (works with CGO_ENABLED=0)
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB opens an embedded SQLite database stored as a file inside the
// container. No network connection is used, so it works even when the platform
// blocks outbound database traffic. The file is ephemeral (reset on redeploy),
// which is fine for a demo — the schema is recreated on startup via AutoMigrate.
func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("crud.db"), &gorm.Config{})
	if err != nil {
		// Don't crash the app on a DB error; the server still starts and /healthz works.
		log.Println("Failed to connect database:", err)
		return
	}
	DB = db
}
