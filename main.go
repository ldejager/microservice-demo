package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"database/sql"

	"github.com/coopernurse/gorp"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
)

var dbmap = initDb()

// Config struct
type DBConfig struct {
	Port     string `default:"3306"`
	Username string
	Password string
	Hostname string `default:"localhost"`
	Database string `default:"tasks"`
}

// Task struct
type Task struct {
	Id          int64 `db:"task_id"`
	Created     int64
	Completed   int64
	Name        string
	Description string
}

func main() {

	defer dbmap.Db.Close()

	router := gin.Default()

	// Create routes
	router.GET("/", DefaultLanding)
	router.GET("/tasks", TasksList)
	router.POST("/tasks", TaskPost)
	router.GET("/tasks/:id", TasksDetail)
	router.GET("/health", Health)
	router.GET("/ping", PingPong)
	router.Run(":8000")

}

func CreateConnectionString() string {

	var s DBConfig
	err := envconfig.Process("db", &s)

	if err != nil {
		log.Fatal(err.Error())
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		s.Username, s.Password, s.Hostname, s.Port, s.Database)
}

func DefaultLanding(c *gin.Context) {
	c.String(200, "Simple Microservices Demo")
}

func createTask(name, description string) Task {
	task := Task{
		Created:     time.Now().UnixNano(),
		Completed:   0,
		Name:        name,
		Description: description,
	}

	err := dbmap.Insert(&task)

	HandleError(err, "Failed to create task")
	return task
}

func getTask(task_id int) Task {
	task := Task{}
	err := dbmap.SelectOne(&task, "select * from tasks where task_id=?", task_id)
	HandleError(err, "Failed to select task")
	return task
}

func TasksList(c *gin.Context) {
	var tasks []Task
	_, err := dbmap.Select(&tasks, "select * from tasks order by task_id")
	HandleError(err, "Select failed")
	content := gin.H{}
	for k, v := range tasks {
		content[strconv.Itoa(k)] = v
	}
	c.IndentedJSON(200, content)
}

func TasksDetail(c *gin.Context) {
	task_id := c.Params.ByName("id")
	t_id, _ := strconv.Atoi(task_id)
	task := getTask(t_id)
	content := gin.H{"name": task.Name, "description": task.Description}
	c.IndentedJSON(200, content)
}

func TaskPost(c *gin.Context) {
	var json Task

	c.Bind(&json)
	task := createTask(json.Name, json.Description)
	if task.Name == json.Name {
		content := gin.H{
			"result":      "Success",
			"name":        task.Name,
			"description": task.Description,
		}
		c.IndentedJSON(201, content)
	} else {
		c.IndentedJSON(500, gin.H{"result": "An error occured"})
	}
}

func Health(c *gin.Context) {
	connectionString := CreateConnectionString()
	db, err := sql.Open("mysql", connectionString)
	defer db.Close()
	err = db.Ping()
	if err != nil {
		content := gin.H{"health": "dead"}
		c.IndentedJSON(200, content)
	} else {
		content := gin.H{"health": "alive"}
		c.IndentedJSON(200, content)
	}
}

func PingPong(c *gin.Context) {
	c.String(200, "pong")
}

func initDb() *gorp.DbMap {

	log.Println("Starting API...")

	if os.Getenv("DB_USERNAME") == "" {
		log.Fatal("DB_USERNAME must be set and non-empty")
	}
	if os.Getenv("DB_PASSWORD") == "" {
		log.Fatal("DB_PASSWORD must be set and non-empty")
	}
	if os.Getenv("DB_DATABASE") == "" {
		log.Fatal("DB_DATABASE must be set and non-empty")
	}

	connectionString := CreateConnectionString()

	db, err := sql.Open("mysql", connectionString)
	HandleError(err, "Database connection failed!")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	dbmap.AddTableWithName(Task{}, "tasks").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	HandleError(err, "Create tables failed")

	return dbmap
}

func HandleError(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
