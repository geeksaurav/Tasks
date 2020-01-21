package views

import (
	"io/ioutil"
	"net/http"
	"os"
	// "strconv"
	"strings"
	"html/template"
	"time"
	"db"
	"log"
	"fmt"
	"types"
)

var (
	homeTemplate	*template.Template
	deletedTemplate	*template.Template
	completedTemplate	*template.Template
	loginTemplate	*template.Template
	searchTemplate	*template.Template
	editTemplate	*template.Template
	templates	*template.Template
	message	string
	err	error
)


// PopulateTemplates does what it says
func PopulateTemplates(){
	var allFiles []string
	templatesDir := "./public/templates/"
	files, err := ioutil.ReadDir(templatesDir)

	if err!= nil {
		fmt.Println("Error reading templates Dir")
	}

	for _, file := range files{
		filename := file.Name()
		if strings.HasSuffix(filename, ".html"){
			allFiles = append(allFiles, templatesDir+filename)
		}
	}

	if err!= nil {
		fmt.Println(err)
		os.Exit(1)
	}

	templates = template.Must(template.ParseFiles(allFiles...))

	// template.Must()
	if err!= nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// fmt.Println(templates)

	homeTemplate = templates.Lookup("home.html")
	deletedTemplate = templates.Lookup("deleted.html")
	editTemplate = templates.Lookup("edit.html")
	searchTemplate = templates.Lookup("search.html")
	completedTemplate = templates.Lookup("completed.html")
	loginTemplate = templates.Lookup("login.html")
}

//CompleteTaskFunc hey there its a comment for the linter 
func CompleteTaskFunc(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		message = "all pending tasks GET"
	} else {
		message = "all pending tasks POST"
	}
	w.Write([]byte(message))
}

// DeleteTaskFunc as
func DeleteTaskFunc(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		message = "all pending tasks GET"
	} else {
		message = "all pending tasks POST"
	}
	w.Write([]byte(message))
}

// ShowTrashTaskFunc as
func ShowTrashTaskFunc(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		message = "all pending tasks GET"
	} else {
		message = "all pending tasks POST"
	}
	w.Write([]byte(message))
}

// TrashTaskFunc as
func TrashTaskFunc(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		message = "all pending tasks GET"
	} else {
		message = "all pending tasks POST"
	}
	w.Write([]byte(message))
}

// EditTaskFunc as
func EditTaskFunc(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		message = "all pending tasks GET"
	} else {
		message = "all pending tasks POST"
	}
	w.Write([]byte(message))
}

// ShowCompleteTasksFunc as
func ShowCompleteTasksFunc(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		message = "all pending tasks GET"
	} else {
		message = "all pending tasks POST"
	}
	w.Write([]byte(message))
}

// RestoreTaskFunc as
func RestoreTaskFunc(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		message = "all pending tasks GET"
	} else {
		message = "all pending tasks POST"
	}
	w.Write([]byte(message))
}

// AddTaskFunc as
func AddTaskFunc(w http.ResponseWriter, r *http.Request){
	title := "random title"
    content := "random content"
    truth := AddTask(title, content)
    if truth != nil {
        log.Fatal("Error adding task")
    }
    w.Write([]byte("Added task"))
}

// UpdateTaskFunc as
func UpdateTaskFunc(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		message = "all pending tasks GET"
	} else {
		message = "all pending tasks POST"
	}
	w.Write([]byte(message))
}

// SearchTaskFunc as
func SearchTaskFunc(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		message = "all pending tasks GET"
	} else {
		message = "all pending tasks POST"
	}
	w.Write([]byte(message))
}

// GetLogin as
func GetLogin(w http.ResponseWriter, r *http.Request){
	fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        loginTemplate.Execute(w, nil)
    } else {
        r.ParseForm()
        // logic part of log in
        fmt.Println("username:", r.Form["username"])
        fmt.Println("password:", r.Form["password"])
    }
}

// PostRegister ad
func PostRegister(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		message = "all pending tasks GET"
	} else {
		message = "all pending tasks POST"
	}
	w.Write([]byte(message))
}

// HandleAdmin as
func HandleAdmin(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		message = "all pending tasks GET"
	} else {
		message = "all pending tasks POST"
	}
	w.Write([]byte(message))
}

// PostAddUser af
func PostAddUser(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		message = "all pending tasks GET"
	} else {
		message = "all pending tasks POST"
	}
	w.Write([]byte(message))
}

// PostChange af
func PostChange(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		message = "all pending tasks GET"
	} else {
		message = "all pending tasks POST"
	}
	w.Write([]byte(message))
}

// HandleLogout as
func HandleLogout(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		message = "all pending tasks GET"
	} else {
		message = "all pending tasks POST"
	}
	w.Write([]byte(message))
}

// ShowAllTasksFunc => used to handle the "/" path of  url i.e default one
func ShowAllTasksFunc(w http.ResponseWriter, r *http.Request){
	PopulateTemplates()
	if r.Method == "GET" {
		context := GetTasks()

		if message != "" {
            context.Message = message
		}
		
		context.CSRFToken = "abcd"
		message = ""
		expiration := time.Now().Add(365*24*time.Hour)
		cookie := http.Cookie{Name: "csrftoken", Value: context.CSRFToken, Expires: expiration}
		
		http.SetCookie(w, &cookie)

		homeTemplate.Execute(w, context)

		w.Write([]byte(context.Tasks[1].Title))
	} else {
		message = "Method not allowed"
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

// GetTasks is the function or context which gets all the tasks from the db
func GetTasks() types.Context {
	var task []types.Task
	var context types.Context
	var TaskID int
    var TaskTitle string
    var TaskContent string
    var TaskCreated time.Time
	var getTasksql string
	
	getTasksql = "select id, title, content, created_date from task;"

	rows, err := db.Database.Query(getTasksql)
	if err != nil{
		fmt.Println(err)
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&TaskID, &TaskTitle, &TaskContent, &TaskCreated)
		if err != nil {
			fmt.Println(err)
		}
		TaskCreated = TaskCreated.Local()
		a := types.Task{Id: TaskID, Title: TaskTitle, Content: TaskContent,
            Created: TaskCreated.Format(time.UnixDate)[0:20]}
        task = append(task, a)
	}
	context = types.Context{Tasks: task}
	return context
}

// AddTask is the function which added the to the db
func AddTask(title, content string) error {
    query:="insert into task(title, content, created_date, last_modified_at) values(?,?,datetime(), datetime())"
    restoreSQL, err := db.Database.Prepare(query)
    if err != nil {
        fmt.Println(err)
    }
    tx, err := db.Database.Begin()
    _, err = tx.Stmt(restoreSQL).Exec(title, content)
    if err != nil {
        fmt.Println(err)
        tx.Rollback()
    } else {
        log.Print("insert successful")
        tx.Commit()
    }
    return err
}