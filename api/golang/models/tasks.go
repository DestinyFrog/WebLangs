package tasks

import (
	"fmt"

	"app/db"
)

type Tasks struct {
	Id			int32		`json:"id"`
	Title		string		`json:"title"`
	Completed	bool		`json:"completed"`
}

type TasksRequest struct {
	Title		*string		`json:"title"`
	Completed	*bool		`json:"completed"`
}

func (o *Tasks) ToString() string {
	return fmt.Sprintf("%d. %s : %t", o.Id, o.Title, o.Completed)
}

func GetAll() (data []Tasks, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT id, title, completed FROM tasks;")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var obj Tasks

		err = rows.Scan(&obj.Id, &obj.Title, &obj.Completed)
		if err != nil {
			return
		}

		data = append(data, obj)
	}

	return
}

func GetOneById(id int) (data Tasks, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	err = conn.QueryRow("SELECT * FROM tasks WHERE id = $1;", id).Scan(&data.Id, &data.Title, &data.Completed)
	if err != nil {
		return
	}

	return
}

func CreateOne(data TasksRequest) (err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	_, err = conn.Exec("INSERT INTO tasks(title, completed) VALUES ($1, $2);", data.Title, data.Completed)
	if err != nil {
		return
	}

	return
}

func (o *Tasks) SetTitle(title string) (err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	_, err = conn.Exec("UPDATE tasks SET title=$2 WHERE id=$1;", o.Id, title)
	if err != nil {
		return
	}

	o.Title = title
	return
}

func (o *Tasks) SetCompleted(completed bool) (err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	_, err = conn.Exec("UPDATE tasks SET completed=$1 WHERE id=$2;", completed, o.Id)
	if err != nil {
		return
	}

	o.Completed = completed
	return
}

func (o *Tasks) Delete() (err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	_, err = conn.Exec("DELETE FROM tasks WHERE id=$1;", o.Id)
	if err != nil {
		return
	}

	return
}