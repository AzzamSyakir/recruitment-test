package config

import (
	"database/sql"
	"fmt"
	"os"
	"recruitment-test/application/migration"

	_ "github.com/go-sql-driver/mysql" // Import driver MySQL
	"github.com/rs/zerolog/log"
)

var DB *sql.DB

// InitDB digunakan untuk menghubungkan ke database.
func InitDB() *sql.DB {

	//baca env nya
	sqlInfo := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Membuka koneksi ke database
	db, err := sql.Open("mysql", sqlInfo)
	if err != nil {
		log.Fatal().Err(err).Msg("Gagal membuka koneksi database")
	}

	// Memeriksa koneksi ke database
	err = db.Ping()
	if err != nil {
		log.Fatal().Err(err).Msg("Gagal melakukan ping ke database")
	}

	log.Info().Msg("Terhubung ke database!")

	// Panggil fungsi migrate untuk inisialisasi migrasi database
	migration.UserMigrate(db)  // User
	migration.TokenMigrate(db) // tokens
	migration.TaskMigrate(db)  // tasks

	DB = db

	return db
}
