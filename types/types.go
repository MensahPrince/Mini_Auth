package types

type DBSTATUS struct {
	Success bool
	Message string
}

type USER struct {
	Password string
	Name     string
	Email    string
}
