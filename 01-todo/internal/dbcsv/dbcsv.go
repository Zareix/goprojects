package dbcsv

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

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
	_, err := os.Stat("todo.csv")
	if os.IsNotExist(err) {
		file, err := os.Create("todo.csv")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		headers := []string{"Id", "Title", "Done", "Created"}
		err = writer.Write(headers)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func CreateTodo(title string) {
	todo := Todo{Id: getNewTodoId(), Title: title, Created: time.Now(), Done: false}

	saveTodo(todo)
}

func GetTodos() []Todo {
	file, err := os.OpenFile("todo.csv", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	todos := []Todo{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if record[0] == "Id" {
			continue
		}

		id, _ := strconv.Atoi(record[0])
		created, _ := time.Parse(time.RFC3339, record[3])
		todo := Todo{Id: id, Title: record[1], Done: record[2] == "true", Created: created}
		todos = append(todos, todo)
	}

	return todos
}

func CompleteTodo(idToComplete string) {
	todos := GetTodos()
	for _, todo := range todos {
		if strconv.Itoa(todo.Id) == idToComplete {
			RemoveTodo(idToComplete)
			saveTodo(Todo{
				Id:      todo.Id,
				Title:   todo.Title,
				Created: todo.Created,
				Done:    true,
			})
			return
		}
	}
}

func RemoveTodo(idToRemove string) bool {
	file, err := os.OpenFile("todo.csv", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	lineNumber := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		lineNumber++

		if idToRemove == record[0] {
			deleteLineFromFile("todo.csv", lineNumber)
			return true
		}
	}
	return false
}

func getNewTodoId() int {
	todos := GetTodos()
	maxId := 0
	for _, todo := range todos {
		if todo.Id > maxId {
			maxId = todo.Id
		}
	}
	newId := maxId + 1
	return newId
}

func deleteLineFromFile(filename string, lineToDelete int) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	tempFile, err := os.CreateTemp("", "temp")
	if err != nil {
		return err
	}
	defer tempFile.Close()

	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(tempFile)
	defer writer.Flush()

	lineNum := 1
	for scanner.Scan() {
		if lineNum != lineToDelete {
			fmt.Fprintln(writer, scanner.Text())
		}
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return os.Rename(tempFile.Name(), filename)
}

func saveTodo(todo Todo) {
	file, err := os.OpenFile("todo.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	values := []string{strconv.Itoa(todo.Id), todo.Title, strconv.FormatBool(todo.Done), todo.Created.Format(time.RFC3339)}
	err = writer.Write(values)
	if err != nil {
		log.Fatal(err)
	}
}
