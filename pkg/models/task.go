package models

type PostTask struct {
  TaskType string `json:"task_type"`
  ScheduleTime time.Time `json:"schedule_time"`
  Payload map[string]any `json:"payload"`
}
