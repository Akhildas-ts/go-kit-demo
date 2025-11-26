package task

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	spec "github.com/yourusername/task-service/spec/task"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateTask(ctx context.Context, title, description string) (uuid.UUID, error) {
	id := uuid.New()
	query := `INSERT INTO tasks (id, title, description) VALUES ($1, $2, $3)`
	_, err := r.db.ExecContext(ctx, query, id, title, description)
	return id, err
}

func (r *Repository) ListTasks(ctx context.Context) ([]spec.Task, error) {
	query := `SELECT id, title, description, completed, created_at FROM tasks ORDER BY created_at DESC`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []spec.Task
	for rows.Next() {
		var t spec.Task
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *Repository) GetTaskByID(ctx context.Context, id uuid.UUID) (*spec.Task, error) {
	query := `SELECT id, title, description, completed, created_at FROM tasks WHERE id = $1`
	var t spec.Task
	err := r.db.QueryRowContext(ctx, query, id).Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *Repository) CompleteTask(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE tasks SET completed = TRUE WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
