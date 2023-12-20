package entities

import "time"

// Task adalah struktur untuk merepresentasikan entitas tugas
type Task struct {
	ID          string     `json:"id"` // Tambahkan tag json agar dapat di-encode ke format JSON
	Title       string     `json:"title"`
	Owner_id    string     `json:"owner_id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	DueDate     time.Time  `json:"due_date"`
}

// TaskStatus adalah tipe enumerasi untuk status tugas
type TaskStatus string

const (
	NotDone   TaskStatus = "belum selesai"
	Done      TaskStatus = "selesai"
	Postponed TaskStatus = "ditunda"
)
