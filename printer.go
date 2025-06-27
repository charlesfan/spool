package spool

type Printer interface {
	Name() string
	PrintJPEG(content []byte, metadata map[string]string) (jobID string, err error)
	WatchJob(jobID string, callback func(status string)) error
}
