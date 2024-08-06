package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/mergestat/timediff"
)

type Todo struct {
	Id      int
	Title   string
	Created time.Time
	Done    bool
}

func (t Todo) String() string {
	if t.Done {
		return fmt.Sprintf("%d\t%s\t%s\t✅\n", t.Id, t.Title, timediff.TimeDiff(t.Created))
	} else {
		return fmt.Sprintf("%d\t%s\t%s\t❌\n", t.Id, t.Title, timediff.TimeDiff(t.Created))
	}
}

func init() {
	db := getConnection()
	defer db.Close()
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS todo (
		id integer not null primary key, 
		title text not null, 
		created timestamp not null, 
		done boolean not null
	);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func CreateTodo(title string) {
	db := getConnection()
	defer db.Close()

	_, err := db.Exec("INSERT INTO todo (title, created, done) VALUES (?, ?, ?)", title, time.Now().Format(time.RFC3339), false)
	if err != nil {
		log.Fatal(err)
	}
}

func GetTodos() []Todo {
	db := getConnection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM todo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	todos := []Todo{}
	for rows.Next() {
		var id int
		var title string
		var created time.Time
		var done bool
		err = rows.Scan(&id, &title, &created, &done)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, Todo{Id: id, Title: title, Created: created, Done: done})
	}

	return todos
}

func CompleteTodo(idToComplete string) {
	db := getConnection()
	defer db.Close()

	_, err := db.Exec("UPDATE todo SET done = ? WHERE id = ?", true, idToComplete)
	if err != nil {
		log.Fatal(err)
	}
}

func RemoveTodo(idToRemove string) bool {
	db := getConnection()
	defer db.Close()

	res, err := db.Exec("DELETE FROM todo WHERE id = ?", idToRemove)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	return rowsAffected > 0
}

func getConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "./todo.db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
