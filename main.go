package main

import (
    "log"
	"net/http"
	"views"
)

func main(){
	
	http.HandleFunc("/complete/", views.CompleteTaskFunc)
    http.HandleFunc("/delete/", views.DeleteTaskFunc)
    http.HandleFunc("/deleted/", views.ShowTrashTaskFunc)
    http.HandleFunc("/trash/", views.TrashTaskFunc)
    http.HandleFunc("/edit/", views.EditTaskFunc)
    http.HandleFunc("/completed/", views.ShowCompleteTasksFunc)
    http.HandleFunc("/restore/", views.RestoreTaskFunc)
	http.HandleFunc("/add/", views.AddTaskFunc)
	http.HandleFunc("/update/", views.UpdateTaskFunc)
	http.HandleFunc("/search/", views.SearchTaskFunc)
	http.HandleFunc("/login", views.GetLogin)
	http.HandleFunc("/register", views.PostRegister)
	http.HandleFunc("/admin", views.HandleAdmin)
	http.HandleFunc("/add_user", views.PostAddUser)
	http.HandleFunc("/change", views.PostChange)
	http.HandleFunc("/logout", views.HandleLogout)
	http.HandleFunc("/", views.ShowAllTasksFunc)
	
	// http.Handle("/static/", http.FileServer(http.Dir("public")))
	
	log.Print("Running server on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

