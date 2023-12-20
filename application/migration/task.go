package migration

import (
	"database/sql"
	"log"
)

// TaskMigrate digunakan untuk menjalankan migrasi tabel task.
func TaskMigrate(db *sql.DB) {
	// SQL statement untuk memeriksa apakah tabel task sudah ada
	checkTableSQL := `
        SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = DATABASE() AND table_name = 'task'
    `

	// Menjalankan perintah SQL untuk memeriksa apakah tabel sudah ada
	var tableCount int
	err := db.QueryRow(checkTableSQL).Scan(&tableCount)
	if err != nil {
		// Menangani kesalahan jika terjadi kesalahan saat memeriksa tabel
		log.Fatal(err)
		return
	}

	if tableCount > 0 {
		// Jika tabel sudah ada, tampilkan pesan
		log.Println("Tabel task sudah di migrasi")
		return
	}

	// SQL statement untuk membuat tabel task
	createTableSQL := `
        CREATE TABLE IF NOT EXISTS task (
            id CHAR(36) NOT NULL PRIMARY KEY,
            title VARCHAR(255) NOT NULL,
		owner_id CHAR(36),
            description VARCHAR(255) NOT NULL UNIQUE,
            status ENUM('belum selesai', 'selesai', 'ditunda') NOT NULL,
		FOREIGN KEY (owner_id) REFERENCES user(id)
            due_date TIMESTAMP NOT NULL
        )
    `

	// Menjalankan perintah SQL untuk membuat tabel
	_, err = db.Exec(createTableSQL)
	if err != nil {
		// Menangani kesalahan jika terjadi kesalahan saat migrasi
		log.Fatal(err)
		return
	}

	// Pesan sukses jika migrasi berhasil
	log.Println("Migrasi tabel task berhasil")
}
