package data

import "time"

//Task struct
type Task struct {
	Id        string
	Uuid      string
	TaskName  string
	UserId    int
	CreatedAt time.Time
}

//Note struct
type Note struct {
	Id          string
	Uuid        string
	Description string
	TaskId      int
	CreatedAt   time.Time
}

// CreateTask method
func (user *User) CreateTask(taskName string) (conv Task, err error) {
	statement := "INSERT INTO tasks (uuid, task_name, created_at) VALUES ($1, $2, $3) RETURNING id, uuid, task_name, user_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), taskName, time.Now()).
		Scan(&conv.Id, &conv.Uuid, &conv.CreatedAt)
	return
}
