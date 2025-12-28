package repository

import (
  "context"
  "time"
  "github.com/jmoiron/sqlx"
)

type Repository struct {
  db sqlx.DB
}

func NewRepository(db sqlx.DB) *Repository {
  return &Repository{db}
}

func (r *Repository) CreateTask(
  ctx context.Context,
  taskType string,
  scheduleTime time.Time,
  payload string,
) error {
  
  _, err := r.db.ExecContext(ctx, `
    INSERT INTO tasks (task_type, payload, schedule_for)
    VALUES (?, ?, ?)
  `, taskType, payload, scheduleTime)
  
  return err
  
}
