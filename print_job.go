package spool

type PrintJob struct {
	ID       string
	Name     string
	Status   string
	Metadata map[string]string
}
