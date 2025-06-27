package spool

import (
	"fmt"
	"time"
)

type MockPrinter struct {
	name string
}

func (mp *MockPrinter) Name() string {
	return mp.name
}

func (mp *MockPrinter) PrintJPEG(content []byte, metadata map[string]string) (string, error) {
	jobID := fmt.Sprintf("job-%d", time.Now().UnixNano())
	fmt.Printf("[%s] 列印中... JobID: %s, Metadata: %+v\n", mp.name, jobID, metadata)
	return jobID, nil
}

func (mp *MockPrinter) WatchJob(jobID string, callback func(status string)) error {
	go func() {
		time.Sleep(2 * time.Second)
		callback("completed")
	}()
	return nil
}

func NewMockPrinter(n string) *MockPrinter {
	return &MockPrinter{name: n}
}
