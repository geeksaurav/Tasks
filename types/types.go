package types

//Task is the struct used to indetify tasks
type Task struct {
	Id int
	Title string
	Content string
	Created string
}

// Context is the struct passed to the templates
type Context struct {
	Tasks []Task
	Navigation string
	Search string
	Message string
	CSRFToken string
}