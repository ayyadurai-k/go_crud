package initializers

import (
	"log"
	"os"

	"github.com/glebarez/sqlite" // pure-Go SQLite driver (works with CGO_ENABLED=0)
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB opens an embedded SQLite database stored as a file inside the
// container. No network connection is used, so it works even when the platform
// blocks outbound database traffic. The file is ephemeral (reset on redeploy),
// which is fine for a demo — the schema is recreated on startup via AutoMigrate.
func ConnectDB() {
	// Store the SQLite file in /tmp, which is writable even when the container's
	// working directory / root filesystem is read-only (common on PaaS). Override
	// with DB_PATH if a persistent/volume path is available.
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "/tmp/crud.db"
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		// Don't crash the app on a DB error; the server still starts and /healthz works.
		log.Println("Failed to connect database:", err)
		return
	}
	DB = db
}
