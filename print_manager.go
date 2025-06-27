package spool

import (
	"fmt"
	"time"
)

type PrintManager struct {
	Printers     []Printer
	Dispatcher   Dispatcher
	metadataFunc func([]byte) map[string]string
	Callback     func(PrintJob)
}

func NewPrintManager() *PrintManager {
	pm := &PrintManager{Dispatcher: &dispatcher{}}
	pm.metadataFunc = func(img []byte) map[string]string {
		return map[string]string{
			"user": "default",
			"time": time.Now().Format(time.RFC3339),
		}
	}

	return pm
}

func (pm *PrintManager) RegisterPrinter(p Printer) {
	pm.Printers = append(pm.Printers, p)
}

func (pm *PrintManager) SetMetadataFunc(f func([]byte) map[string]string) {
	pm.metadataFunc = f
}

func (pm *PrintManager) SubmitPrintJob(jpeg []byte) error {
	meta := pm.metadataFunc(jpeg)
	printer := pm.Dispatcher.Next(pm.Printers, meta)
	if printer == nil {
		return fmt.Errorf("no printer available")
	}

	jobID, err := printer.PrintJPEG(jpeg, meta)
	if err != nil {
		return err
	}

	go func() {
		_ = printer.WatchJob(jobID, func(status string) {
			result := PrintJob{
				ID:       jobID,
				Name:     printer.Name(),
				Status:   status,
				Metadata: meta,
			}
			go pm.Callback(result)
		})
	}()
	return nil
}
