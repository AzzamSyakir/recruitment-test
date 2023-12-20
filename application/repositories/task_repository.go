package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"recruitment-test/application/entities"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (tr *TaskRepository) CreateTask(task entities.Task, taskNewID string) (string, error) {
	result, err := tr.db.Exec("INSERT INTO tasks (id, title, owner_id, description, status, due_date) VALUES (?, ?, ?, ?, ?)",
		taskNewID, task.Title, task.Description, task.Status, task.DueDate)
	if err != nil {
		return "", err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", err
	}

	if rowsAffected == 0 {
		return "", errors.New("failed to create task")
	}

	return taskNewID, nil
}

func (tr *TaskRepository) FetchTasks() ([]entities.Task, error) {
	rows, err := tr.db.Query("SELECT id, title, description, status, due_date FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []entities.Task

	for rows.Next() {
		var task entities.Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.DueDate,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
func (tr *TaskRepository) UpdateTask(id int, updatedTask entities.Task) error {
	// Validasi bisnis atau logika lainnya jika diperlukan
	// ...

	result, err := tr.db.Exec("UPDATE tasks SET title=?, description=?, status=?, due_date=? WHERE id=?",
		updatedTask.Title, updatedTask.Description, string(updatedTask.Status), updatedTask.DueDate, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("task with ID %d not found", id)
	}

	return nil
}

func (tr *TaskRepository) GetTask(id int) (*entities.Task, error) {
	task := &entities.Task{}

	err := tr.db.QueryRow("SELECT id, title, description, status, due_date FROM tasks WHERE id=?", id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.DueDate,
	)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (tr *TaskRepository) DeleteTask(id int) error {
	result, err := tr.db.Exec("DELETE FROM tasks WHERE id=?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("task with ID %d not found", id)
	}

	return nil
}
