package models

import "time"

type TaskSubmission struct {
	ID          uint      `json:"submission_id" gorm:"primaryKey"`
	CreatedAt   time.Time `json:"created_time"`
	UserId      int32     `json:"user_id"`
	TaskCode    string    `json:"task_code"`
	TaskVersion int32     `json:"task_version"`
	LangCode    string    `json:"lang_code"`
}

type ExecSubmission struct {
	ID        uint      `json:"submission_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_time"`
	UserId    int32     `json:"user_id"`
	UserCode  string    `json:"user_code"`
	LangCode  string    `json:"lang_code"`
	StdInput  string    `json:"std_input"`
}
